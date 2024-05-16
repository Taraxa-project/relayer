package to_eth

import (
	"context"
	"log"
	"math/big"

	tara_client_contract_interface "github.com/Taraxa-project/relayer/clients/eth/tara_client_contract_client/contract_interface"
	tara_rpc_types "github.com/Taraxa-project/relayer/clients/tara/rpc_client/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func transformPillarBlockData(pillarBlockData *tara_rpc_types.PillarBlockData) (block tara_client_contract_interface.PillarBlockWithChanges, signatures []tara_client_contract_interface.CompactSignature) {
	block.Block.Period = big.NewInt(int64(pillarBlockData.PillarBlock.PbftPeriod))
	block.Block.BridgeRoot = pillarBlockData.PillarBlock.BridgeRoot
	block.Block.StateRoot = pillarBlockData.PillarBlock.StateRoot
	block.Block.PrevHash = pillarBlockData.PillarBlock.PreviousBlockHash
	for _, votesCountChange := range pillarBlockData.PillarBlock.VoteCountsChanges {
		block.ValidatorChanges = append(block.ValidatorChanges, tara_client_contract_interface.PillarBlockVoteCountChange{Validator: votesCountChange.Address, Change: votesCountChange.Value})
	}

	for _, signature := range pillarBlockData.Signatures {
		signatures = append(signatures, tara_client_contract_interface.CompactSignature{R: signature.R, Vs: signature.Vs})
	}

	return
}

func (r *Relayer) processBridgeStateProof(finalizeBlocksTx *types.Transaction) {
	// Note: we want to wait until the finalizeBlocks tx is mined to see if it didn't fail. If it did fail, there is no need to call applyState on bridge contract
	log.Println("Waiting for finalize blocks tx to be mined. Tx hash: ", finalizeBlocksTx.Hash())
	finalizeBlocksTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, finalizeBlocksTx)
	if err != nil {
		log.Fatal("WaitMined finalize blocks tx failed. Err: ", err)
	}
	// Tx failed -> status == 0
	if finalizeBlocksTxReceipt.Status == 0 {
		log.Fatal("Finalize blocks tx failed execution. Tx hash: ", finalizeBlocksTx.Hash())
	}

	// Get state with proof from tara bridge contract
	taraStateWithProof, err := r.taraBridge.GetStateWithProof(nil)
	if err != nil {
		log.Fatal("GetStateWithProof err: ", err)
	}

	// Send tara state with proof to eth bridge contract
	applyStateTx, err := r.ethBridge.ApplyState(r.taraAuth, taraStateWithProof)
	if err != nil {
		log.Fatal("ApplyState err: ", err)
	}
	log.Println("Apply state tx sent to eth bridge contracts. Tx hash: ", applyStateTx.Hash(), ". State: ", taraStateWithProof)

	log.Println("Waiting for apply state tx to be mined. Tx hash: ", applyStateTx.Hash())
	applyStateTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, applyStateTx)
	if err != nil {
		log.Fatal("WaitMined apply state tx failed. Err: ", err)
	}
	// Tx failed -> status == 0
	if applyStateTxReceipt.Status == 0 {
		log.Fatal("Apply state tx failed execution. Tx hash: ", applyStateTx.Hash())
	}
}

func (r *Relayer) processNewPillarBlock(pillarBlockData *tara_rpc_types.PillarBlockData) {
	log.Println("Process new pillar block data: ", pillarBlockData)
	pillarBlocksInterval := uint64(r.taraxaNodeConfig.Hardforks.FicusHf.PillarBlocksInterval)
	currentBlockNumber, err := r.taraxaClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("BlockNumber err: ", err)
	}
	expectedLatestPillarBlockPeriod := currentBlockNumber - currentBlockNumber%pillarBlocksInterval

	latestFinalizedPillarBlock, err := r.taraClientOnEth.GetFinalized(nil)
	if err != nil {
		log.Fatal("GetFinalizedPillarBlock err: ", err)
	}
	latestFinalizedPillarBlockPeriod := latestFinalizedPillarBlock.Block.Period.Uint64()

	if latestFinalizedPillarBlockPeriod < pillarBlocksInterval {
		log.Fatal("Latest finalized pillar block period is: ", latestFinalizedPillarBlockPeriod, ". Should be at least ", pillarBlocksInterval)
	}

	numOfProcessedBlocks := uint64(0)
	maxNumOfBlocksInBatch := uint64(10)

	var blocks []tara_client_contract_interface.PillarBlockWithChanges
	var blocksSignatures [][]tara_client_contract_interface.CompactSignature

	// Process all missing pillar blocks between latestFinalizedPillarBlockPeriod + pillarBlocksInterval and expectedLatestPillarBlockPeriod
	for period := latestFinalizedPillarBlockPeriod + pillarBlocksInterval; period <= expectedLatestPillarBlockPeriod; period += pillarBlocksInterval {
		if pillarBlockData != nil && pillarBlockData.PillarBlock.PbftPeriod == period {
			block, signatures := transformPillarBlockData(pillarBlockData)

			blocks = append(blocks, block)
			blocksSignatures = append(blocksSignatures, signatures)
		} else {
			tmpPillarBlockData, err := r.taraxaClient.GetPillarBlockData(period, true)
			if err != nil {
				log.Fatal("GetPillarBlockData err: ", err)
			} else {
				// TODO: might be empty because nodes dont have it ????
				block, signatures := transformPillarBlockData(tmpPillarBlockData)

				blocks = append(blocks, block)
				blocksSignatures = append(blocksSignatures, signatures)
			}
		}
		numOfProcessedBlocks++

		// Send blocks into the tara client contract on ethereum
		if numOfProcessedBlocks == maxNumOfBlocksInBatch || period == expectedLatestPillarBlockPeriod {
			finalizeBlocksTx, err := r.taraClientOnEth.FinalizeBlocks(r.ethAuth, blocks, blocksSignatures[len(blocksSignatures)-1])
			if err != nil {
				log.Fatal("FinalizeBlocks tx failed: ", err)
			}
			log.Println("Finalize blocks tx sent. Tx hash: ", finalizeBlocksTx.Hash(), ". Blocks: ", blocks, ", last block signaturtures: ", blocksSignatures)

			// Clear blocks & signatures slices before processing new ones
			blocks = nil
			blocksSignatures = nil

			r.processBridgeStateProof(finalizeBlocksTx)
		}
	}
}

func (r *Relayer) ProcessPillarBlocks(ctx context.Context) {
	r.processNewPillarBlock(nil)
	// Listen to new pillar block data
	newPillarBlockData := make(chan *tara_rpc_types.PillarBlockData)
	sub, err := r.taraxaClient.Client.Client().EthSubscribe(ctx, newPillarBlockData, "newPillarBlockData", "includeSignatures")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case pillarBlockData := <-newPillarBlockData:
			// Send new pillar block to tara client contract on ethereum
			r.processNewPillarBlock(pillarBlockData)
		}
	}
}

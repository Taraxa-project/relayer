package to_eth

import (
	"context"
	"math/big"

	log "github.com/sirupsen/logrus"

	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	tara_client_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client/contract_interface"
	tara_rpc_types "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (r *Relayer) getStateWithProof(epoch *big.Int, period *big.Int) (*bridge_contract_interface.SharedStructsStateWithProof, error) {
	opts := bind.CallOpts{}
	if period != nil {
		opts.BlockNumber = period
	}
	taraStateWithProof, err := r.taraBridge.GetStateWithProof(&opts)
	if err != nil {
		return nil, err
	}

	// TODO: implement some binary search?
	bigPillarBlocksInterval := big.NewInt(0).SetUint64(uint64(r.taraxaNodeConfig.Hardforks.FicusHf.PillarBlocksInterval))
	if epoch == nil || epoch.Cmp(taraStateWithProof.State.Epoch) == 0 {
		return &taraStateWithProof, nil
	}

	if taraStateWithProof.State.Epoch.Cmp(epoch) > 0 {
		return r.getStateWithProof(epoch, period.Sub(period, bigPillarBlocksInterval))
	}

	return r.getStateWithProof(epoch, period.Add(period, bigPillarBlocksInterval))
}

func (r *Relayer) bridgeState() {
	lastFinalizedEpoch, err := r.taraBridge.FinalizedEpoch(nil)
	if err != nil {
		log.WithError(err).Fatal("lastFinalizedEpoch")
	}
	lastAppliedEpoch, err := r.ethBridge.FinalizedEpoch(nil)
	if err != nil {
		log.WithError(err).Fatal("lastAppliedEpoch")
	}
	if lastFinalizedEpoch == lastAppliedEpoch {
		return
	}

	for epoch := lastFinalizedEpoch; epoch.Cmp(lastAppliedEpoch) <= 0; epoch.Add(epoch, big.NewInt(1)) {
		taraStateWithProof, err := r.getStateWithProof(epoch, nil)
		if err != nil {
			log.WithError(err).WithField("epoch", epoch).Fatal("getStateWithProof")
		}
		applyStateTx, err := r.ethBridge.ApplyState(r.taraAuth, *taraStateWithProof)
		if err != nil {
			log.WithError(err).Fatal("ApplyState")
		}
		log.WithFields(log.Fields{"tx_hash": applyStateTx.Hash, "state": taraStateWithProof}).Println("Apply state tx sent to eth bridge contracts")

		log.WithField("hash", applyStateTx.Hash()).Info("Waiting for apply state tx to be mined")
		applyStateTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, applyStateTx)
		if err != nil {
			log.WithError(err).Fatal("WaitMined apply state tx failed")
		}
		// Tx failed -> status == 0
		if applyStateTxReceipt.Status == 0 {
			log.WithField("hash", applyStateTx.Hash()).Fatal("Apply state tx failed execution")
		}
	}
}

func (r *Relayer) processPillarBlocks() {
	pillarBlocksInterval := uint64(r.taraxaNodeConfig.Hardforks.FicusHf.PillarBlocksInterval)
	ficusHfBlockNum := uint64(r.taraxaNodeConfig.Hardforks.FicusHf.BlockNum)
	currentBlockNumber, err := r.taraxaClient.BlockNumber(context.Background())
	if err != nil {
		log.WithError(err).Fatal("BlockNumber")
	}
	expectedLatestPillarBlockPeriod := currentBlockNumber - currentBlockNumber%pillarBlocksInterval

	latestFinalizedPillarBlock, err := r.taraClientOnEth.GetFinalized(nil)
	if err != nil {
		log.WithError(err).Fatal("GetFinalizedPillarBlock")
	}
	latestFinalizedPillarBlockPeriod := latestFinalizedPillarBlock.Block.Period.Uint64()

	numOfProcessedBlocks := uint64(0)
	maxNumOfBlocksInBatch := uint64(50)
	if latestFinalizedPillarBlockPeriod == 0 {
		latestFinalizedPillarBlockPeriod = ficusHfBlockNum - pillarBlocksInterval
		// TODO: Do we need this limit here for the initialization?
		maxNumOfBlocksInBatch = 1
	}

	var blocks []tara_client_contract_interface.PillarBlockWithChanges
	var blocksSignatures [][]tara_client_contract_interface.CompactSignature

	// Process all missing pillar blocks between latestFinalizedPillarBlockPeriod + pillarBlocksInterval and expectedLatestPillarBlockPeriod
	period := latestFinalizedPillarBlockPeriod + pillarBlocksInterval
	for ; period <= expectedLatestPillarBlockPeriod; period += pillarBlocksInterval {
		tmpPillarBlockData, err := r.taraxaClient.GetPillarBlockData(period, true)
		if err != nil {
			log.Fatal("GetPillarBlockData err: ", err)
		} else {
			// TODO: might be empty because nodes dont have it ????
			block, signatures := transformPillarBlockData(tmpPillarBlockData)

			blocks = append(blocks, block)
			blocksSignatures = append(blocksSignatures, signatures)
		}
		numOfProcessedBlocks++

		// Send blocks into the tara client contract on ethereum
		if numOfProcessedBlocks == maxNumOfBlocksInBatch {
			break
		}
	}
	finalizeBlocksTx, err := r.taraClientOnEth.FinalizeBlocks(r.ethAuth, blocks, blocksSignatures[len(blocksSignatures)-1])
	if err != nil {
		log.Fatal("FinalizeBlocks tx failed: ", err)
	}
	log.Println("Finalize blocks tx sent. Tx hash: ", finalizeBlocksTx.Hash(), ". Blocks: ", blocks, ", last block signaturtures: ", blocksSignatures)
	log.Println("Waiting for finalize blocks tx to be mined. Tx hash: ", finalizeBlocksTx.Hash())
	finalizeBlocksTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, finalizeBlocksTx)
	if err != nil {
		log.Fatal("WaitMined finalize blocks tx failed. Err: ", err)
	}
	// Tx failed -> status == 0
	if finalizeBlocksTxReceipt.Status == 0 {
		log.Fatal("Finalize blocks tx failed execution. Tx hash: ", finalizeBlocksTx.Hash())
	}
	// This means that we have more blocks to process
	if period != expectedLatestPillarBlockPeriod {
		r.processPillarBlocks()
	} else {
		r.bridgeState()
	}
}

func (r *Relayer) ListenForPillarBlockUpdates(ctx context.Context) {
	// sync pillar blocks to tara client contract on ethereum
	r.processPillarBlocks()
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
		case <-newPillarBlockData:
			r.processPillarBlocks()
		}
	}
}

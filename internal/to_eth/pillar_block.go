package to_eth

import (
	"context"
	"math/big"

	log "github.com/sirupsen/logrus"

	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	tara_client_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client/contract_interface"
	tara_rpc_types "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client/types"
	"github.com/ethereum/go-ethereum"
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

func (r *Relayer) getStateWithProof(epoch *big.Int, block_num *big.Int) (*bridge_contract_interface.SharedStructsStateWithProof, error) {
	if block_num == nil {
		block, err := r.taraxaClient.Client.BlockByNumber(context.Background(), nil)
		if err != nil || block == nil {
			log.WithField("block", block).WithError(err).Fatal("BlockByNumber")
		}
		block_num = block.Number()
	}
	opts := bind.CallOpts{BlockNumber: block_num}

	taraStateWithProof, err := r.taraBridge.GetStateWithProof(&opts)
	log.WithField("state", taraStateWithProof).WithField("epoch", epoch).Println("GetStateWithProof")
	if err != nil {
		log.WithError(err).Error("taraBridge.GetStateWithProof")
		return nil, err
	}

	// TODO: implement some binary search?
	bigPillarBlocksInterval := big.NewInt(0).SetUint64(uint64(r.taraxaNodeConfig.Hardforks.FicusHf.PillarBlocksInterval))
	if epoch == nil || epoch.Cmp(taraStateWithProof.State.Epoch) == 0 {
		return &taraStateWithProof, nil
	}

	if taraStateWithProof.State.Epoch.Cmp(epoch) > 0 {
		return r.getStateWithProof(epoch, block_num.Sub(block_num, bigPillarBlocksInterval))
	}

	return r.getStateWithProof(epoch, block_num.Add(block_num, bigPillarBlocksInterval))
}

func (r *Relayer) bridgeState() {
	lastFinalizedEpoch, err := r.taraBridge.FinalizedEpoch(nil)
	if err != nil {
		log.WithError(err).Fatal("lastFinalizedEpoch")
	}
	lastAppliedEpoch, err := r.ethBridge.AppliedEpoch(nil)
	if err != nil {
		log.WithError(err).Fatal("lastAppliedEpoch")
	}
	if lastFinalizedEpoch.Cmp(lastAppliedEpoch) == 0 {
		log.WithFields(log.Fields{"lastFinalizedEpoch": lastFinalizedEpoch, "lastAppliedEpoch": lastAppliedEpoch}).Info("No new state to pass")
		return
	}
	epoch := lastAppliedEpoch.Add(lastAppliedEpoch, big.NewInt(1))
	for ; epoch.Cmp(lastFinalizedEpoch) <= 0; epoch.Add(epoch, big.NewInt(1)) {
		log.WithField("epoch", epoch).Info("Applying state")
		taraStateWithProof, err := r.getStateWithProof(epoch, nil)
		if err != nil {
			log.WithError(err).WithField("epoch", epoch).Fatal("getStateWithProof")
		}
		local := r.ethAuth
		// TODO: fix the estimation?
		local.GasLimit = 1000000
		applyStateTx, err := r.ethBridge.ApplyState(local, *taraStateWithProof)
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
		log.WithField("hash", applyStateTx.Hash()).Info("Apply state tx mined")
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

	maxNumOfBlocksInBatch := 20
	if latestFinalizedPillarBlockPeriod == 0 {
		latestFinalizedPillarBlockPeriod = ficusHfBlockNum - pillarBlocksInterval
		// TODO: Do we need this limit here for the initialization?
		maxNumOfBlocksInBatch = 1
	}

	var blocks []tara_client_contract_interface.PillarBlockWithChanges
	var blocksSignatures [][]tara_client_contract_interface.CompactSignature

	// Process all missing pillar blocks between latestFinalizedPillarBlockPeriod + pillarBlocksInterval and expectedLatestPillarBlockPeriod
	pendingBridgeRoot := r.lastAppliedBridgeRoot
	period := latestFinalizedPillarBlockPeriod + pillarBlocksInterval
	for ; period <= expectedLatestPillarBlockPeriod; period += pillarBlocksInterval {
		tmpPillarBlockData, err := r.taraxaClient.GetPillarBlockData(period, true)
		log.WithFields(log.Fields{"block": tmpPillarBlockData, "period": period}).Println("GetPillarBlockData")
		if err == ethereum.NotFound {
			log.WithField("period", period).Info("Pillar block not found, probably not finalized yet")
			break
		}
		if err != nil {
			log.WithError(err).Error("GetPillarBlockData")
		} else {
			// TODO: might be empty because nodes don't have it ????
			block, signatures := transformPillarBlockData(tmpPillarBlockData)

			if pendingBridgeRoot != block.Block.BridgeRoot {
				pendingBridgeRoot = block.Block.BridgeRoot
			}

			blocks = append(blocks, block)
			blocksSignatures = append(blocksSignatures, signatures)
		}

		// Send blocks into the tara client contract on ethereum
		if len(blocks) == maxNumOfBlocksInBatch {
			break
		}
	}
	// TODO: don't process without blocks of if bridgeRoot wasn't change
	if len(blocks) == 0 { // || (len(blocks) == maxNumOfBlocksInBatch && pendingBridgeRoot == r.lastAppliedBridgeRoot) {
		log.WithField("blocks", len(blocks)).Info("No new pillar blocks to process")
		r.bridgeState()
		return
	}
	finalizeBlocksTx, err := r.taraClientOnEth.FinalizeBlocks(r.ethAuth, blocks, blocksSignatures[len(blocksSignatures)-1])
	if err != nil {
		log.Fatal("FinalizeBlocks tx failed: ", err)
	}
	log.Println("Finalize blocks tx sent. Tx hash: ", finalizeBlocksTx.Hash(), ". Blocks: ", blocks, ", last block signatures: ", blocksSignatures)
	log.Println("Waiting for finalize blocks tx to be mined. Tx hash: ", finalizeBlocksTx.Hash())
	finalizeBlocksTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, finalizeBlocksTx)
	if err != nil {
		log.Fatal("WaitMined finalize blocks tx failed. Err: ", err)
	}
	// Tx failed -> status == 0
	if finalizeBlocksTxReceipt.Status == 0 {
		log.Fatal("Finalize blocks tx failed execution. Tx hash: ", finalizeBlocksTx.Hash())
	}
	r.lastAppliedBridgeRoot = pendingBridgeRoot
	// This means that we have more blocks to process
	if period != expectedLatestPillarBlockPeriod {
		log.WithField("period", period).WithField("expectedLatest", expectedLatestPillarBlockPeriod).Info("We have more pillar blocks, processing next batch")
		r.processPillarBlocks()
	} else {
		log.Info("All pillar blocks processed, syncing bridge state")
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

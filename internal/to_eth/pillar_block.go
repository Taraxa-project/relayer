package to_eth

import (
	"context"
	"math/big"
	"time"

	"relayer/bindings/BridgeBase"
	"relayer/bindings/TaraClient"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func transformPillarBlockData(pillarBlockData *PillarBlockData) (block TaraClient.PillarBlockWithChanges, signatures []TaraClient.CompactSignature) {
	block.Block.Period = big.NewInt(int64(pillarBlockData.PillarBlock.PbftPeriod))
	block.Block.BridgeRoot = pillarBlockData.PillarBlock.BridgeRoot
	block.Block.StateRoot = pillarBlockData.PillarBlock.StateRoot
	block.Block.PrevHash = pillarBlockData.PillarBlock.PreviousBlockHash
	block.Block.Epoch = big.NewInt(int64(pillarBlockData.PillarBlock.Epoch))
	for _, votesCountChange := range pillarBlockData.PillarBlock.VoteCountsChanges {
		block.ValidatorChanges = append(block.ValidatorChanges, TaraClient.PillarBlockVoteCountChange{Validator: votesCountChange.Address, Change: votesCountChange.Value})
	}

	for _, signature := range pillarBlockData.Signatures {
		signatures = append(signatures, TaraClient.CompactSignature{R: signature.R, Vs: signature.Vs})
	}

	return
}

func (r *Relayer) getStateWithProof(epoch *big.Int, block_num *big.Int) (*BridgeBase.SharedStructsStateWithProof, error) {
	if block_num == nil {
		block, err := r.taraxaClient.BlockByNumber(context.Background(), nil)
		if err != nil || block == nil {
			r.log.WithField("block", block).WithError(err).Fatal("BlockByNumber")
		}
		block_num = block.Number()
	}
	opts := bind.CallOpts{BlockNumber: block_num}

	taraStateWithProof, err := r.taraBridge.GetStateWithProof(&opts)
	r.log.WithField("state", taraStateWithProof).WithField("epoch", epoch).Println("GetStateWithProof")
	if err != nil {
		r.log.WithError(err).Error("taraBridge.GetStateWithProof")
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
		r.log.WithError(err).Fatal("lastFinalizedEpoch")
	}
	r.latestAppliedEpoch, err = r.ethBridge.AppliedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Fatal("lastAppliedEpoch")
	}
	if lastFinalizedEpoch.Cmp(r.latestAppliedEpoch) == 0 {
		r.log.WithFields(log.Fields{"lastFinalizedEpoch": lastFinalizedEpoch, "latestAppliedEpoch": r.latestAppliedEpoch}).Info("No new state to pass")
		return
	}
	if r.latestAppliedEpoch.Cmp(r.latestClientEpoch) == 0 {
		r.log.WithFields(log.Fields{"r.latestAppliedEpoch": r.latestAppliedEpoch, "r.latestClientEpoch": r.latestClientEpoch}).Info("We don't have a pillar block with this epoch in the client")
		return
	}

	epoch := big.NewInt(0)
	epoch.Add(r.latestAppliedEpoch, big.NewInt(1))

	for ; epoch.Cmp(lastFinalizedEpoch) <= 0; epoch.Add(epoch, big.NewInt(1)) {
		r.log.WithField("epoch", epoch).Info("Applying state")
		taraStateWithProof, err := r.getStateWithProof(epoch, nil)
		if err != nil {
			r.log.WithError(err).WithField("epoch", epoch).Fatal("getStateWithProof")
		}
		applyStateTx, err := r.ethBridge.ApplyState(r.ethAuth, *taraStateWithProof)
		if err != nil {
			r.log.WithError(err).Fatal("ApplyState")
		}
		r.log.WithFields(log.Fields{"tx_hash": applyStateTx.Hash()}).Info("Apply state tx sent to eth bridge contract")

		r.log.WithField("hash", applyStateTx.Hash()).Debug("Waiting for apply state tx to be mined")
		applyStateTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, applyStateTx)

		if err != nil {
			r.log.WithError(err).WithField("hash", applyStateTx.Hash()).Fatal("WaitMined apply state tx failed")
		}
		// Tx failed -> status == 0
		if applyStateTxReceipt.Status == 0 {
			r.log.WithField("hash", applyStateTx.Hash()).Fatal("Apply state tx failed execution")
		}
		r.log.WithField("hash", applyStateTx.Hash()).Info("Apply state tx mined")
	}
}

func (r *Relayer) processPillarBlocks() {
	pillarBlocksInterval := uint64(r.taraxaNodeConfig.Hardforks.FicusHf.PillarBlocksInterval)
	ficusHfBlockNum := uint64(r.taraxaNodeConfig.Hardforks.FicusHf.BlockNum)
	currentBlockNumber, err := r.taraxaClient.BlockNumber(context.Background())
	if err != nil {
		r.log.WithError(err).Fatal("BlockNumber")
	}
	expectedLatestPillarBlockPeriod := currentBlockNumber - currentBlockNumber%pillarBlocksInterval

	latestFinalizedPillarBlock, err := r.taraClientOnEth.GetFinalized(nil)
	if err != nil {
		r.log.WithError(err).Fatal("GetFinalizedPillarBlock")
	}
	latestFinalizedPillarBlockPeriod := latestFinalizedPillarBlock.Block.Period.Uint64()

	maxNumOfBlocksInBatch := 20
	if latestFinalizedPillarBlockPeriod == 0 {
		latestFinalizedPillarBlockPeriod = ficusHfBlockNum - pillarBlocksInterval
		// TODO: Do we need this limit here for the initialization?
		maxNumOfBlocksInBatch = 1
	}

	var blocks []TaraClient.PillarBlockWithChanges
	var blocksSignatures [][]TaraClient.CompactSignature

	// Process all missing pillar blocks between latestFinalizedPillarBlockPeriod + pillarBlocksInterval and expectedLatestPillarBlockPeriod
	pendingBridgeRoot := r.latestBridgeRoot
	pendingEpoch := r.latestClientEpoch
	period := latestFinalizedPillarBlockPeriod + pillarBlocksInterval
	for ; period <= expectedLatestPillarBlockPeriod; period += pillarBlocksInterval {
		tmpPillarBlockData, err := r.taraxaClient.GetPillarBlockData(period, true)
		r.log.WithFields(log.Fields{"block": tmpPillarBlockData, "period": period}).Debug("GetPillarBlockData")
		if err == ethereum.NotFound {
			r.log.WithField("period", period).Info("Pillar block not found, probably not finalized yet")
			break
		}
		if err != nil {
			r.log.WithError(err).Error("GetPillarBlockData")
		} else {
			// TODO: might be empty because nodes don't have it ????
			block, signatures := transformPillarBlockData(tmpPillarBlockData)

			if pendingBridgeRoot != block.Block.BridgeRoot {
				pendingBridgeRoot = block.Block.BridgeRoot
				pendingEpoch = block.Block.Epoch
			}

			blocks = append(blocks, block)
			blocksSignatures = append(blocksSignatures, signatures)
		}

		// Send blocks into the tara client contract on ethereum
		if len(blocks) == maxNumOfBlocksInBatch {
			break
		}
	}
	// TODO: don't process without blocks of if bridgeRoot wasn't changed, but send it if we have maxNumOfBlocksInBatch
	if len(blocks) == 0 || (pendingBridgeRoot == r.latestBridgeRoot && len(blocks) != maxNumOfBlocksInBatch) {
		r.log.WithField("blocks", len(blocks)).Info("No new pillar blocks to process")
		return
	}

	time.Sleep(30 * time.Second)
	finalizeBlocksTx, err := r.taraClientOnEth.FinalizeBlocks(r.ethAuth, blocks, blocksSignatures[len(blocksSignatures)-1])
	if err != nil {
		r.log.Fatal("FinalizeBlocks tx failed: ", err)
	}
	r.log.WithFields(logrus.Fields{"hash": finalizeBlocksTx.Hash(), "blocks_count": len(blocks)}).Info("Waiting for finalize blocks tx to be mined")
	finalizeBlocksTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, finalizeBlocksTx)
	if err != nil {
		r.log.Fatal("WaitMined finalize blocks tx failed. Err: ", err)
	}
	// Tx failed -> status == 0
	if finalizeBlocksTxReceipt.Status == 0 {
		r.log.Fatal("Finalize blocks tx failed execution. Tx hash: ", finalizeBlocksTx.Hash())
	}

	r.latestBridgeRoot = pendingBridgeRoot
	r.latestClientEpoch = pendingEpoch

	// This means that we have more blocks to process
	if period != expectedLatestPillarBlockPeriod {
		r.log.WithFields(logrus.Fields{"period": period, "expectedLatest": expectedLatestPillarBlockPeriod}).Info("Processing next batch")
		r.processPillarBlocks()
	}

	r.log.Info("All pillar blocks processed, syncing bridge state")
	r.bridgeState()
}

func (r *Relayer) ListenForPillarBlockUpdates(ctx context.Context) {
	// Listen to new pillar block data
	newPillarBlockData := make(chan *PillarBlockData)
	sub, err := r.taraxaClient.Client.Client().EthSubscribe(ctx, newPillarBlockData, "newPillarBlockData", "includeSignatures")
	if err != nil {
		r.log.WithError(err).Fatal("Failed to subscribe to new pillar block data")
	}

	for {
		select {
		case err := <-sub.Err():
			r.log.WithError(err).Fatal("Subscription error")
		case <-newPillarBlockData:
			r.processPillarBlocks()
		}
	}
}

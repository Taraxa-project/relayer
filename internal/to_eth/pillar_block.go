package to_eth

import (
	"context"
	"time"

	"relayer/bindings/TaraClient"
	"relayer/internal/types"

	log "github.com/sirupsen/logrus"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (r *Relayer) processPillarBlocks() {
	pillarBlocksInterval := uint64(r.taraxaNodeConfig.Hardforks.FicusHf.PillarBlocksInterval)
	ficusHfBlockNum := uint64(r.taraxaNodeConfig.Hardforks.FicusHf.BlockNum)
	currentBlockNumber, err := r.taraxaClient.BlockNumber(context.Background()) //nolint
	if err != nil {
		r.log.WithError(err).Panic("BlockNumber")
	}
	expectedLatestPillarBlockPeriod := currentBlockNumber - currentBlockNumber%pillarBlocksInterval

	latestFinalizedPillarBlock, err := r.taraClientOnEth.GetFinalized(nil)
	if err != nil {
		r.log.WithError(err).Panic("GetFinalizedPillarBlock")
	}
	latestFinalizedPillarBlockPeriod := latestFinalizedPillarBlock.Block.Period.Uint64()

	if latestFinalizedPillarBlockPeriod == 0 {
		latestFinalizedPillarBlockPeriod = ficusHfBlockNum - pillarBlocksInterval
	}

	var blocks []TaraClient.PillarBlockWithChanges

	// Process all missing pillar blocks between latestFinalizedPillarBlockPeriod + pillarBlocksInterval and expectedLatestPillarBlockPeriod
	pendingBridgeRoot := r.latestBridgeRoot
	pendingEpoch := r.latestClientEpoch
	lastProcessedPillarBlock := &types.PillarBlockData{}
	period := latestFinalizedPillarBlockPeriod + pillarBlocksInterval
	for ; period <= expectedLatestPillarBlockPeriod; period += pillarBlocksInterval {
		lastProcessedPillarBlock, err = r.taraxaClient.GetPillarBlockData(period)
		r.StakeState.UpdateState(&lastProcessedPillarBlock.PillarBlock)

		r.log.WithFields(log.Fields{"block": lastProcessedPillarBlock, "period": period}).Trace("GetPillarBlockData")
		if err == ethereum.NotFound {
			r.log.WithField("period", period).Debug("Pillar block not found, probably not finalized yet")
			break
		}
		if err != nil {
			r.log.WithError(err).Error("GetPillarBlockData")
			break
		}

		block := lastProcessedPillarBlock.TransformPillarBlockData()

		if pendingBridgeRoot != block.Block.BridgeRoot {
			pendingBridgeRoot = block.Block.BridgeRoot
			pendingEpoch = block.Block.Epoch
		}

		blocks = append(blocks, block)

		// Send blocks into the tara client contract on ethereum
		if len(blocks) == r.pillarBlocksInBatch {
			break
		}
	}
	// TODO: don't process without blocks of if bridgeRoot wasn't changed, but send it if we have r.pillarBlocksInBatch
	if len(blocks) == 0 || (pendingBridgeRoot == r.latestBridgeRoot && len(blocks) != r.pillarBlocksInBatch) {
		r.log.WithField("blocks", len(blocks)).Debug("No new pillar blocks to process")
		return
	}

	reducedSignatures, err := r.StakeState.ReduceSignatures(lastProcessedPillarBlock)
	if err != nil {
		r.log.WithError(err).Panic("ReduceSignatures")
	}
	finalizeBlocksTx, err := r.taraClientOnEth.FinalizeBlocks(r.ethAuth, blocks, reducedSignatures)
	if err != nil {
		r.log.Panic("TaraClient.FinalizeBlocks tx failed: ", err)
	}
	r.log.WithFields(log.Fields{"hash": finalizeBlocksTx.Hash(), "blocks_count": len(blocks), "period": period}).Info("Waiting for finalize blocks tx to be mined")
	finalizeBlocksTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, finalizeBlocksTx)
	if err != nil {
		r.log.Panic("WaitMined TaraClient.FinalizeBlocks tx failed. Err: ", err)
	}
	// Tx failed -> status == 0
	if finalizeBlocksTxReceipt.Status == 0 {
		r.log.Panic("TaraClient.FinalizeBlocks tx failed execution. Tx hash: ", finalizeBlocksTx.Hash())
	}
	time.Sleep(30 * time.Second)
	r.latestBridgeRoot = pendingBridgeRoot
	r.latestClientEpoch = pendingEpoch

	// This means that we have more blocks to process
	if period != expectedLatestPillarBlockPeriod {
		r.log.WithFields(log.Fields{"period": period, "expectedLatest": expectedLatestPillarBlockPeriod}).Debug("Processing next batch")
		r.processPillarBlocks()
	}

	r.log.Trace("All pillar blocks processed, syncing bridge state")
	r.bridgeState()
}

func (r *Relayer) ListenForPillarBlockUpdates(ctx context.Context) {
	// Listen to new pillar block data
	newPillarBlockData := make(chan *types.PillarBlockData)
	sub, err := r.taraxaClient.Client.Client().EthSubscribe(ctx, newPillarBlockData, "newPillarBlockData", true)
	if err != nil {
		r.log.WithError(err).Panic("Failed to subscribe to new pillar block data")
	}

	for {
		select {
		case err := <-sub.Err():
			r.log.WithError(err).Panic("Subscription error")
		case <-newPillarBlockData:
			r.processPillarBlocks()
		}
	}
}

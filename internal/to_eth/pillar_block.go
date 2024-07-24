package to_eth

import (
	"context"
	"time"

	"relayer/bindings/TaraClient"
	"relayer/internal/types"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

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

	if latestFinalizedPillarBlockPeriod == 0 {
		latestFinalizedPillarBlockPeriod = ficusHfBlockNum - pillarBlocksInterval
		// TODO: Do we need this limit here for the initialization?
		r.pillarBlocksInBatch = 1
	}

	var blocks []TaraClient.PillarBlockWithChanges
	var blocksSignatures [][]TaraClient.CompactSignature

	// Process all missing pillar blocks between latestFinalizedPillarBlockPeriod + pillarBlocksInterval and expectedLatestPillarBlockPeriod
	pendingBridgeRoot := r.latestBridgeRoot
	pendingEpoch := r.latestClientEpoch
	period := latestFinalizedPillarBlockPeriod + pillarBlocksInterval
	for ; period <= expectedLatestPillarBlockPeriod; period += pillarBlocksInterval {

		tmpPillarBlockData, err := r.taraxaClient.GetPillarBlockData(period)
		r.StakeState.UpdateState(&tmpPillarBlockData.PillarBlock)

		r.log.WithFields(log.Fields{"block": tmpPillarBlockData, "period": period}).Debug("GetPillarBlockData")
		if err == ethereum.NotFound {
			r.log.WithField("period", period).Debug("Pillar block not found, probably not finalized yet")
			break
		}
		if err != nil {
			r.log.WithError(err).Error("GetPillarBlockData")
		} else {
			// TODO: might be empty because nodes don't have it ????
			reducetSignatures, err := r.StakeState.ReduceSignatures(tmpPillarBlockData)
			if err == nil {
				tmpPillarBlockData.Signatures = reducetSignatures
			}

			block, signatures := tmpPillarBlockData.TransformPillarBlockData()

			if pendingBridgeRoot != block.Block.BridgeRoot {
				pendingBridgeRoot = block.Block.BridgeRoot
				pendingEpoch = block.Block.Epoch
			}

			blocks = append(blocks, block)
			blocksSignatures = append(blocksSignatures, signatures)
		}

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
	time.Sleep(30 * time.Second)
	r.latestBridgeRoot = pendingBridgeRoot
	r.latestClientEpoch = pendingEpoch

	// This means that we have more blocks to process
	if period != expectedLatestPillarBlockPeriod {
		r.log.WithFields(logrus.Fields{"period": period, "expectedLatest": expectedLatestPillarBlockPeriod}).Debug("Processing next batch")
		r.processPillarBlocks()
	}

	r.log.Debug("All pillar blocks processed, syncing bridge state")
	r.bridgeState()
}

func (r *Relayer) ListenForPillarBlockUpdates(ctx context.Context) {
	// Listen to new pillar block data
	newPillarBlockData := make(chan *types.PillarBlockData)
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

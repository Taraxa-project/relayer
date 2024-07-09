package to_eth

import (
	"context"
	"math/big"
	"relayer/bindings/BridgeBase"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/sirupsen/logrus"
)

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

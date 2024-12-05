package to_eth

import (
	"context"
	"math/big"
	"relayer/bindings/BridgeBase"

	"relayer/internal/proof"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/sirupsen/logrus"
)

func (r *Relayer) LatestBlockNumber() (*big.Int, error) {
	block, err := r.taraxaClient.BlockByNumber(context.Background(), nil)
	if err == nil {
		return block.Number(), nil
	}
	return nil, err
}

func (r *Relayer) GetStateWithProof(opts *bind.CallOpts) (BridgeBase.SharedStructsStateWithProof, error) {
	return r.taraBridge.GetStateWithProof(opts)
}

func (r *Relayer) FinalizationInterval() *big.Int {
	return big.NewInt(0).SetUint64(uint64(r.taraxaNodeConfig.Hardforks.FicusHf.PillarBlocksInterval))
}

func (r *Relayer) bridgeState() {
	lastFinalizedEpoch, err := r.taraBridge.FinalizedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Panic("lastFinalizedEpoch")
	}
	r.latestAppliedEpoch, err = r.ethBridge.AppliedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Panic("lastAppliedEpoch")
	}
	if lastFinalizedEpoch.Cmp(r.latestAppliedEpoch) == 0 {
		r.log.WithFields(log.Fields{"lastFinalizedEpoch": lastFinalizedEpoch, "latestAppliedEpoch": r.latestAppliedEpoch}).Debug("No new state to pass")
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
		taraStateWithProof, err := proof.GetStateWithProof(r, r.log, epoch, nil)
		if err != nil {
			r.log.WithError(err).WithField("epoch", epoch).Panic("getStateWithProof")
		}
		applyStateTx, err := r.ethBridge.ApplyState(r.ethAuth, *taraStateWithProof)
		if err != nil {
			r.log.WithError(err).Panic("ApplyState")
		}
		r.log.WithFields(log.Fields{"hash": applyStateTx.Hash()}).Debug("Apply state tx sent to eth bridge contract")

		r.log.WithField("hash", applyStateTx.Hash()).Debug("Waiting for apply state tx to be mined")
		applyStateTxReceipt, err := bind.WaitMined(context.Background(), r.ethClient, applyStateTx)

		if err != nil {
			r.log.WithError(err).WithField("hash", applyStateTx.Hash()).Panic("WaitMined apply state tx failed")
		}
		// Tx failed -> status == 0
		if applyStateTxReceipt.Status == 0 {
			r.log.WithField("hash", applyStateTx.Hash()).Panic("Apply state tx failed execution")
		}
		r.log.WithField("hash", applyStateTx.Hash()).Info("Apply state tx mined")
	}
}

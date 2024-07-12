package to_tara

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"relayer/bindings/BridgeBase"
	"relayer/internal/state"
	"runtime/debug"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/sirupsen/logrus"
)

func (r *Relayer) finalize() {
	shouldFinalize, err := r.ethBridge.ShouldFinalizeEpoch(nil)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to call ShouldFinalizeEpoch")
	}
	r.log.WithField("shouldFinalize", shouldFinalize).Debug("ShouldFinalizeEpoch")
	trx, err := r.ethBridge.FinalizeEpoch(r.ethAuth)
	if err != nil {
		r.log.WithField("trx", trx).WithError(err).Debug("Failed to call finalize")
		return
	}
	r.log.WithField("hash", trx.Hash()).Debug("FinalizeEpoch trx sent")

	receipt, err := bind.WaitMined(context.Background(), r.ethClient, trx)
	if err != nil {
		r.log.WithError(err).Warn("Failed to wait for finalize")
		return
	}
	if receipt.Status != 1 {
		r.log.WithField("status", receipt.Status).Fatal("Finalize failed")
		return
	}
	r.log.WithField("block", receipt.BlockNumber.Uint64()).Info("Finalized bridge on block")
}

func (r *Relayer) getDecodedProofs(epoch, block *big.Int) (accountProof [][]byte, rootProof [][]byte, err error) {
	key, err := r.ethClientContract.BridgeRootKeyByEpoch(nil, epoch)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to get bridge root key")
	}
	strKey := "0x" + hex.EncodeToString(key[:])

	r.log.WithFields(logrus.Fields{"epoch": epoch, "block": block}).Debug("Getting proof")

	client := gethclient.New(r.ethClient.Client())

	root, err := client.GetProof(context.Background(), r.bridgeContractAddr, []string{strKey}, block)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to get proof")
	}
	r.log.WithField("root", root).Debug("Got proof")
	if len(root.StorageProof) != 1 {
		r.log.WithField("len", len(root.StorageProof)).Fatal("Invalid storage proof length")
	}
	if root.StorageProof[0].Value.Cmp(big.NewInt(0)) == 0 {
		err = errors.New("no value for epoch")
		return
	}

	accountProof, err = decodeProofs(root.AccountProof)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to decode account proof")
	}

	rootProof, err = decodeProofs(root.StorageProof[0].Proof)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to decode storage proof")
	}
	return
}

func (r *Relayer) processBridgeRoots() {
	lastClientEpoch, err := r.ethClientContract.LastEpoch(nil)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to get last epoch")
	}
	ethFinalizedEpoch, err := r.ethBridge.FinalizedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to get finalized epoch")
	}
	if lastClientEpoch.Cmp(ethFinalizedEpoch) == 0 {
		r.log.WithFields(logrus.Fields{"lastClientEpoch": lastClientEpoch, "ethFinalizedEpoch": ethFinalizedEpoch}).Debug("No new bridge roots to process")
		return
	}
	finalizedBlock, err := r.beaconLightClient.BlockNumber(nil)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to get finalized block")
	}
	epoch := big.NewInt(0).Add(lastClientEpoch, big.NewInt(1))
	for ; epoch.Cmp(ethFinalizedEpoch) <= 0; epoch.Add(epoch, big.NewInt(1)) {
		accountProof, rootProof, err := r.getDecodedProofs(epoch, finalizedBlock)
		if err != nil {
			r.log.WithFields(logrus.Fields{"epoch": epoch}).WithError(err).Warn("Failed to get proofs")
			return
		}

		trx, err := r.ethClientContract.ProcessBridgeRoot(r.taraAuth, accountProof, rootProof)
		if err != nil {
			r.log.WithError(err).Fatal("Failed to call ProcessBridgeRoot")
		}

		r.log.WithField("hash", trx.Hash()).Debug("ProcessBridgeRoot trx sent")

		_, err = bind.WaitMined(context.Background(), r.taraxaClient, trx)
		if err != nil {
			r.log.Fatalf("Failed to wait for ProcessBridgeRoot: %v", err)
		}

		r.log.WithField("hash", trx.Hash()).Info("ProcessBridgeRoot trx mined")
	}
}

func (r *Relayer) BlockByNumber(ctx context.Context, blockNum *big.Int) (*big.Int, error) {
	return r.beaconLightClient.BlockNumber(nil)
}

func (r *Relayer) GetStateWithProof(opts *bind.CallOpts) (BridgeBase.SharedStructsStateWithProof, error) {
	return r.ethBridge.GetStateWithProof(opts)
}

func (r *Relayer) FinalizationInterval() *big.Int {
	finalizationInterval, _ := r.ethBridge.FinalizationInterval(nil)
	return finalizationInterval
}

func (r *Relayer) applyStates() {
	lastClientEpoch, err := r.ethClientContract.LastEpoch(nil)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to get finalized epoch")
	}
	lastAppliedEpoch, err := r.taraBridge.AppliedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Fatal("Failed to get last applied epoch")
	}

	if lastAppliedEpoch.Cmp(lastClientEpoch) == 0 {
		r.log.WithFields(logrus.Fields{"lastAppliedEpoch": lastAppliedEpoch, "lastClientEpoch": lastClientEpoch}).Debug("No new states to process")
		return
	}
	epoch := big.NewInt(0).Add(lastAppliedEpoch, big.NewInt(1))
	for ; epoch.Cmp(lastClientEpoch) <= 0; epoch.Add(epoch, big.NewInt(1)) {
		state, err := state.GetStateWithProof(r, r.log, epoch, nil)
		if err != nil {
			r.log.WithError(err).Fatal("Failed to get state with proof")
		}

		trx, err := r.taraBridge.ApplyState(r.taraAuth, *state)
		if err != nil {
			debug.PrintStack()
			r.log.WithError(err).Fatal("Failed to apply state")
		}
		r.log.WithField("hash", trx.Hash()).Debug("Apply state trx sent")

		_, err = bind.WaitMined(context.Background(), r.taraxaClient, trx)
		if err != nil {
			r.log.WithError(err).Fatal("Failed to wait for apply state trx")
		}

		r.log.WithField("hash", trx.Hash()).Info("Apply state trx mined")
	}
}

func decodeProofs(hexStrings []string) ([][]byte, error) {
	decodedBytes := make([][]byte, len(hexStrings))
	for i, proof := range hexStrings {
		// Check for '0x' prefix and remove it if present
		cleanProof := strings.TrimPrefix(proof, "0x")
		data, err := hex.DecodeString(cleanProof)
		if err != nil {
			return nil, errors.New("failed to decode proof at index " + fmt.Sprint(i) + ": " + err.Error())
		}
		decodedBytes[i] = data
	}
	return decodedBytes, nil
}

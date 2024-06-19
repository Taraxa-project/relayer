package to_tara

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func (r *Relayer) finalize() {
	trx, err := r.ethBridge.FinalizeEpoch(r.ethAuth)
	if err != nil {
		r.log.WithField("trx", trx).WithError(err).Debug("Failed to call finalize")
		return
	}
	receipt, err := bind.WaitMined(context.Background(), r.ethClient, trx)
	if err != nil {
		r.log.WithError(err).Warn("Failed to wait for finalize")
		return
	}
	r.log.WithField("block", receipt.BlockNumber.Uint64()).Info("Finalized bridge on block")
}

func (r *Relayer) getProofs(finalizedBlock *big.Int) (accountProof, epochProof, bridgeRootProof [][]byte) {
	client := gethclient.New(r.ethClient.Client())

	bridgeRootIdx := 0
	epochIdx := 1
	keys := []string{r.epochKey, r.bridgeRootKey}
	root, err := client.GetProof(context.Background(), r.bridgeContractAddr, keys, finalizedBlock)
	if err != nil {
		r.log.Fatalf("Failed to get proof: %v", err)
	}
	if len(root.StorageProof) != 1 {
		r.log.Fatalf("Invalid storage proof length: %d", len(root.StorageProof))
	}

	accountProof, err = decodeProofs(root.AccountProof)
	if err != nil {
		r.log.Fatalf("Failed to decode account proof: %v", err)
	}

	epochProof, err = decodeProofs(root.StorageProof[epochIdx].Proof)
	if err != nil {
		r.log.Fatalf("Failed to decode epoch proof: %v", err)
	}

	bridgeRootProof, err = decodeProofs(root.StorageProof[bridgeRootIdx].Proof)
	if err != nil {
		r.log.Fatalf("Failed to decode bridgeRoot proof: %v", err)
	}

	return
}

func (r *Relayer) ProcessHeaderWithProofs(epoch int64, blockNumber uint64) error {
	updateData, err := r.GetBeaconBlockData(epoch)
	if blockNumber > updateData.FinalizedHeader.Execution.BlockNumber {
		return fmt.Errorf("block number %d is greater than the block number in the finalized header %d", blockNumber, updateData.FinalizedHeader.Execution.BlockNumber)
	}
	// print(*updateData)
	if err != nil {
		return fmt.Errorf("failed to get beacon block data: %v", err)
	}
	finalizedBlock := big.NewInt(0).SetUint64(updateData.FinalizedHeader.Execution.BlockNumber)

	accountProof, epochProof, storageProof := r.getProofs(finalizedBlock)

	trx, err := r.ethClientContract.ProcessHeaderWithProofs(r.taraAuth, *updateData, accountProof, epochProof, storageProof)
	if err != nil {
		r.log.Fatalf("Failed to call ProcessHeaderWithProofs: %v", err)
	}

	r.log.Println("ProcessHeaderWithProofs trx: ", trx.Hash().Hex())

	_, err = bind.WaitMined(context.Background(), r.taraxaClient, trx)
	if err != nil {
		r.log.Fatalf("Failed to wait for ProcessHeaderWithProofs: %v", err)
	}
	return nil
}

func (r *Relayer) applyState(finalizedBlock uint64) {
	opts := bind.CallOpts{}
	opts.BlockNumber.SetUint64(finalizedBlock)

	proof, err := r.ethBridge.GetStateWithProof(&opts)
	if err != nil {
		r.log.Fatalf("Failed to get state with proof: %v", err)
	}
	trx, err := r.taraBridge.ApplyState(r.taraAuth, proof)
	if err != nil {
		r.log.Fatalf("Failed to apply state: %v", err)
	}
	r.log.Printf("Apply state trx: %v", trx.Hash().Hex())
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

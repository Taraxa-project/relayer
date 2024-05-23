package to_tara

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func (r *Relayer) finalize() {
	trx, err := r.ethBridge.FinalizeEpoch(r.ethAuth)
	if err != nil {
		log.WithField("trx", trx).WithError(err).Info("Failed to call finalize")
		return
	}
	receipt, err := bind.WaitMined(context.Background(), r.ethClient, trx)
	if err != nil {
		log.Println("Failed to wait for finalize:", err)
		return
	}
	log.Printf("Finalized bridge on block: %d", receipt.BlockNumber.Uint64())
}

func (r *Relayer) getProof(finalizedBlock *big.Int) {
	key, err := r.ethClientContract.BridgeRootKey(nil)
	if err != nil {
		log.Fatalf("Failed to get bridge root key: %v", err)
	}
	strKey := "0x" + hex.EncodeToString(key[:])

	log.Printf("Bridge root key: %s and block %s", strKey, finalizedBlock.String())

	client := gethclient.New(r.ethClient.Client())

	root, err := client.GetProof(context.Background(), r.bridgeContractAddr, []string{strKey}, finalizedBlock)
	if err != nil {
		log.Fatalf("Failed to get proof: %v", err)
	}
	if len(root.StorageProof) != 1 {
		log.Fatalf("Invalid storage proof length: %d", len(root.StorageProof))
	}

	// log.Printf("Root: %v", root)

	accountProof, err := decodeProofs(root.AccountProof)
	if err != nil {
		log.Fatalf("Failed to decode account proof: %v", err)
	}

	storageProof, err := decodeProofs(root.StorageProof[0].Proof)
	if err != nil {
		log.Fatalf("Failed to decode storage proof: %v", err)
	}

	trx, err := r.ethClientContract.ProcessBridgeRoot(r.taraAuth, accountProof, storageProof)
	if err != nil {
		log.Fatalf("Failed to call ProcessBridgeRoot: %v", err)
	}

	log.Println("ProcessBridgeRoot trx: ", trx.Hash().Hex())

	_, err = bind.WaitMined(context.Background(), r.taraxaClient, trx)
	if err != nil {
		log.Fatalf("Failed to wait for ProcessBridgeRoot: %v", err)
	}
}

func (r *Relayer) applyState(finalizedBlock *big.Int) {
	opts := bind.CallOpts{}
	if finalizedBlock != nil {
		opts.BlockNumber = finalizedBlock
	}
	proof, err := r.ethBridge.GetStateWithProof(&opts)
	if err != nil {
		log.Fatalf("Failed to get state with proof: %v", err)
	}
	trx, err := r.taraBridge.ApplyState(r.taraAuth, proof)
	if err != nil {
		log.Fatalf("Failed to apply state: %v", err)
	}
	log.Printf("Apply state trx: %v", trx.Hash().Hex())
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

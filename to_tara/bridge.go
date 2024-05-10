package to_tara

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func (r *Relayer) finalize() {
	trx, err := r.ethBridge.FinalizeEpoch(r.taraAuth)
	if err != nil {
		log.Fatalf("Failed to call finalize: %v", err)
	}
	receipt, err := bind.WaitMined(context.Background(), r.ethClient, trx)
	if err != nil {
		log.Fatalf("Failed to wait for finalize: %v", err)
	}
	log.Printf("Receipt: %v", receipt)
	r.finalizedBlock = receipt.BlockNumber
	r.onFinalizedBlockNumber <- receipt.BlockNumber.Uint64()
}

func (r *Relayer) getProof() {
	key, err := r.ethClientContract.BridgeRootKey(nil)
	if err != nil {
		log.Fatalf("Failed to get bridge root: %v", err)
	}

	client := gethclient.New(r.ethClient.Client())

	root, err := client.GetProof(context.Background(), r.bridgeContractAddr, []string{string(key[:])}, r.finalizedBlock)
	if err != nil {
		log.Fatalf("Failed to get proof: %v", err)
	}
	if len(root.StorageProof) != 1 {
		log.Fatalf("Invalid storage proof length: %d", len(root.StorageProof))
		return
	}

	accountProof := make([][]byte, len(root.AccountProof))
	for i, proof := range root.AccountProof {
		accountProof[i] = []byte(proof)
	}

	storageProof := make([][]byte, len(root.StorageProof[0].Proof))
	for i, proof := range root.StorageProof[0].Proof {
		storageProof[i] = []byte(proof)
	}

	trx, err := r.ethClientContract.ProcessBridgeRoot(r.taraAuth, accountProof, storageProof)
	if err != nil {
		log.Fatalf("Failed to call ProcessBridgeRoot: %v", err)
		return
	}
	_, err = bind.WaitMined(context.Background(), r.ethClient, trx)
	if err != nil {
		log.Fatalf("Failed to wait for ProcessBridgeRoot: %v", err)
	}
}

func (r *Relayer) applyState() {
	proof, err := r.ethBridge.GetStateWithProof(nil)
	if err != nil {
		log.Fatalf("Failed to get state with proof: %v", err)
	}

	trx, err := r.taraBridge.ApplyState(r.taraAuth, proof)
	if err != nil {
		log.Fatalf("Failed to apply state: %v", err)
	}
	log.Printf("Apply state trx: %v", trx.Hash().Hex())
}

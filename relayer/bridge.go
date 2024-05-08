package relayer

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

func (r *Relayer) callFinalize() *types.Transaction {
	trx, err := r.ethBridge.FinalizeEpoch(r.taraAuth)
	if err != nil {
		log.Fatalf("Failed to call finalize: %v", err)
		return nil
	}
	return trx
}

func (r *Relayer) waitForFinalize(trx *types.Transaction) {
	go func() {
		for {
			receipt, err := r.ethClient.TransactionReceipt(r.taraAuth.Context, trx.Hash())
			if err == ethereum.NotFound {
				// Retry after a delay
				time.Sleep(5 * time.Second) // Adjust the delay as needed
				continue
			} else if err != nil {
				log.Fatalf("Failed to get receipt: %v", err)
				return
			}
			log.Printf("Receipt: %v", receipt)
			r.finalizedBlock = receipt.BlockNumber
			break // Break out of the loop if receipt is found
		}
	}()
}

func (r *Relayer) getProof() {
	// root, err := r.taraBridge.bridgeRootKey(nil)
	// if err != nil {
	// 	log.Fatalf("Failed to get bridge root: %v", err)
	// }

	client := gethclient.New(r.ethClient.Client())

	_, err := client.GetProof(context.Background(), r.bridgeContractAddr, []string{"0x0000000000000000000000000000000000000000000000000000000000000006"}, r.finalizedBlock)
	if err != nil {
		log.Fatalf("Failed to get proof: %v", err)
	}
	//processBridgeRoot
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

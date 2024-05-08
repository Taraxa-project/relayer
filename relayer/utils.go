package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const SLOTS_PER_EPOCH = 32
const EPOCHS_PER_SYNC_COMMITTEE_PERIOD = 256

// GetSlotFromEpoch calculates the first slot number of a given epoch.
func GetSlotFromEpoch(epoch int64) int64 {
	return epoch * SLOTS_PER_EPOCH
}

func GetPeriodFromEpoch(epoch int64) int64 {
	return epoch / EPOCHS_PER_SYNC_COMMITTEE_PERIOD
}

func connectToChain(ctx context.Context, url string, key *ecdsa.PrivateKey) (*ethclient.Client, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to node: %v", err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to retrieve chain ID: %v", err)
	}

	log.Printf("Connected to chain id: %d, on %s", chainID, url)

	// Prepare Taraxa transact options
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	return client, auth, nil
}

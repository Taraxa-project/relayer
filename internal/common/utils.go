package common

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const SLOTS_PER_EPOCH = 32
const EPOCHS_PER_SYNC_COMMITTEE_PERIOD = 256

// GetSlotFromEpoch calculates the first slot number of a given epoch.
func GetSlotFromEpoch(epoch int64) int64 {
	return epoch * SLOTS_PER_EPOCH
}

func GetPeriodFromSlot(slot int64) int64 {
	return slot / EPOCHS_PER_SYNC_COMMITTEE_PERIOD / SLOTS_PER_EPOCH
}

func GetPeriodFromEpoch(epoch int64) int64 {
	return epoch / EPOCHS_PER_SYNC_COMMITTEE_PERIOD
}

func ConnectToChain(ctx context.Context, url string, key *ecdsa.PrivateKey) (*ethclient.Client, *bind.TransactOpts, error) {
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

type Clients struct {
	TaraxaClient *ethclient.Client
	TaraxaAuth   *bind.TransactOpts
	EthClient    *ethclient.Client
	EthAuth      *bind.TransactOpts
}

func CreateClients(ctx context.Context, taraUrl, ethUrl string, key *ecdsa.PrivateKey) (*Clients, error) {
	taraxaClient, taraAuth, err := ConnectToChain(ctx, taraUrl, key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Taraxa network: %v", err)
	}

	ethClient, ethAuth, err := ConnectToChain(ctx, ethUrl, key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ETH network: %v", err)
	}
	return &Clients{
		TaraxaClient: taraxaClient,
		TaraxaAuth:   taraAuth,
		EthClient:    ethClient,
		EthAuth:      ethAuth,
	}, nil
}

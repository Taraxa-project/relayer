package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"relayer/BeaconLightClient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RelayerConfig struct {
	BeaconNodeEndpoint string
	TaraxaNodeURL      string
	TaraxaContractAddr common.Address
	Key                *ecdsa.PrivateKey
	LightNodeEndpoint  string
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	beaconNodeEndpoint string
	lightNodeEndpoint  string
	taraxaContract     common.Address
	taraxaClient       *ethclient.Client
	beaconLightClient  *BeaconLightClient.BeaconLightClient
	auth               *bind.TransactOpts
	onFinalizedEpoch   chan int64
	currentPeriod      int64
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *RelayerConfig) (*Relayer, error) {
	taraxaClient, err := ethclient.Dial(cfg.TaraxaNodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Taraxa node: %v", err)
	}

	chainID, err := taraxaClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chain ID: %v", err)
	}

	log.Printf("Connected to chain id: %d, on %s", chainID, cfg.TaraxaNodeURL)

	// Prepare transact options
	auth, err := bind.NewKeyedTransactorWithChainID(cfg.Key, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	beaconLightClient, err := BeaconLightClient.NewBeaconLightClient(cfg.TaraxaContractAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	return &Relayer{
		beaconNodeEndpoint: cfg.BeaconNodeEndpoint,
		taraxaContract:     cfg.TaraxaContractAddr,
		taraxaClient:       taraxaClient,
		beaconLightClient:  beaconLightClient,
		auth:               auth,
		lightNodeEndpoint:  cfg.LightNodeEndpoint,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	go r.startEventProcessing(ctx)
	go r.processNewBlocks(ctx)
}

func (r *Relayer) Close() {
	close(r.onFinalizedEpoch)
}

func (r *Relayer) processNewBlocks(ctx context.Context) {
	for {
		select {
		case epoch := <-r.onFinalizedEpoch:
			log.Printf("Processing new block for epoch: %d", epoch)

			r.UpdateLightClient(epoch, r.currentPeriod != GetPeriodFromEpoch(epoch))
			if r.currentPeriod != GetPeriodFromEpoch(epoch) {
				r.currentPeriod = GetPeriodFromEpoch(epoch)
			}
		case <-ctx.Done():
			log.Println("Stopping new block processing")
			return
		}
	}
}

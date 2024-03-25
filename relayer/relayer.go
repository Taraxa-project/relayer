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

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	beaconNodeEndpoint string
	lightNodeEndpoint  string
	taraxaContract     common.Address
	taraxaClient       *ethclient.Client
	beaconLightClient  *BeaconLightClient.BeaconLightClient
	auth               *bind.TransactOpts
	onFinalizedEpoch   chan int64
}

// NewRelayer creates a new Relayer instance
func NewRelayer(beaconNodeEndpoint, taraxaNodeURL string, taraxaContractAddr common.Address, key *ecdsa.PrivateKey) (*Relayer, error) {
	taraxaClient, err := ethclient.Dial(taraxaNodeURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Taraxa node: %v", err)
	}

	chainID, err := taraxaClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chain ID: %v", err)
	}

	log.Printf("Connected to chain id: %d, on %s", chainID, taraxaNodeURL)

	// Prepare transact options
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	beaconLightClient, err := BeaconLightClient.NewBeaconLightClient(taraxaContractAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	return &Relayer{
		beaconNodeEndpoint: beaconNodeEndpoint,
		taraxaContract:     taraxaContractAddr,
		taraxaClient:       taraxaClient,
		beaconLightClient:  beaconLightClient,
		auth:               auth,
		lightNodeEndpoint:  "https://www.lightclientdata.org",
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

			slot := GetSlotFromEpoch(epoch)
			r.updateNewHeader(slot)
			// r.updateCommittee(slot)
			// Add logic to process the block here
		case <-ctx.Done():
			log.Println("Stopping new block processing")
			return
		}
	}
}

func (r *Relayer) updateCommittee(slot int64) {
	if slot%16384 == 0 {
		log.Println("Updating sync committee")
		//TODO implement
	}
}

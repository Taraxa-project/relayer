package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"relayer/BeaconLightClient"

	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RelayerConfig struct {
	BeaconNodeEndpoint  string
	TaraxaRPCURL        string
	EthRPCURL           string
	TaraxaEthClientAddr common.Address
	TaraxaBridgeAddr    common.Address
	EthBridgeAddr       common.Address
	Key                 *ecdsa.PrivateKey
	LightNodeEndpoint   string
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	beaconNodeEndpoint string
	lightNodeEndpoint  string
	taraxaClient       *ethclient.Client
	taraAuth           *bind.TransactOpts
	ethClient          *ethclient.Client
	ethAuth            *bind.TransactOpts
	beaconLightClient  *BeaconLightClient.BeaconLightClient
	ethBridge          *bridge_contract_interface.BridgeContractInterface
	taraBridge         *bridge_contract_interface.BridgeContractInterface
	onFinalizedEpoch   chan int64
	currentPeriod      int64
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *RelayerConfig) (*Relayer, error) {
	taraxaClient, taraAuth, err := connectToChain(context.Background(), cfg.TaraxaRPCURL, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Taraxa network: %v", err)
	}

	ethClient, ethAuth, err := connectToChain(context.Background(), cfg.EthRPCURL, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ETH network: %v", err)
	}

	beaconLightClient, err := BeaconLightClient.NewBeaconLightClient(cfg.TaraxaEthClientAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	ethBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.TaraxaBridgeAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	taraBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.EthBridgeAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	return &Relayer{
		beaconNodeEndpoint: cfg.BeaconNodeEndpoint,
		taraxaClient:       taraxaClient,
		taraAuth:           taraAuth,
		ethClient:          ethClient,
		ethAuth:            ethAuth,
		beaconLightClient:  beaconLightClient,
		ethBridge:          ethBridge,
		taraBridge:         taraBridge,
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

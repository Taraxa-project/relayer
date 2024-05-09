package to_eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"relayer/BeaconLightClient"
	"relayer/common"

	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RelayerConfig struct {
	BeaconNodeEndpoint  string
	TaraxaRPCURL        string
	EthRPCURL           string
	TaraxaEthClientAddr eth_common.Address
	TaraxaBridgeAddr    eth_common.Address
	EthBridgeAddr       eth_common.Address
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
	finalizedBlock     *big.Int
	bridgeContractAddr eth_common.Address
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *RelayerConfig) (*Relayer, error) {
	taraxaClient, taraAuth, err := common.ConnectToChain(context.Background(), cfg.TaraxaRPCURL, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Taraxa network: %v", err)
	}

	ethClient, ethAuth, err := common.ConnectToChain(context.Background(), cfg.EthRPCURL, cfg.Key)
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
		bridgeContractAddr: cfg.EthBridgeAddr,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	// go r.startEventProcessing(ctx)
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

			// r.UpdateLightClient(epoch, r.currentPeriod != common.GetPeriodFromEpoch(epoch))
			// if r.currentPeriod != common.GetPeriodFromEpoch(epoch) {
			// 	r.currentPeriod = common.GetPeriodFromEpoch(epoch)
			// }
		case <-ctx.Done():
			log.Println("Stopping new block processing")
			return
		}
	}
}

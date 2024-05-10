package to_eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"relayer/common"

	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	tara_client_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client/contract_interface"
	tara_rpc_types "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	BeaconNodeEndpoint    string
	TaraxaRPCURL          string
	EthRPCURL             string
	TaraxaClientOnEthAddr eth_common.Address
	TaraxaBridgeAddr      eth_common.Address
	EthBridgeAddr         eth_common.Address
	Key                   *ecdsa.PrivateKey
	LightNodeEndpoint     string
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	beaconNodeEndpoint string
	lightNodeEndpoint  string
	taraxaClient       *TaraxaClientWrapper
	taraxaNodeConfig   *tara_rpc_types.TaraConfig
	taraAuth           *bind.TransactOpts
	ethClient          *ethclient.Client
	ethAuth            *bind.TransactOpts
	ethBridge          *bridge_contract_interface.BridgeContractInterface
	taraBridge         *bridge_contract_interface.BridgeContractInterface
	taraClientOnEth    *tara_client_interface.TaraClientContractInterface
	onFinalizedEpoch   chan int64
	bridgeContractAddr eth_common.Address
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *Config) (*Relayer, error) {
	tcl, taraAuth, err := common.ConnectToChain(context.Background(), cfg.TaraxaRPCURL, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Taraxa network: %v", err)
	}
	taraxaClient := NewClient(tcl)
	taraConfig, err := taraxaClient.GetTaraConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get Taraxa Config: %v", err)
	}

	ethClient, ethAuth, err := common.ConnectToChain(context.Background(), cfg.EthRPCURL, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ETH network: %v", err)
	}

	taraClientOnEth, err := tara_client_interface.NewTaraClientContractInterface(cfg.TaraxaClientOnEthAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	ethBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.TaraxaBridgeAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the EthBridge contract: %v", err)
	}

	taraBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.EthBridgeAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the TaraBridge contract: %v", err)
	}

	return &Relayer{
		beaconNodeEndpoint: cfg.BeaconNodeEndpoint,
		taraxaClient:       taraxaClient,
		taraxaNodeConfig:   taraConfig,
		taraAuth:           taraAuth,
		ethClient:          ethClient,
		ethAuth:            ethAuth,
		taraClientOnEth:    taraClientOnEth,
		ethBridge:          ethBridge,
		taraBridge:         taraBridge,
		lightNodeEndpoint:  cfg.LightNodeEndpoint,
		bridgeContractAddr: cfg.EthBridgeAddr,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	// go r.startEventProcessing(ctx)
	go r.ProcessPillarBlocks(ctx)
}

func (r *Relayer) Close() {
	close(r.onFinalizedEpoch)
}

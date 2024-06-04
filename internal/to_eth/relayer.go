package to_eth

import (
	"context"
	"fmt"
	"relayer/internal/common"

	log "github.com/sirupsen/logrus"

	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	tara_client_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/eth/tara_client_contract_client/contract_interface"
	tara_rpc_types "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	TaraxaClientOnEthAddr eth_common.Address
	TaraxaBridgeAddr      eth_common.Address
	EthBridgeAddr         eth_common.Address
	Clients               *common.Clients
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	taraxaClient          *TaraxaClientWrapper
	taraxaNodeConfig      *tara_rpc_types.TaraConfig
	taraAuth              *bind.TransactOpts
	ethClient             *ethclient.Client
	ethAuth               *bind.TransactOpts
	ethBridge             *bridge_contract_interface.BridgeContractInterface
	taraBridge            *bridge_contract_interface.BridgeContractInterface
	taraClientOnEth       *tara_client_interface.TaraClientContractInterface
	onFinalizedEpoch      chan int64
	bridgeContractAddr    eth_common.Address
	lastAppliedBridgeRoot eth_common.Hash
	log                   *log.Entry
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *Config) (*Relayer, error) {
	relayerLogger := log.WithField("relayer", "to_eth")

	taraxaClient := NewClient(cfg.Clients.TaraxaClient)
	taraConfig, err := taraxaClient.GetTaraConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get Taraxa Config: %v", err)
	}

	taraClientOnEth, err := tara_client_interface.NewTaraClientContractInterface(cfg.TaraxaClientOnEthAddr, cfg.Clients.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	taraBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.TaraxaBridgeAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the TaraBridge contract: %v", err)
	}

	ethBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.EthBridgeAddr, cfg.Clients.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the EthBridge contract: %v", err)
	}

	return &Relayer{
		taraxaClient:       taraxaClient,
		taraxaNodeConfig:   taraConfig,
		taraAuth:           cfg.Clients.TaraxaAuth,
		ethClient:          cfg.Clients.EthClient,
		ethAuth:            cfg.Clients.EthAuth,
		taraClientOnEth:    taraClientOnEth,
		ethBridge:          ethBridge,
		taraBridge:         taraBridge,
		bridgeContractAddr: cfg.EthBridgeAddr,
		log:                relayerLogger,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	br, err := r.taraClientOnEth.GetFinalizedBridgeRoot(nil)
	if err != nil {
		r.log.WithError(err).Error("Failed to get last applied bridge root")
	}
	r.lastAppliedBridgeRoot = br
	// sync
	r.processPillarBlocks()

	go r.ListenForPillarBlockUpdates(ctx)
}

func (r *Relayer) Close() {
	close(r.onFinalizedEpoch)
}

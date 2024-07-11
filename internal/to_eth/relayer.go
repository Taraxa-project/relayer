package to_eth

import (
	"context"
	"fmt"
	"math/big"
	"path/filepath"
	"relayer/bindings/BridgeBase"
	"relayer/bindings/TaraClient"
	"relayer/internal/common"
	"relayer/internal/logging"
	"relayer/internal/types"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	TaraxaClientOnEthAddr eth_common.Address
	TaraxaBridgeAddr      eth_common.Address
	EthBridgeAddr         eth_common.Address
	Clients               *common.Clients
	DataDir               string
	LogLevel              string
	PillarBlocksInBatch   int
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	taraxaClient        *TaraxaClientWrapper
	taraxaNodeConfig    *types.TaraConfig
	ethClient           *ethclient.Client
	ethAuth             *bind.TransactOpts
	ethBridge           *BridgeBase.BridgeBase
	taraBridge          *BridgeBase.BridgeBase
	taraClientOnEth     *TaraClient.TaraClient
	onFinalizedEpoch    chan int64
	bridgeContractAddr  eth_common.Address
	latestBridgeRoot    eth_common.Hash
	latestClientEpoch   *big.Int
	latestAppliedEpoch  *big.Int
	log                 *log.Logger
	pillarBlocksInBatch int
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *Config) (*Relayer, error) {
	relayerLogger := logging.MakeLogger("to_eth", filepath.Join(cfg.DataDir, "logs", "to_eth.log"), cfg.LogLevel)

	taraxaClient := NewClient(cfg.Clients.TaraxaClient)
	taraConfig, err := taraxaClient.GetTaraConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get Taraxa Config: %v", err)
	}

	taraClientOnEth, err := TaraClient.NewTaraClient(cfg.TaraxaClientOnEthAddr, cfg.Clients.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	taraBridge, err := BridgeBase.NewBridgeBase(cfg.TaraxaBridgeAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the TaraBridge contract: %v", err)
	}

	ethBridge, err := BridgeBase.NewBridgeBase(cfg.EthBridgeAddr, cfg.Clients.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the EthBridge contract: %v", err)
	}

	return &Relayer{
		taraxaClient:        taraxaClient,
		taraxaNodeConfig:    taraConfig,
		ethClient:           cfg.Clients.EthClient,
		ethAuth:             cfg.Clients.EthAuth,
		taraClientOnEth:     taraClientOnEth,
		ethBridge:           ethBridge,
		taraBridge:          taraBridge,
		bridgeContractAddr:  cfg.EthBridgeAddr,
		log:                 relayerLogger,
		pillarBlocksInBatch: cfg.PillarBlocksInBatch,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	finalized_block, err := r.taraClientOnEth.GetFinalized(nil)
	if err != nil {
		r.log.WithError(err).Error("Failed to get finalized block")
	}
	r.latestBridgeRoot = finalized_block.Block.BridgeRoot
	r.latestClientEpoch = finalized_block.Block.Epoch

	r.latestAppliedEpoch, err = r.ethBridge.AppliedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Error("Failed to get last applied epoch")
	}
	// sync client with pillar blocks
	r.processPillarBlocks()
	// check it in case we missed a state bridging and don't have a new pillar blocks to bridge
	r.bridgeState()

	go r.ListenForPillarBlockUpdates(ctx)
}

func (r *Relayer) Close() {
	close(r.onFinalizedEpoch)
}

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
	"relayer/internal/state"
	"relayer/internal/types"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
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
	ethClient           *common.GasLimitClient
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
	StakeState          *state.State
	ready_to_shutdown   bool
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

	totalWeight, err := taraClientOnEth.TotalWeight(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get total stake: %v", err)
	}
	state := state.NewState(int32(totalWeight.Int64()), func(a eth_common.Address) int32 {
		votes, err := taraClientOnEth.ValidatorVoteCounts(nil, a)
		if err != nil {
			return 0
		}
		return int32(votes.Int64())
	})

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
		StakeState:          state,
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
	r.log.WithFields(log.Fields{"ClientPeriod": finalized_block.Block.Period, "ClientEpoch": r.latestClientEpoch, "BridgeAppliedEpoch": r.latestAppliedEpoch}).Info("Starting relayer")
	// sync client with pillar blocks
	r.processPillarBlocks()
	// check it in case we missed a state bridging and don't have a new pillar blocks to bridge
	r.bridgeState()

	r.ready_to_shutdown = true

	go r.ListenForPillarBlockUpdates(ctx)
	r.log.Info("Relayer started")
}

func (r *Relayer) Shutdown() {
	for !r.ready_to_shutdown {
		time.Sleep(100 * time.Millisecond)
	}
}

func (r *Relayer) SetReadyToShutdown() {
	r.ready_to_shutdown = true
}

package to_tara

import (
	"context"
	"fmt"
	"path/filepath"
	"relayer/bindings/BeaconLightClient"
	"relayer/bindings/BridgeBase"
	"relayer/bindings/EthClient"
	"relayer/internal/common"
	"relayer/internal/logging"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	BeaconNodeEndpoint    string
	BeaconLightClientAddr eth_common.Address
	EthClientOnTaraAddr   eth_common.Address
	TaraxaBridgeAddr      eth_common.Address
	EthBridgeAddr         eth_common.Address
	Clients               *common.Clients
	DataDir               string
	LogLevel              string
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	beaconNodeEndpoint        string
	taraxaClient              *ethclient.Client
	taraAuth                  *bind.TransactOpts
	ethClient                 *common.GasLimitClient
	ethAuth                   *bind.TransactOpts
	beaconLightClient         *BeaconLightClient.BeaconLightClient
	ethClientContract         *EthClient.EthClient
	ethBridge                 *BridgeBase.BridgeBase
	taraBridge                *BridgeBase.BridgeBase
	onFinalizedEpoch          chan int64
	onFinalizedBlockNumber    chan uint64
	onSyncCommitteeUpdate     chan int64
	currentContractSyncPeriod int64
	currentCommittee          *BeaconLightClient.BeaconChainSyncCommittee // This is only used for caching latest sync committee as it lasts ~27h
	currentCommitteePeriod    int64
	bridgeContractAddr        eth_common.Address
	log                       *log.Logger
	started                   bool
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *Config) (*Relayer, error) {
	relayerLogger := logging.MakeLogger("to_tara", filepath.Join(cfg.DataDir, "logs", "to_tara.log"), cfg.LogLevel)

	beaconLightClient, err := BeaconLightClient.NewBeaconLightClient(cfg.BeaconLightClientAddr, cfg.Clients.TaraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	taraBridge, err := BridgeBase.NewBridgeBase(cfg.TaraxaBridgeAddr, cfg.Clients.TaraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the TaraBridge contract: %v", err)
	}

	ethBridge, err := BridgeBase.NewBridgeBase(cfg.EthBridgeAddr, cfg.Clients.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the EthBridge contract: %v", err)
	}

	ethClientContract, err := EthClient.NewEthClient(cfg.EthClientOnTaraAddr, cfg.Clients.TaraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the ethClientContract contract: %v", err)
	}

	return &Relayer{
		beaconNodeEndpoint: cfg.BeaconNodeEndpoint,
		taraxaClient:       cfg.Clients.TaraxaClient,
		taraAuth:           cfg.Clients.TaraxaAuth,
		ethClient:          cfg.Clients.EthClient,
		ethAuth:            cfg.Clients.EthAuth,
		beaconLightClient:  beaconLightClient,
		ethClientContract:  ethClientContract,
		ethBridge:          ethBridge,
		taraBridge:         taraBridge,
		bridgeContractAddr: cfg.EthBridgeAddr,
		log:                relayerLogger,
		currentCommittee:   nil,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	r.onFinalizedBlockNumber = make(chan uint64)
	r.onSyncCommitteeUpdate = make(chan int64)

	slot, err := r.beaconLightClient.Slot(nil)
	if err != nil {
		r.log.WithError(err).Panic("Failed to get current slot from contract")
	}

	r.currentContractSyncPeriod = common.GetPeriodFromSlot(int64(slot))
	r.log.WithField("BeaconPeriod", r.currentContractSyncPeriod).Info("Starting relayer")

	r.processBridgeRoots()
	r.applyStates()

	r.checkAndFinalize()
	r.checkAndUpdateNextSyncCommittee(r.currentContractSyncPeriod)

	r.started = true

	go r.startEventProcessing(ctx)
	go r.processNewBlocks(ctx)
}

func (r *Relayer) Shutdown() {
	for !r.started {
		time.Sleep(100 * time.Millisecond)
	}
}

func (r *Relayer) processNewBlocks(ctx context.Context) {
	var finalizedBlockNumber uint64
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case epoch := <-r.onFinalizedEpoch:
			r.log.WithField("epoch", epoch).Trace("Processing new block for epoch")
			if r.currentContractSyncPeriod < common.GetPeriodFromEpoch(epoch-3) { // -3 so we have new period finalized :)
				go r.checkAndUpdateNextSyncCommittee(common.GetPeriodFromEpoch(epoch))
			}
			if finalizedBlockNumber != 0 {
				r.log.WithFields(log.Fields{"epoch": epoch, "block": finalizedBlockNumber}).Debug("Updating Beacon Light Client")
				_, err := r.updateLightClient(epoch, finalizedBlockNumber)
				if err != nil {
					r.log.WithError(err).Info("Did not to update light client")
				} else {
					r.processBridgeRoots()
					r.applyStates()
					finalizedBlockNumber = 0
				}
			}
		case blockNumber := <-r.onFinalizedBlockNumber:
			r.log.WithField("blockNumber", blockNumber).Info("Received finalized block number")
			if finalizedBlockNumber != 0 {
				r.log.WithFields(log.Fields{"finalizedBlockNumber": finalizedBlockNumber, "current": blockNumber}).Info("Finalized block number was not processed yet")
				continue
			}
			finalizedBlockNumber = blockNumber
		case <-ticker.C:
			if finalizedBlockNumber == 0 {
				r.log.Debug("Checking if we need to finalize")
				go r.checkAndFinalize()
			}
		case period := <-r.onSyncCommitteeUpdate:
			if period == 0 {
				continue
			}
			r.log.WithField("period", period).Info("Next sync committee updated")
			r.currentContractSyncPeriod = period - 1
		case <-ctx.Done():
			r.log.Info("Stopping new block processing")
			return
		}
	}
}

func (r *Relayer) checkAndFinalize() {
	r.finalize()
	finalizedEpoch, err := r.ethBridge.FinalizedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Warn("Failed to get finalized epoch from ETH contract")
		return
	}
	appliedEpoch, err := r.taraBridge.AppliedEpoch(nil)
	if err != nil {
		r.log.WithError(err).Warn("Failed to get finalized epoch from TARA contract")
		return
	}

	r.log.WithFields(log.Fields{"finalized": finalizedEpoch, "applied": appliedEpoch}).Trace("Checking if need to finalize")
	if finalizedEpoch.Cmp(appliedEpoch) > 0 {
		r.log.WithFields(log.Fields{"finalized": finalizedEpoch, "applied": appliedEpoch}).Debug("Finalizing ETH epoch in TARA contract")

		lastFinalizedBlock, err := r.ethBridge.LastFinalizedBlock(nil)
		if err != nil {
			r.log.WithError(err).Panic("Failed to get last finalized block")
		}
		r.onFinalizedBlockNumber <- lastFinalizedBlock.Uint64()
	}
}

func (r *Relayer) checkAndUpdateNextSyncCommittee(period int64) {
	root, err := r.beaconLightClient.SyncCommitteeRoots(nil, uint64(period+1))
	if err != nil {
		r.log.WithError(err).Panic("Failed to get sync committee roots")
	}

	if root == [32]byte{} {
		r.updateSyncCommittee(period)
	}
}

package to_tara

import (
	"context"
	"fmt"
	"path/filepath"
	"relayer/bindings/BeaconLightClient"
	"relayer/bindings/BridgeBase"
	"relayer/bindings/EthClient"
	"relayer/internal/logging"
	"relayer/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	BeaconNodeEndpoint    string
	BeaconLightClientAddr common.Address
	EthClientOnTaraAddr   common.Address
	TaraxaBridgeAddr      common.Address
	EthBridgeAddr         common.Address
	Clients               *utils.Clients
	DataDir               string
	LogLevel              string
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	beaconNodeEndpoint        string
	taraxaClient              *ethclient.Client
	taraAuth                  *bind.TransactOpts
	ethClient                 *utils.GasLimitClient
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
	bridgeContractAddr        common.Address
	log                       *log.Logger
	ready_to_shutdown         bool
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

func (r *Relayer) syncCommittee() {
	update, err := r.GetLightClientFinalityUpdate()
	if err != nil {
		r.log.WithError(err).Error("Failed to get light client finality update")
		return
	}

	finSlot := uint64(update.Data.FinalizedHeader.Beacon.Slot)
	finPeriod := utils.GetPeriodFromSlot(int64(finSlot))
	for finPeriod >= r.currentCommitteePeriod {
		r.log.WithFields(log.Fields{"fin_period": finPeriod, "contractPeriod": r.currentContractSyncPeriod}).Fatal("Syncing committee")
		r.updateSyncCommittee(r.currentCommitteePeriod)
	}
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	r.onFinalizedBlockNumber = make(chan uint64)
	r.onSyncCommitteeUpdate = make(chan int64)

	slot, err := r.beaconLightClient.Slot(nil)
	if err != nil {
		r.log.WithError(err).Panic("Failed to get current slot from contract")
	}

	r.currentContractSyncPeriod = utils.GetPeriodFromSlot(int64(slot))
	r.log.WithField("BeaconPeriod", r.currentContractSyncPeriod).Info("Starting relayer")
	r.syncCommittee()

	r.processBridgeRoots()
	r.applyStates()

	go r.startEventProcessing(ctx)
	go r.processNewBlocks(ctx)

	r.checkAndFinalize()
	r.checkAndUpdateNextSyncCommittee(r.currentContractSyncPeriod)

	r.ready_to_shutdown = true
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

func (r *Relayer) processNewBlocks(ctx context.Context) {
	var finalizedBlockNumber uint64
	mainTicker := time.NewTicker(2 * time.Minute)
	defer mainTicker.Stop()

	const finalizationTimeout = 15 * time.Minute
	// Separate ticker for epoch timeout
	epochTimeoutTicker := time.NewTicker(finalizationTimeout)
	defer epochTimeoutTicker.Stop()

	for {
		select {
		case epoch := <-r.onFinalizedEpoch:
			epochTimeoutTicker.Reset(finalizationTimeout) // Reset epoch timeout ticker
			r.log.WithField("epoch", epoch).Trace("Processing new block for epoch")
			if r.currentContractSyncPeriod < utils.GetPeriodFromEpoch(epoch-3) { // -3 so we have new period finalized :)
				go r.checkAndUpdateNextSyncCommittee(utils.GetPeriodFromEpoch(epoch))
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
		case <-mainTicker.C:
			if finalizedBlockNumber == 0 {
				r.log.Debug("Checking if we need to finalize")
				go r.checkAndFinalize()
			}
		case <-epochTimeoutTicker.C:
			// If this ticker fires, it means no epoch was received for 15 minutes
			r.log.Fatal("No finalized epoch received in 15 minutes! Exiting...")
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

package to_tara

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"relayer/bindings/BeaconLightClient"
	"relayer/bindings/EthClient"
	"relayer/internal/common"
	"time"

	log "github.com/sirupsen/logrus"

	bridge_contract_interface "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/bridge_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	BeaconNodeEndpoint    string
	TaraxaRPCURL          string
	EthRPCURL             string
	BeaconLightClientAddr eth_common.Address
	EthClientOnTaraAddr   eth_common.Address
	TaraxaBridgeAddr      eth_common.Address
	EthBridgeAddr         eth_common.Address
	Key                   *ecdsa.PrivateKey
	LightNodeEndpoint     string
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	beaconNodeEndpoint     string
	lightNodeEndpoint      string
	taraxaClient           *ethclient.Client
	taraAuth               *bind.TransactOpts
	ethClient              *ethclient.Client
	ethAuth                *bind.TransactOpts
	beaconLightClient      *BeaconLightClient.BeaconLightClient
	ethClientContract      *EthClient.EthClient
	ethBridge              *bridge_contract_interface.BridgeContractInterface
	taraBridge             *bridge_contract_interface.BridgeContractInterface
	onFinalizedEpoch       chan int64
	onFinalizedBlockNumber chan uint64
	onSyncCommitteeUpdate  chan int64
	currentPeriod          int64
	bridgeContractAddr     eth_common.Address
}

// NewRelayer creates a new Relayer instance
func NewRelayer(cfg *Config) (*Relayer, error) {
	taraxaClient, taraAuth, err := common.ConnectToChain(context.Background(), cfg.TaraxaRPCURL, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Taraxa network: %v", err)
	}

	ethClient, ethAuth, err := common.ConnectToChain(context.Background(), cfg.EthRPCURL, cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ETH network: %v", err)
	}

	beaconLightClient, err := BeaconLightClient.NewBeaconLightClient(cfg.BeaconLightClientAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	taraBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.TaraxaBridgeAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the TaraBridge contract: %v", err)
	}

	ethBridge, err := bridge_contract_interface.NewBridgeContractInterface(cfg.EthBridgeAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the EthBridge contract: %v", err)
	}

	ethClientContract, err := EthClient.NewEthClient(cfg.EthClientOnTaraAddr, taraxaClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the ethClientContract contract: %v", err)
	}

	return &Relayer{
		beaconNodeEndpoint: cfg.BeaconNodeEndpoint,
		taraxaClient:       taraxaClient,
		taraAuth:           taraAuth,
		ethClient:          ethClient,
		ethAuth:            ethAuth,
		beaconLightClient:  beaconLightClient,
		ethClientContract:  ethClientContract,
		ethBridge:          ethBridge,
		taraBridge:         taraBridge,
		lightNodeEndpoint:  cfg.LightNodeEndpoint,
		bridgeContractAddr: cfg.EthBridgeAddr,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	r.onFinalizedBlockNumber = make(chan uint64)
	r.onSyncCommitteeUpdate = make(chan int64)

	slot, err := r.beaconLightClient.Slot(nil)
	if err != nil {
		log.WithError(err).Fatal("Failed to get current slot from contract")
	}

	r.currentPeriod = common.GetPeriodFromSlot(int64(slot))

	root, err := r.beaconLightClient.SyncCommitteeRoots(nil, uint64(r.currentPeriod+1))
	if err != nil {
		log.WithError(err).Fatal("Failed to get sync committee roots")
	}

	log.WithField("current period", r.currentPeriod).Info("Beacon light client deployed, starting relayer")

	if root == [32]byte{} {
		r.updateSyncCommittee(r.currentPeriod)
	}

	go r.startEventProcessing(ctx)
	go r.processNewBlocks(ctx)
	r.checkAndFinalize()
}

func (r *Relayer) Close() {
	close(r.onFinalizedEpoch)
	close(r.onFinalizedBlockNumber)
	close(r.onSyncCommitteeUpdate)
}

func (r *Relayer) processNewBlocks(ctx context.Context) {
	var finalizedBlockNumber uint64
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case epoch := <-r.onFinalizedEpoch:
			log.Printf("Processing new block for epoch: %d", epoch)
			if r.currentPeriod != common.GetPeriodFromEpoch(epoch) {
				r.updateSyncCommittee(common.GetPeriodFromEpoch(epoch))
			}
			if finalizedBlockNumber != 0 {
				log.Println("Updating light client with epoch", epoch, "and block number", finalizedBlockNumber)
				blockNum, err := r.updateLightClient(epoch, 0)
				if err != nil {
					log.Println("Did not to update light client:", err)
				} else {
					r.getProof(blockNum)
					r.applyState(blockNum)
					finalizedBlockNumber = 0
				}
			}
		case blockNumber := <-r.onFinalizedBlockNumber:
			log.Println("Received finalized block number", blockNumber)
			if finalizedBlockNumber != 0 {
				log.Println("Finalized block number was not processed yet, skipping this one")
				continue
			}
			finalizedBlockNumber = blockNumber
		case <-ticker.C:
			log.Println("Checking for if we need to finalize")
			if finalizedBlockNumber == 0 {
				go r.checkAndFinalize()
			}
		case period := <-r.onSyncCommitteeUpdate:
			r.currentPeriod = period
			log.WithField("period", period).Info("Sync committee updated")
		case <-ctx.Done():
			log.Println("Stopping new block processing")
			return
		}
	}
}

func (r *Relayer) checkAndFinalize() {
	r.finalize()
	finalizedEpoch, err := r.ethBridge.FinalizedEpoch(nil)
	if err != nil {
		log.Warningf("Failed to get finalized epoch from ETH contract: %v", err)
		return
	}
	appliedEpoch, err := r.taraBridge.AppliedEpoch(nil)
	if err != nil {
		log.Warningf("Failed to get finalized epoch from TARA contract: %v", err)
		return
	}
	if finalizedEpoch.Cmp(appliedEpoch) > 0 {
		log.Printf("Finalizing ETH epoch %d on TARA epoch %d", finalizedEpoch, appliedEpoch)

		lastFinalizedBlock, err := r.ethBridge.LastFinalizedBlock(nil)
		if err != nil {
			log.Fatalf("Failed to get last finalized block: %v", err)
		}
		r.onFinalizedBlockNumber <- lastFinalizedBlock.Uint64()
	}
}

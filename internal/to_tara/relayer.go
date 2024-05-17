package to_tara

import (
	"context"
	"fmt"
	"time"

	"github.com/Taraxa-project/relayer/bindings/BeaconLightClient"
	"github.com/Taraxa-project/relayer/bindings/EthClient"
	"github.com/Taraxa-project/relayer/internal/common"

	eth_client "github.com/Taraxa-project/relayer/clients/eth"
	tara_client "github.com/Taraxa-project/relayer/clients/tara"
	log "github.com/sirupsen/logrus"
)

type RelayerConfig struct {
	*tara_client.TaraClientConfig
	BeaconNodeEndpoint string `json:"beacon_node_endpoint"`
	LightNodeEndpoint  string `json:"light_node_endpoint"`
}

// Relayer encapsulates the functionality to relay data from Ethereum to Taraxa
type Relayer struct {
	ethClient  *eth_client.EthClient
	taraClient *tara_client.TaraClient

	beaconLightClient *BeaconLightClient.BeaconLightClient // TODO: move to taraClient
	ethClientContract *EthClient.EthClient                 // TODO: move to taraClient

	Config RelayerConfig

	onFinalizedEpoch       chan int64
	onFinalizedBlockNumber chan uint64
	currentPeriod          int64
}

// NewRelayer creates a new Relayer instance
func NewRelayer(taraClient *tara_client.TaraClient, ethClient *eth_client.EthClient, config RelayerConfig) (*Relayer, error) {
	// // Get tara config
	// taraNodeConfig, err := taraClient.RpcClient.GetTaraConfig()
	// if err != nil {
	// 	log.Fatal("GetTaraConfig err: ", err)
	// }
	// if taraClient.Config.ChainID.Uint64() != taraNodeConfig.ChainId {
	// 	log.Fatal("Configured chainID != retrived chain ID ", taraClient.Config.ChainID.Uint64(), taraNodeConfig.ChainId)
	// }
	// log.Println("Configured ficus hardfork number ", uint64(taraNodeConfig.Hardforks.FicusHf.BlockNum), ", pillar blocks interval: ", uint64(taraNodeConfig.Hardforks.FicusHf.PillarBlocksInterval))

	// // TODO: taraClient.EthClient should work too !!!
	// currentBlockNumber, err := taraClient.RpcClient.EthClient.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatal("BlockNumber err: ", err)
	// }
	// log.Println("Current tara block number: ", currentBlockNumber)
	// if currentBlockNumber < uint64(taraNodeConfig.Hardforks.FicusHf.BlockNum) {
	// 	log.Fatal("No need to run relayer yet. Ficus hardfork block num: ", uint64(taraNodeConfig.Hardforks.FicusHf.BlockNum), ", current block num ", currentBlockNumber)
	// }

	// TODO: move beacon client into the taraClient
	beaconLightClient, err := BeaconLightClient.NewBeaconLightClient(taraClient.Config.BeaconLightClientContractAddress, taraClient.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the BeaconLightClient contract: %v", err)
	}

	// TODO: move ethClientContract into the taraClient
	ethClientContract, err := EthClient.NewEthClient(taraClient.Config.EthClientContractAddress, taraClient.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the ethClientContract contract: %v", err)
	}

	relayer := new(Relayer)
	relayer.ethClient = ethClient
	relayer.taraClient = taraClient
	relayer.Config = config
	relayer.beaconLightClient = beaconLightClient
	relayer.ethClientContract = ethClientContract

	return relayer, nil
}

func (r *Relayer) Start(ctx context.Context) {
	r.onFinalizedEpoch = make(chan int64)
	r.onFinalizedBlockNumber = make(chan uint64)

	slot, err := r.beaconLightClient.Slot(nil)
	if err != nil {
		log.Fatalf("Failed to get current slot from contract: %v", err)
		return
	}
	r.currentPeriod = common.GetPeriodFromSlot(int64(slot))

	go r.startEventProcessing(ctx)
	go r.processNewBlocks(ctx)
	r.checkAndFinalize()
}

func (r *Relayer) Close() {
	close(r.onFinalizedEpoch)
	close(r.onFinalizedBlockNumber)
}

func (r *Relayer) processNewBlocks(ctx context.Context) {
	var finalizedBlockNumber uint64
	ticker := time.NewTicker(20 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case epoch := <-r.onFinalizedEpoch:
			log.Printf("Processing new block for epoch: %d", epoch)
			if r.currentPeriod != common.GetPeriodFromEpoch(epoch) {
				r.updateSyncCommittee(epoch)
				r.currentPeriod = common.GetPeriodFromEpoch(epoch)
			}
			if finalizedBlockNumber != 0 {
				log.Println("Updating light client with epoch", epoch, "and block number", finalizedBlockNumber)
				blockNum, err := r.updateLightClient(epoch, finalizedBlockNumber)
				if err != nil {
					log.Fatalf("Did not to update light client: %v", err)
				} else {
					go func() {
						r.getProof(blockNum)
						r.applyState(blockNum)
					}()
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
			r.checkAndFinalize()
		case <-ctx.Done():
			log.Println("Stopping new block processing")
			return
		}
	}
}

func (r *Relayer) checkAndFinalize() {
	ethEpoch, err := r.ethClient.BridgeContractClient.FinalizedEpoch(nil)
	if err != nil {
		log.Warningf("Failed to get finalized epoch from ETH contract: %v", err)
		return
	}
	taraEpoch, err := r.taraClient.BridgeContractClient.FinalizedEpoch(nil)
	if err != nil {
		log.Warningf("Failed to get finalized epoch from TARA contract: %v", err)
		return
	}
	if ethEpoch != taraEpoch {
		log.Printf("Finalizing ETH epoch %d on TARA epoch %d", ethEpoch, taraEpoch)
		r.finalize()
	}
}

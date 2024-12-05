package proof

import (
	"math/big"
	"relayer/bindings/BridgeBase"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// BlockchainClient defines the interface for blockchain operations required by getStateWithProof.
type BlockchainClient interface {
	LatestBlockNumber() (*big.Int, error)                                                  // Returns the block number
	GetStateWithProof(opts *bind.CallOpts) (BridgeBase.SharedStructsStateWithProof, error) // Gets the state with proof
	FinalizationInterval() *big.Int                                                        // Gets the finalization interval
}

// getStateWithProof fetches the state with proof for a given epoch and block number using the provided client.
func GetStateWithProof(client BlockchainClient, logger *log.Logger, epoch *big.Int, block_num *big.Int) (*BridgeBase.SharedStructsStateWithProof, error) {
	if block_num == nil {
		block, err := client.LatestBlockNumber()
		if err != nil || block == nil {
			logger.WithField("block", block).WithError(err).Panic("BlockByNumber")
			return nil, err
		}
		block_num = block
		logger.WithField("block", block_num).Info("BlockByNumber")
	}
	opts := bind.CallOpts{BlockNumber: block_num}

	stateWithProof, err := client.GetStateWithProof(&opts)
	logger.WithFields(log.Fields{"state": stateWithProof, "epoch": epoch, "opts": opts}).Info("GetStateWithProof")
	if err != nil {
		logger.WithError(err).Error("GetStateWithProof")
		return nil, err
	}

	// TODO: implement some binary search?
	interval := client.FinalizationInterval()
	if epoch == nil || epoch.Cmp(stateWithProof.State.Epoch) == 0 {
		return &stateWithProof, nil
	}

	if stateWithProof.State.Epoch.Cmp(epoch) > 0 {
		return GetStateWithProof(client, logger, epoch, block_num.Sub(block_num, interval))
	}

	return GetStateWithProof(client, logger, epoch, block_num.Add(block_num, interval))
}

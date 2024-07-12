package to_eth

import (
	"context"

	"relayer/internal/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TaraxaClientWrapper struct {
	*ethclient.Client
}

func NewClient(sharedClient *ethclient.Client) *TaraxaClientWrapper {
	taraxaClientWrapper := new(TaraxaClientWrapper)
	taraxaClientWrapper.Client = sharedClient

	return taraxaClientWrapper
}

func (t *TaraxaClientWrapper) GetPillarBlockData(period uint64) (*types.PillarBlockData, error) {
	var pillarBlockData *types.PillarBlockData
	includeSignatures := true
	err := t.Client.Client().CallContext(context.Background(), &pillarBlockData, "taraxa_getPillarBlockData", period, includeSignatures)
	if err == nil && pillarBlockData == nil {
		err = ethereum.NotFound
	}

	return pillarBlockData, err
}

func (t *TaraxaClientWrapper) GetTaraConfig() (*types.TaraConfig, error) {
	var taraConfig *types.TaraConfig
	err := t.Client.Client().CallContext(context.Background(), &taraConfig, "taraxa_getConfig")
	if err == nil && taraConfig == nil {
		err = ethereum.NotFound
	}

	return taraConfig, err
}

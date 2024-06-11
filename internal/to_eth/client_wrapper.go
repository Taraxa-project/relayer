package to_eth

import (
	"context"

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

func (t *TaraxaClientWrapper) GetPillarBlockData(period uint64, includeBinaryData bool) (*PillarBlockData, error) {
	var pillarBlockData *PillarBlockData
	err := t.Client.Client().CallContext(context.Background(), &pillarBlockData, "taraxa_getPillarBlockData", period, includeBinaryData)
	if err == nil && pillarBlockData == nil {
		err = ethereum.NotFound
	}

	return pillarBlockData, err
}

func (t *TaraxaClientWrapper) GetTaraConfig() (*TaraConfig, error) {
	var taraConfig *TaraConfig
	err := t.Client.Client().CallContext(context.Background(), &taraConfig, "taraxa_getConfig")
	if err == nil && taraConfig == nil {
		err = ethereum.NotFound
	}

	return taraConfig, err
}

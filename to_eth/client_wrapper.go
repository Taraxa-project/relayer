package to_eth

import (
	"context"

	tara_rpc_types "github.com/Taraxa-project/taraxa-contracts-go-clients/clients/tara/rpc_client/types"

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

func (t *TaraxaClientWrapper) GetPillarBlockData(period uint64, includeBinaryData bool) (*tara_rpc_types.PillarBlockData, error) {
	var pillarBlockData *tara_rpc_types.PillarBlockData
	err := t.Client.Client().CallContext(context.Background(), &pillarBlockData, "taraxa_getPillarBlockData", period, includeBinaryData)
	if err == nil && pillarBlockData == nil {
		err = ethereum.NotFound
	}

	return pillarBlockData, err
}

func (t *TaraxaClientWrapper) GetTaraConfig() (*tara_rpc_types.TaraConfig, error) {
	var taraConfig *tara_rpc_types.TaraConfig
	err := t.Client.Client().CallContext(context.Background(), &taraConfig, "taraxa_getConfig")
	if err == nil && taraConfig == nil {
		err = ethereum.NotFound
	}

	return taraConfig, err
}

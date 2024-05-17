package rpc_client

import (
	"context"

	"github.com/Taraxa-project/relayer/clients/client_base"
	tara_rpc_types "github.com/Taraxa-project/relayer/clients/tara/rpc_client/types"

	"github.com/ethereum/go-ethereum"
)

type RpcClient struct {
	*client_base.ClientBase
}

func NewRpcClient(config client_base.NetConfig) (*RpcClient, error) {
	var err error

	rpcClient := new(RpcClient)
	rpcClient.ClientBase, err = client_base.NewClientBase(config, nil)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}

func NewSharedRpcClient(sharedClient *client_base.ClientBase) *RpcClient {
	rpcClient := new(RpcClient)
	rpcClient.ClientBase = sharedClient

	return rpcClient
}

func (RpcClient *RpcClient) GetPillarBlockData(period uint64, includeBinaryData bool) (*tara_rpc_types.PillarBlockData, error) {
	var pillarBlockData *tara_rpc_types.PillarBlockData
	err := RpcClient.EthClient.Client().CallContext(context.Background(), &pillarBlockData, "taraxa_getPillarBlockData", period, includeBinaryData)
	if err == nil && pillarBlockData == nil {
		err = ethereum.NotFound
	}

	return pillarBlockData, err
}

func (RpcClient *RpcClient) GetTaraConfig() (*tara_rpc_types.TaraConfig, error) {
	var taraConfig *tara_rpc_types.TaraConfig
	err := RpcClient.EthClient.Client().CallContext(context.Background(), &taraConfig, "taraxa_getConfig")
	if err == nil && taraConfig == nil {
		err = ethereum.NotFound
	}

	return taraConfig, err
}

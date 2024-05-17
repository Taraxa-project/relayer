package tara_client

import (
	"github.com/Taraxa-project/relayer/clients/bridge_contract_client"
	"github.com/Taraxa-project/relayer/clients/tara/rpc_client"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Taraxa-project/relayer/clients/client_base"
)

type TaraClient struct {
	*client_base.ClientBase
	Config               TaraClientConfig
	BridgeContractClient *bridge_contract_client.BridgeContractClient
	RpcClient            *rpc_client.RpcClient
}

type TaraClientConfig struct {
	client_base.NetConfig
	BridgeContractAddress            common.Address `json:"bridge_contract_address"`
	DposContractAddress              common.Address `json:"dpos_contract_address"`
	EthClientContractAddress         common.Address `json:"eth_client_contract_address"`
	BeaconLightClientContractAddress common.Address `json:"beacon_light_client_contract_address"`
}

func NewTaraClient(config TaraClientConfig, privateKeyStr *string) (*TaraClient, error) {
	var err error

	ethClient := new(TaraClient)
	ethClient.ClientBase, err = client_base.NewClientBase(config.NetConfig, privateKeyStr)
	if err != nil {
		return nil, err
	}

	taraClient := new(TaraClient)
	taraClient.BridgeContractClient, err = bridge_contract_client.NewSharedBridgeContractClient(ethClient.ClientBase, config.BridgeContractAddress)
	if err != nil {
		return nil, err
	}

	taraClient.RpcClient = rpc_client.NewSharedRpcClient(ethClient.ClientBase)
	taraClient.Config = config

	return taraClient, nil
}
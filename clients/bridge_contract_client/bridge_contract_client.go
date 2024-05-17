package bridge_contract_client

import (
	bridge_contract_interface "github.com/Taraxa-project/relayer/clients/bridge_contract_client/contract_interface"
	"github.com/Taraxa-project/relayer/clients/client_base"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BridgeContractClient struct {
	*client_base.ClientBase
	*bridge_contract_interface.BridgeContractInterface
}

func NewBridgeContractClient(config client_base.NetConfig, communicationProtocol client_base.CommunicationProtocol, privateKeyStr *string) (*BridgeContractClient, error) {
	var err error

	bridgeContractClient := new(BridgeContractClient)
	bridgeContractClient.ClientBase, err = client_base.NewClientBase(config, communicationProtocol, privateKeyStr)
	if err != nil {
		return nil, err
	}

	// TODO: fix config
	bridgeContractClient.BridgeContractInterface, err = bridge_contract_interface.NewBridgeContractInterface(bridgeContractClient.Config.ContractAddress, bridgeContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return bridgeContractClient, nil
}

func NewSharedBridgeContractClient(clientBase *client_base.ClientBase, contractAddress common.Address) (*BridgeContractClient, error) {
	var err error

	bridgeContractClient := new(BridgeContractClient)
	bridgeContractClient.ClientBase = clientBase

	bridgeContractClient.BridgeContractInterface, err = bridge_contract_interface.NewBridgeContractInterface(contractAddress, bridgeContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return bridgeContractClient, nil
}

func (BridgeContractClient *BridgeContractClient) GetStateWithProof() (bridge_contract_interface.SharedStructsStateWithProof, error) {
	return BridgeContractClient.BridgeContractInterface.GetStateWithProof(&bind.CallOpts{})
}

func (BridgeContractClient *BridgeContractClient) ApplyState(stateWithProof bridge_contract_interface.SharedStructsStateWithProof) (*types.Transaction, error) {
	transactOpts, err := BridgeContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return BridgeContractClient.BridgeContractInterface.ApplyState(transactOpts, stateWithProof)
}

func (BridgeContractClient *BridgeContractClient) FinalizeEpoch() (*types.Transaction, error) {
	transactOpts, err := BridgeContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return BridgeContractClient.BridgeContractInterface.FinalizeEpoch(transactOpts)
}

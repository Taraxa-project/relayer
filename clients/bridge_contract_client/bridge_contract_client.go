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
	bridgeInterface *bridge_contract_interface.BridgeContractInterface
}

func NewBridgeContractClient(config client_base.NetConfig, communicationProtocol client_base.CommunicationProtocol) (*BridgeContractClient, error) {
	var err error

	bridgeContractClient := new(BridgeContractClient)
	bridgeContractClient.ClientBase, err = client_base.NewClientBase(config, communicationProtocol)
	if err != nil {
		return nil, err
	}

	// TODO: fix config
	bridgeContractClient.bridgeInterface, err = bridge_contract_interface.NewBridgeContractInterface(bridgeContractClient.Config.ContractAddress, bridgeContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return bridgeContractClient, nil
}

func NewSharedBridgeContractClient(clientBase *client_base.ClientBase, contractAddress common.Address) (*BridgeContractClient, error) {
	var err error

	bridgeContractClient := new(BridgeContractClient)
	bridgeContractClient.ClientBase = clientBase

	bridgeContractClient.bridgeInterface, err = bridge_contract_interface.NewBridgeContractInterface(contractAddress, bridgeContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return bridgeContractClient, nil
}

func (BridgeContractClient *BridgeContractClient) GetStateWithProof() (bridge_contract_interface.SharedStructsStateWithProof, error) {
	return BridgeContractClient.bridgeInterface.GetStateWithProof(&bind.CallOpts{})
}

func (BridgeContractClient *BridgeContractClient) ApplyState(transactor *client_base.Transactor, stateWithProof bridge_contract_interface.SharedStructsStateWithProof) (*types.Transaction, error) {
	transactOpts, err := BridgeContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return BridgeContractClient.bridgeInterface.ApplyState(transactOpts, stateWithProof)
}

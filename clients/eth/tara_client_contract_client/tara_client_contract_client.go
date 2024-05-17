package tara_client_contract_client

import (
	"github.com/Taraxa-project/relayer/clients/client_base"
	tara_client_contract_interface "github.com/Taraxa-project/relayer/clients/eth/tara_client_contract_client/contract_interface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TaraClientContractClient contains variables needed for communication with taraxa client smart contract on eth
type TaraClientContractClient struct {
	*client_base.ClientBase
	*tara_client_contract_interface.TaraClientContractInterface
}

func NewTaraClientContractClient(config client_base.NetConfig, privateKeyStr *string) (*TaraClientContractClient, error) {
	var err error

	taraClientContractClient := new(TaraClientContractClient)
	taraClientContractClient.ClientBase, err = client_base.NewClientBase(config, privateKeyStr)
	if err != nil {
		return nil, err
	}

	taraClientContractClient.TaraClientContractInterface, err = tara_client_contract_interface.NewTaraClientContractInterface(taraClientContractClient.Config.ContractAddress, taraClientContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient, nil
}

func NewSharedTaraClientContractClient(clientBase *client_base.ClientBase, contractAddress common.Address) (*TaraClientContractClient, error) {
	var err error

	taraClientContractClient := new(TaraClientContractClient)
	taraClientContractClient.ClientBase = clientBase

	taraClientContractClient.TaraClientContractInterface, err = tara_client_contract_interface.NewTaraClientContractInterface(contractAddress, taraClientContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return taraClientContractClient, nil
}

func (taraClientContractClient *TaraClientContractClient) GetFinalizedPillarBlock() (tara_client_contract_interface.PillarBlockFinalizedBlock, error) {
	return taraClientContractClient.TaraClientContractInterface.GetFinalized(&bind.CallOpts{})
}

func (taraClientContractClient *TaraClientContractClient) FinalizeBlocks(blocks []tara_client_contract_interface.PillarBlockWithChanges, lastBlockSigs []tara_client_contract_interface.CompactSignature) (*types.Transaction, error) {
	transactOpts, err := taraClientContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return taraClientContractClient.TaraClientContractInterface.FinalizeBlocks(transactOpts, blocks, lastBlockSigs)
}

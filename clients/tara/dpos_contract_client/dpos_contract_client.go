package dpos_contract_client

import (
	"math/big"

	"github.com/Taraxa-project/relayer/clients/client_base"
	dpos_interface "github.com/Taraxa-project/relayer/clients/tara/dpos_contract_client/dpos_interface"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DposContractClient contains variables needed for communication with taraxa dpos smart contract
type DposContractClient struct {
	*client_base.ClientBase
	dposInterface *dpos_interface.DposInterface
}

func NewDposContractClient(config client_base.NetConfig, communicationProtocol client_base.CommunicationProtocol) (*DposContractClient, error) {
	var err error

	dposContractClient := new(DposContractClient)
	dposContractClient.ClientBase, err = client_base.NewClientBase(config, communicationProtocol)
	if err != nil {
		return nil, err
	}

	dposContractClient.dposInterface, err = dpos_interface.NewDposInterface(dposContractClient.Config.ContractAddress, dposContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return dposContractClient, nil
}

func NewSharedDposContractClient(sharedClient *client_base.ClientBase, contractAddress common.Address) (*DposContractClient, error) {
	var err error

	dposContractClient := new(DposContractClient)
	dposContractClient.ClientBase = sharedClient

	dposContractClient.dposInterface, err = dpos_interface.NewDposInterface(contractAddress, dposContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return dposContractClient, nil
}

func (DposContractClient *DposContractClient) GetTotalEligibleVotesCount() (uint64, error) {
	return DposContractClient.dposInterface.GetTotalEligibleVotesCount(&bind.CallOpts{})
}

func (DposContractClient *DposContractClient) GetValidator(validator common.Address) (dpos_interface.DposInterfaceValidatorBasicInfo, error) {
	return DposContractClient.dposInterface.GetValidator(&bind.CallOpts{}, validator)
}

func (DposContractClient *DposContractClient) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return DposContractClient.dposInterface.GetValidatorEligibleVotesCount(&bind.CallOpts{}, validator)
}

func (DposContractClient *DposContractClient) IsValidatorEligible(validator common.Address) (bool, error) {
	return DposContractClient.dposInterface.IsValidatorEligible(&bind.CallOpts{}, validator)
}

func (DposContractClient *DposContractClient) GetValidators() ([]dpos_interface.DposInterfaceValidatorData, error) {
	return DposContractClient.getValidators(&bind.CallOpts{})
}

func (DposContractClient *DposContractClient) GetValidatorsAtBlock(block_num *big.Int) ([]dpos_interface.DposInterfaceValidatorData, error) {
	return DposContractClient.getValidators(&bind.CallOpts{BlockNumber: block_num})
}

func (DposContractClient *DposContractClient) getValidators(opts *bind.CallOpts) ([]dpos_interface.DposInterfaceValidatorData, error) {
	var validators []dpos_interface.DposInterfaceValidatorData

	for batch := uint32(0); ; batch++ {
		validatorsBatch, err := DposContractClient.dposInterface.GetValidators(opts, batch)
		if err != nil {
			return nil, err
		}

		if len(validatorsBatch.Validators) != 0 {
			validators = append(validators, validatorsBatch.Validators...)
		}

		if validatorsBatch.End == true {
			break
		}
	}

	return validators, nil
}

func (DposContractClient *DposContractClient) GetOwnerValidators(owner common.Address) ([]dpos_interface.DposInterfaceValidatorData, error) {
	var validators []dpos_interface.DposInterfaceValidatorData

	for batch := uint32(0); ; batch++ {
		validatorsBatch, err := DposContractClient.dposInterface.GetValidatorsFor(&bind.CallOpts{}, owner, batch)
		if err != nil {
			return nil, err
		}

		if len(validatorsBatch.Validators) != 0 {
			validators = append(validators, validatorsBatch.Validators...)
		}

		if validatorsBatch.End == true {
			break
		}
	}

	return validators, nil
}

func (DposContractClient *DposContractClient) GetDelegations(delegator common.Address) ([]dpos_interface.DposInterfaceDelegationData, error) {
	var delegations []dpos_interface.DposInterfaceDelegationData

	for batch := uint32(0); ; batch++ {
		delegationsBatch, err := DposContractClient.dposInterface.GetDelegations(&bind.CallOpts{}, delegator, batch)
		if err != nil {
			return nil, err
		}

		if len(delegationsBatch.Delegations) != 0 {
			delegations = append(delegations, delegationsBatch.Delegations...)
		}

		if delegationsBatch.End == true {
			break
		}
	}

	return delegations, nil
}

func (DposContractClient *DposContractClient) GetUndelegations(delegator common.Address) ([]dpos_interface.DposInterfaceUndelegationData, error) {
	var undelegations []dpos_interface.DposInterfaceUndelegationData

	for batch := uint32(0); ; batch++ {
		undelegationsBatch, err := DposContractClient.dposInterface.GetUndelegations(&bind.CallOpts{}, delegator, batch)
		if err != nil {
			return nil, err
		}

		if len(undelegationsBatch.Undelegations) != 0 {
			undelegations = append(undelegations, undelegationsBatch.Undelegations...)
		}

		if undelegationsBatch.End == true {
			break
		}
	}

	return undelegations, nil
}

func (DposContractClient *DposContractClient) Delegate(transactor *client_base.Transactor, amount *big.Int, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	transactOpts.Value = amount // in wei
	return DposContractClient.dposInterface.Delegate(transactOpts, validator)
}

func (DposContractClient *DposContractClient) Undelegate(transactor *client_base.Transactor, amount *big.Int, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.Undelegate(transactOpts, validator, amount)
}

func (DposContractClient *DposContractClient) ConfirmUndelegate(transactor *client_base.Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.ConfirmUndelegate(transactOpts, validator)
}

func (DposContractClient *DposContractClient) CancelUndelegate(transactor *client_base.Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.CancelUndelegate(transactOpts, validator)
}

func (DposContractClient *DposContractClient) RedelegateUndelegate(transactor *client_base.Transactor, amount *big.Int, validatorFrom common.Address, validatorTo common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.ReDelegate(transactOpts, validatorFrom, validatorTo, amount)
}

func (DposContractClient *DposContractClient) ClaimRewards(transactor *client_base.Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.ClaimRewards(transactOpts, validator)
}

func (DposContractClient *DposContractClient) ClaimCommissionRewards(transactor *client_base.Transactor, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.ClaimCommissionRewards(transactOpts, validator)
}

func (DposContractClient *DposContractClient) RegisterValidator(transactor *client_base.Transactor, validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.RegisterValidator(transactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

func (DposContractClient *DposContractClient) SetValidatorInfo(transactor *client_base.Transactor, validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.SetValidatorInfo(transactOpts, validator, description, endpoint)
}

func (DposContractClient *DposContractClient) SetCommission(transactor *client_base.Transactor, validator common.Address, commission uint16) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.CreateNewTransactOpts(transactor)
	if err != nil {
		return nil, err
	}

	return DposContractClient.dposInterface.SetCommission(transactOpts, validator, commission)
}

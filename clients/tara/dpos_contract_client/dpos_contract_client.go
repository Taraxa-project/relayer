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
	*dpos_interface.DposInterface
}

func NewDposContractClient(config client_base.NetConfig, privateKeyStr *string) (*DposContractClient, error) {
	var err error

	dposContractClient := new(DposContractClient)
	dposContractClient.ClientBase, err = client_base.NewClientBase(config, privateKeyStr)
	if err != nil {
		return nil, err
	}

	dposContractClient.DposInterface, err = dpos_interface.NewDposInterface(dposContractClient.Config.ContractAddress, dposContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return dposContractClient, nil
}

func NewSharedDposContractClient(sharedClient *client_base.ClientBase, contractAddress common.Address) (*DposContractClient, error) {
	var err error

	dposContractClient := new(DposContractClient)
	dposContractClient.ClientBase = sharedClient

	dposContractClient.DposInterface, err = dpos_interface.NewDposInterface(contractAddress, dposContractClient.EthClient)
	if err != nil {
		return nil, err
	}

	return dposContractClient, nil
}

func (DposContractClient *DposContractClient) GetTotalEligibleVotesCount() (uint64, error) {
	return DposContractClient.DposInterface.GetTotalEligibleVotesCount(&bind.CallOpts{})
}

func (DposContractClient *DposContractClient) GetValidator(validator common.Address) (dpos_interface.DposInterfaceValidatorBasicInfo, error) {
	return DposContractClient.DposInterface.GetValidator(&bind.CallOpts{}, validator)
}

func (DposContractClient *DposContractClient) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return DposContractClient.DposInterface.GetValidatorEligibleVotesCount(&bind.CallOpts{}, validator)
}

func (DposContractClient *DposContractClient) IsValidatorEligible(validator common.Address) (bool, error) {
	return DposContractClient.DposInterface.IsValidatorEligible(&bind.CallOpts{}, validator)
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
		validatorsBatch, err := DposContractClient.DposInterface.GetValidators(opts, batch)
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
		validatorsBatch, err := DposContractClient.DposInterface.GetValidatorsFor(&bind.CallOpts{}, owner, batch)
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
		delegationsBatch, err := DposContractClient.DposInterface.GetDelegations(&bind.CallOpts{}, delegator, batch)
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
		undelegationsBatch, err := DposContractClient.DposInterface.GetUndelegations(&bind.CallOpts{}, delegator, batch)
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

func (DposContractClient *DposContractClient) Delegate(amount *big.Int, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	transactOpts.Value = amount // in wei
	return DposContractClient.DposInterface.Delegate(transactOpts, validator)
}

func (DposContractClient *DposContractClient) Undelegate(amount *big.Int, validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.Undelegate(transactOpts, validator, amount)
}

func (DposContractClient *DposContractClient) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.ConfirmUndelegate(transactOpts, validator)
}

func (DposContractClient *DposContractClient) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.CancelUndelegate(transactOpts, validator)
}

func (DposContractClient *DposContractClient) RedelegateUndelegate(amount *big.Int, validatorFrom common.Address, validatorTo common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.ReDelegate(transactOpts, validatorFrom, validatorTo, amount)
}

func (DposContractClient *DposContractClient) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.ClaimRewards(transactOpts, validator)
}

func (DposContractClient *DposContractClient) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.ClaimCommissionRewards(transactOpts, validator)
}

func (DposContractClient *DposContractClient) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.RegisterValidator(transactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

func (DposContractClient *DposContractClient) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.SetValidatorInfo(transactOpts, validator, description, endpoint)
}

func (DposContractClient *DposContractClient) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	transactOpts, err := DposContractClient.GenTransactOpts()
	if err != nil {
		return nil, err
	}

	return DposContractClient.DposInterface.SetCommission(transactOpts, validator, commission)
}

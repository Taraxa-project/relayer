// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dpos_interface

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DposInterfaceDelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegationData struct {
	Account    common.Address
	Delegation DposInterfaceDelegatorInfo
}

// DposInterfaceDelegatorInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegatorInfo struct {
	Stake   *big.Int
	Rewards *big.Int
}

// DposInterfaceUndelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceUndelegationData struct {
	Stake           *big.Int
	Block           uint64
	Validator       common.Address
	ValidatorExists bool
}

// DposInterfaceValidatorBasicInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorBasicInfo struct {
	TotalStake           *big.Int
	CommissionReward     *big.Int
	Commission           uint16
	LastCommissionChange uint64
	UndelegationsCount   uint16
	Owner                common.Address
	Description          string
	Endpoint             string
}

// DposInterfaceValidatorData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorData struct {
	Account common.Address
	Info    DposInterfaceValidatorBasicInfo
}

// DposInterfaceMetaData contains all meta data concerning the DposInterface contract.
var DposInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"}],\"name\":\"CommissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorInfoSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"cancelUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAllRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimCommissionRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"confirmUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"internalType\":\"structDposInterface.DelegatorInfo\",\"name\":\"delegation\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.DelegationData[]\",\"name\":\"delegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getTotalDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total_delegation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getUndelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"validator_exists\",\"type\":\"bool\"}],\"internalType\":\"structDposInterface.UndelegationData[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"undelegations_count\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"validator_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"undelegations_count\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidatorsFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"undelegations_count\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidatorEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrf_key\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"setValidatorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DposInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use DposInterfaceMetaData.ABI instead.
var DposInterfaceABI = DposInterfaceMetaData.ABI

// DposInterface is an auto generated Go binding around an Ethereum contract.
type DposInterface struct {
	DposInterfaceCaller     // Read-only binding to the contract
	DposInterfaceTransactor // Write-only binding to the contract
	DposInterfaceFilterer   // Log filterer for contract events
}

// DposInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DposInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DposInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DposInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DposInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DposInterfaceSession struct {
	Contract     *DposInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DposInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DposInterfaceCallerSession struct {
	Contract *DposInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DposInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DposInterfaceTransactorSession struct {
	Contract     *DposInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DposInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DposInterfaceRaw struct {
	Contract *DposInterface // Generic contract binding to access the raw methods on
}

// DposInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DposInterfaceCallerRaw struct {
	Contract *DposInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DposInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DposInterfaceTransactorRaw struct {
	Contract *DposInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDposInterface creates a new instance of DposInterface, bound to a specific deployed contract.
func NewDposInterface(address common.Address, backend bind.ContractBackend) (*DposInterface, error) {
	contract, err := bindDposInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DposInterface{DposInterfaceCaller: DposInterfaceCaller{contract: contract}, DposInterfaceTransactor: DposInterfaceTransactor{contract: contract}, DposInterfaceFilterer: DposInterfaceFilterer{contract: contract}}, nil
}

// NewDposInterfaceCaller creates a new read-only instance of DposInterface, bound to a specific deployed contract.
func NewDposInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DposInterfaceCaller, error) {
	contract, err := bindDposInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceCaller{contract: contract}, nil
}

// NewDposInterfaceTransactor creates a new write-only instance of DposInterface, bound to a specific deployed contract.
func NewDposInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DposInterfaceTransactor, error) {
	contract, err := bindDposInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceTransactor{contract: contract}, nil
}

// NewDposInterfaceFilterer creates a new log filterer instance of DposInterface, bound to a specific deployed contract.
func NewDposInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DposInterfaceFilterer, error) {
	contract, err := bindDposInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceFilterer{contract: contract}, nil
}

// bindDposInterface binds a generic wrapper to an already deployed contract.
func bindDposInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DposInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DposInterface *DposInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DposInterface.Contract.DposInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DposInterface *DposInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposInterface.Contract.DposInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DposInterface *DposInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DposInterface.Contract.DposInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DposInterface *DposInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DposInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DposInterface *DposInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DposInterface *DposInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DposInterface.Contract.contract.Transact(opts, method, params...)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_DposInterface *DposInterfaceCaller) GetDelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getDelegations", delegator, batch)

	outstruct := new(struct {
		Delegations []DposInterfaceDelegationData
		End         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delegations = *abi.ConvertType(out[0], new([]DposInterfaceDelegationData)).(*[]DposInterfaceDelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_DposInterface *DposInterfaceSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _DposInterface.Contract.GetDelegations(&_DposInterface.CallOpts, delegator, batch)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_DposInterface *DposInterfaceCallerSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _DposInterface.Contract.GetDelegations(&_DposInterface.CallOpts, delegator, batch)
}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address delegator) view returns(uint256 total_delegation)
func (_DposInterface *DposInterfaceCaller) GetTotalDelegation(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getTotalDelegation", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address delegator) view returns(uint256 total_delegation)
func (_DposInterface *DposInterfaceSession) GetTotalDelegation(delegator common.Address) (*big.Int, error) {
	return _DposInterface.Contract.GetTotalDelegation(&_DposInterface.CallOpts, delegator)
}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address delegator) view returns(uint256 total_delegation)
func (_DposInterface *DposInterfaceCallerSession) GetTotalDelegation(delegator common.Address) (*big.Int, error) {
	return _DposInterface.Contract.GetTotalDelegation(&_DposInterface.CallOpts, delegator)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_DposInterface *DposInterfaceCaller) GetTotalEligibleVotesCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getTotalEligibleVotesCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_DposInterface *DposInterfaceSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _DposInterface.Contract.GetTotalEligibleVotesCount(&_DposInterface.CallOpts)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_DposInterface *DposInterfaceCallerSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _DposInterface.Contract.GetTotalEligibleVotesCount(&_DposInterface.CallOpts)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_DposInterface *DposInterfaceCaller) GetUndelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getUndelegations", delegator, batch)

	outstruct := new(struct {
		Undelegations []DposInterfaceUndelegationData
		End           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Undelegations = *abi.ConvertType(out[0], new([]DposInterfaceUndelegationData)).(*[]DposInterfaceUndelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_DposInterface *DposInterfaceSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _DposInterface.Contract.GetUndelegations(&_DposInterface.CallOpts, delegator, batch)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_DposInterface *DposInterfaceCallerSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _DposInterface.Contract.GetUndelegations(&_DposInterface.CallOpts, delegator, batch)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,uint16,address,string,string) validator_info)
func (_DposInterface *DposInterfaceCaller) GetValidator(opts *bind.CallOpts, validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getValidator", validator)

	if err != nil {
		return *new(DposInterfaceValidatorBasicInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DposInterfaceValidatorBasicInfo)).(*DposInterfaceValidatorBasicInfo)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,uint16,address,string,string) validator_info)
func (_DposInterface *DposInterfaceSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _DposInterface.Contract.GetValidator(&_DposInterface.CallOpts, validator)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,uint16,address,string,string) validator_info)
func (_DposInterface *DposInterfaceCallerSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _DposInterface.Contract.GetValidator(&_DposInterface.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_DposInterface *DposInterfaceCaller) GetValidatorEligibleVotesCount(opts *bind.CallOpts, validator common.Address) (uint64, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getValidatorEligibleVotesCount", validator)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_DposInterface *DposInterfaceSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _DposInterface.Contract.GetValidatorEligibleVotesCount(&_DposInterface.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_DposInterface *DposInterfaceCallerSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _DposInterface.Contract.GetValidatorEligibleVotesCount(&_DposInterface.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposInterface *DposInterfaceCaller) GetValidators(opts *bind.CallOpts, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getValidators", batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposInterface *DposInterfaceSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposInterface.Contract.GetValidators(&_DposInterface.CallOpts, batch)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposInterface *DposInterfaceCallerSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposInterface.Contract.GetValidators(&_DposInterface.CallOpts, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposInterface *DposInterfaceCaller) GetValidatorsFor(opts *bind.CallOpts, owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "getValidatorsFor", owner, batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposInterface *DposInterfaceSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposInterface.Contract.GetValidatorsFor(&_DposInterface.CallOpts, owner, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,uint16,address,string,string))[] validators, bool end)
func (_DposInterface *DposInterfaceCallerSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _DposInterface.Contract.GetValidatorsFor(&_DposInterface.CallOpts, owner, batch)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_DposInterface *DposInterfaceCaller) IsValidatorEligible(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _DposInterface.contract.Call(opts, &out, "isValidatorEligible", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_DposInterface *DposInterfaceSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _DposInterface.Contract.IsValidatorEligible(&_DposInterface.CallOpts, validator)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_DposInterface *DposInterfaceCallerSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _DposInterface.Contract.IsValidatorEligible(&_DposInterface.CallOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_DposInterface *DposInterfaceTransactor) CancelUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "cancelUndelegate", validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_DposInterface *DposInterfaceSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.CancelUndelegate(&_DposInterface.TransactOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_DposInterface *DposInterfaceTransactorSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.CancelUndelegate(&_DposInterface.TransactOpts, validator)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x0b83a727.
//
// Solidity: function claimAllRewards() returns()
func (_DposInterface *DposInterfaceTransactor) ClaimAllRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "claimAllRewards")
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x0b83a727.
//
// Solidity: function claimAllRewards() returns()
func (_DposInterface *DposInterfaceSession) ClaimAllRewards() (*types.Transaction, error) {
	return _DposInterface.Contract.ClaimAllRewards(&_DposInterface.TransactOpts)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x0b83a727.
//
// Solidity: function claimAllRewards() returns()
func (_DposInterface *DposInterfaceTransactorSession) ClaimAllRewards() (*types.Transaction, error) {
	return _DposInterface.Contract.ClaimAllRewards(&_DposInterface.TransactOpts)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_DposInterface *DposInterfaceTransactor) ClaimCommissionRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "claimCommissionRewards", validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_DposInterface *DposInterfaceSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.ClaimCommissionRewards(&_DposInterface.TransactOpts, validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_DposInterface *DposInterfaceTransactorSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.ClaimCommissionRewards(&_DposInterface.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_DposInterface *DposInterfaceTransactor) ClaimRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "claimRewards", validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_DposInterface *DposInterfaceSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.ClaimRewards(&_DposInterface.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_DposInterface *DposInterfaceTransactorSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.ClaimRewards(&_DposInterface.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_DposInterface *DposInterfaceTransactor) ConfirmUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "confirmUndelegate", validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_DposInterface *DposInterfaceSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.ConfirmUndelegate(&_DposInterface.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_DposInterface *DposInterfaceTransactorSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.ConfirmUndelegate(&_DposInterface.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_DposInterface *DposInterfaceTransactor) Delegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "delegate", validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_DposInterface *DposInterfaceSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.Delegate(&_DposInterface.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_DposInterface *DposInterfaceTransactorSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _DposInterface.Contract.Delegate(&_DposInterface.TransactOpts, validator)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_DposInterface *DposInterfaceTransactor) ReDelegate(opts *bind.TransactOpts, validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "reDelegate", validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_DposInterface *DposInterfaceSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposInterface.Contract.ReDelegate(&_DposInterface.TransactOpts, validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_DposInterface *DposInterfaceTransactorSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposInterface.Contract.ReDelegate(&_DposInterface.TransactOpts, validator_from, validator_to, amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_DposInterface *DposInterfaceTransactor) RegisterValidator(opts *bind.TransactOpts, validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "registerValidator", validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_DposInterface *DposInterfaceSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _DposInterface.Contract.RegisterValidator(&_DposInterface.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_DposInterface *DposInterfaceTransactorSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _DposInterface.Contract.RegisterValidator(&_DposInterface.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_DposInterface *DposInterfaceTransactor) SetCommission(opts *bind.TransactOpts, validator common.Address, commission uint16) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "setCommission", validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_DposInterface *DposInterfaceSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _DposInterface.Contract.SetCommission(&_DposInterface.TransactOpts, validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_DposInterface *DposInterfaceTransactorSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _DposInterface.Contract.SetCommission(&_DposInterface.TransactOpts, validator, commission)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_DposInterface *DposInterfaceTransactor) SetValidatorInfo(opts *bind.TransactOpts, validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "setValidatorInfo", validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_DposInterface *DposInterfaceSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _DposInterface.Contract.SetValidatorInfo(&_DposInterface.TransactOpts, validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_DposInterface *DposInterfaceTransactorSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _DposInterface.Contract.SetValidatorInfo(&_DposInterface.TransactOpts, validator, description, endpoint)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_DposInterface *DposInterfaceTransactor) Undelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposInterface.contract.Transact(opts, "undelegate", validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_DposInterface *DposInterfaceSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposInterface.Contract.Undelegate(&_DposInterface.TransactOpts, validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_DposInterface *DposInterfaceTransactorSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DposInterface.Contract.Undelegate(&_DposInterface.TransactOpts, validator, amount)
}

// DposInterfaceCommissionRewardsClaimedIterator is returned from FilterCommissionRewardsClaimed and is used to iterate over the raw logs and unpacked data for CommissionRewardsClaimed events raised by the DposInterface contract.
type DposInterfaceCommissionRewardsClaimedIterator struct {
	Event *DposInterfaceCommissionRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceCommissionRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceCommissionRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceCommissionRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceCommissionRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceCommissionRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceCommissionRewardsClaimed represents a CommissionRewardsClaimed event raised by the DposInterface contract.
type DposInterfaceCommissionRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionRewardsClaimed is a free log retrieval operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) FilterCommissionRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*DposInterfaceCommissionRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceCommissionRewardsClaimedIterator{contract: _DposInterface.contract, event: "CommissionRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchCommissionRewardsClaimed is a free log subscription operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) WatchCommissionRewardsClaimed(opts *bind.WatchOpts, sink chan<- *DposInterfaceCommissionRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceCommissionRewardsClaimed)
				if err := _DposInterface.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommissionRewardsClaimed is a log parse operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) ParseCommissionRewardsClaimed(log types.Log) (*DposInterfaceCommissionRewardsClaimed, error) {
	event := new(DposInterfaceCommissionRewardsClaimed)
	if err := _DposInterface.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceCommissionSetIterator is returned from FilterCommissionSet and is used to iterate over the raw logs and unpacked data for CommissionSet events raised by the DposInterface contract.
type DposInterfaceCommissionSetIterator struct {
	Event *DposInterfaceCommissionSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceCommissionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceCommissionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceCommissionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceCommissionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceCommissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceCommissionSet represents a CommissionSet event raised by the DposInterface contract.
type DposInterfaceCommissionSet struct {
	Validator  common.Address
	Commission uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommissionSet is a free log retrieval operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_DposInterface *DposInterfaceFilterer) FilterCommissionSet(opts *bind.FilterOpts, validator []common.Address) (*DposInterfaceCommissionSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceCommissionSetIterator{contract: _DposInterface.contract, event: "CommissionSet", logs: logs, sub: sub}, nil
}

// WatchCommissionSet is a free log subscription operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_DposInterface *DposInterfaceFilterer) WatchCommissionSet(opts *bind.WatchOpts, sink chan<- *DposInterfaceCommissionSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceCommissionSet)
				if err := _DposInterface.contract.UnpackLog(event, "CommissionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommissionSet is a log parse operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_DposInterface *DposInterfaceFilterer) ParseCommissionSet(log types.Log) (*DposInterfaceCommissionSet, error) {
	event := new(DposInterfaceCommissionSet)
	if err := _DposInterface.contract.UnpackLog(event, "CommissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the DposInterface contract.
type DposInterfaceDelegatedIterator struct {
	Event *DposInterfaceDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceDelegated represents a Delegated event raised by the DposInterface contract.
type DposInterfaceDelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposInterfaceDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceDelegatedIterator{contract: _DposInterface.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *DposInterfaceDelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceDelegated)
				if err := _DposInterface.contract.UnpackLog(event, "Delegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) ParseDelegated(log types.Log) (*DposInterfaceDelegated, error) {
	event := new(DposInterfaceDelegated)
	if err := _DposInterface.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the DposInterface contract.
type DposInterfaceRedelegatedIterator struct {
	Event *DposInterfaceRedelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceRedelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceRedelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceRedelegated represents a Redelegated event raised by the DposInterface contract.
type DposInterfaceRedelegated struct {
	Delegator common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) FilterRedelegated(opts *bind.FilterOpts, delegator []common.Address, from []common.Address, to []common.Address) (*DposInterfaceRedelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceRedelegatedIterator{contract: _DposInterface.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *DposInterfaceRedelegated, delegator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceRedelegated)
				if err := _DposInterface.contract.UnpackLog(event, "Redelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedelegated is a log parse operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) ParseRedelegated(log types.Log) (*DposInterfaceRedelegated, error) {
	event := new(DposInterfaceRedelegated)
	if err := _DposInterface.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the DposInterface contract.
type DposInterfaceRewardsClaimedIterator struct {
	Event *DposInterfaceRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceRewardsClaimed represents a RewardsClaimed event raised by the DposInterface contract.
type DposInterfaceRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*DposInterfaceRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceRewardsClaimedIterator{contract: _DposInterface.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *DposInterfaceRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceRewardsClaimed)
				if err := _DposInterface.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) ParseRewardsClaimed(log types.Log) (*DposInterfaceRewardsClaimed, error) {
	event := new(DposInterfaceRewardsClaimed)
	if err := _DposInterface.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceUndelegateCanceledIterator is returned from FilterUndelegateCanceled and is used to iterate over the raw logs and unpacked data for UndelegateCanceled events raised by the DposInterface contract.
type DposInterfaceUndelegateCanceledIterator struct {
	Event *DposInterfaceUndelegateCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceUndelegateCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceUndelegateCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceUndelegateCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceUndelegateCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceUndelegateCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceUndelegateCanceled represents a UndelegateCanceled event raised by the DposInterface contract.
type DposInterfaceUndelegateCanceled struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateCanceled is a free log retrieval operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) FilterUndelegateCanceled(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposInterfaceUndelegateCanceledIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceUndelegateCanceledIterator{contract: _DposInterface.contract, event: "UndelegateCanceled", logs: logs, sub: sub}, nil
}

// WatchUndelegateCanceled is a free log subscription operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) WatchUndelegateCanceled(opts *bind.WatchOpts, sink chan<- *DposInterfaceUndelegateCanceled, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceUndelegateCanceled)
				if err := _DposInterface.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegateCanceled is a log parse operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) ParseUndelegateCanceled(log types.Log) (*DposInterfaceUndelegateCanceled, error) {
	event := new(DposInterfaceUndelegateCanceled)
	if err := _DposInterface.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceUndelegateConfirmedIterator is returned from FilterUndelegateConfirmed and is used to iterate over the raw logs and unpacked data for UndelegateConfirmed events raised by the DposInterface contract.
type DposInterfaceUndelegateConfirmedIterator struct {
	Event *DposInterfaceUndelegateConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceUndelegateConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceUndelegateConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceUndelegateConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceUndelegateConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceUndelegateConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceUndelegateConfirmed represents a UndelegateConfirmed event raised by the DposInterface contract.
type DposInterfaceUndelegateConfirmed struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateConfirmed is a free log retrieval operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) FilterUndelegateConfirmed(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposInterfaceUndelegateConfirmedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceUndelegateConfirmedIterator{contract: _DposInterface.contract, event: "UndelegateConfirmed", logs: logs, sub: sub}, nil
}

// WatchUndelegateConfirmed is a free log subscription operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) WatchUndelegateConfirmed(opts *bind.WatchOpts, sink chan<- *DposInterfaceUndelegateConfirmed, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceUndelegateConfirmed)
				if err := _DposInterface.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegateConfirmed is a log parse operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) ParseUndelegateConfirmed(log types.Log) (*DposInterfaceUndelegateConfirmed, error) {
	event := new(DposInterfaceUndelegateConfirmed)
	if err := _DposInterface.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the DposInterface contract.
type DposInterfaceUndelegatedIterator struct {
	Event *DposInterfaceUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceUndelegated represents a Undelegated event raised by the DposInterface contract.
type DposInterfaceUndelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*DposInterfaceUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceUndelegatedIterator{contract: _DposInterface.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *DposInterfaceUndelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceUndelegated)
				if err := _DposInterface.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_DposInterface *DposInterfaceFilterer) ParseUndelegated(log types.Log) (*DposInterfaceUndelegated, error) {
	event := new(DposInterfaceUndelegated)
	if err := _DposInterface.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceValidatorInfoSetIterator is returned from FilterValidatorInfoSet and is used to iterate over the raw logs and unpacked data for ValidatorInfoSet events raised by the DposInterface contract.
type DposInterfaceValidatorInfoSetIterator struct {
	Event *DposInterfaceValidatorInfoSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceValidatorInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceValidatorInfoSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceValidatorInfoSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceValidatorInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceValidatorInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceValidatorInfoSet represents a ValidatorInfoSet event raised by the DposInterface contract.
type DposInterfaceValidatorInfoSet struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorInfoSet is a free log retrieval operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_DposInterface *DposInterfaceFilterer) FilterValidatorInfoSet(opts *bind.FilterOpts, validator []common.Address) (*DposInterfaceValidatorInfoSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceValidatorInfoSetIterator{contract: _DposInterface.contract, event: "ValidatorInfoSet", logs: logs, sub: sub}, nil
}

// WatchValidatorInfoSet is a free log subscription operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_DposInterface *DposInterfaceFilterer) WatchValidatorInfoSet(opts *bind.WatchOpts, sink chan<- *DposInterfaceValidatorInfoSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceValidatorInfoSet)
				if err := _DposInterface.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorInfoSet is a log parse operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_DposInterface *DposInterfaceFilterer) ParseValidatorInfoSet(log types.Log) (*DposInterfaceValidatorInfoSet, error) {
	event := new(DposInterfaceValidatorInfoSet)
	if err := _DposInterface.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DposInterfaceValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the DposInterface contract.
type DposInterfaceValidatorRegisteredIterator struct {
	Event *DposInterfaceValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DposInterfaceValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DposInterfaceValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DposInterfaceValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DposInterfaceValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DposInterfaceValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DposInterfaceValidatorRegistered represents a ValidatorRegistered event raised by the DposInterface contract.
type DposInterfaceValidatorRegistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_DposInterface *DposInterfaceFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*DposInterfaceValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DposInterfaceValidatorRegisteredIterator{contract: _DposInterface.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_DposInterface *DposInterfaceFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *DposInterfaceValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DposInterface.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DposInterfaceValidatorRegistered)
				if err := _DposInterface.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_DposInterface *DposInterfaceFilterer) ParseValidatorRegistered(log types.Log) (*DposInterfaceValidatorRegistered, error) {
	event := new(DposInterfaceValidatorRegistered)
	if err := _DposInterface.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

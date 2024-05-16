// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pillar_block_interface

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

// PillarBlockInterfaceMetaData contains all meta data concerning the PillarBlockInterface contract.
var PillarBlockInterfaceMetaData = &bind.MetaData{
	ABI: "[]",
}

// PillarBlockInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use PillarBlockInterfaceMetaData.ABI instead.
var PillarBlockInterfaceABI = PillarBlockInterfaceMetaData.ABI

// PillarBlockInterface is an auto generated Go binding around an Ethereum contract.
type PillarBlockInterface struct {
	PillarBlockInterfaceCaller     // Read-only binding to the contract
	PillarBlockInterfaceTransactor // Write-only binding to the contract
	PillarBlockInterfaceFilterer   // Log filterer for contract events
}

// PillarBlockInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PillarBlockInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PillarBlockInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PillarBlockInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PillarBlockInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PillarBlockInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PillarBlockInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PillarBlockInterfaceSession struct {
	Contract     *PillarBlockInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PillarBlockInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PillarBlockInterfaceCallerSession struct {
	Contract *PillarBlockInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PillarBlockInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PillarBlockInterfaceTransactorSession struct {
	Contract     *PillarBlockInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PillarBlockInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PillarBlockInterfaceRaw struct {
	Contract *PillarBlockInterface // Generic contract binding to access the raw methods on
}

// PillarBlockInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PillarBlockInterfaceCallerRaw struct {
	Contract *PillarBlockInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PillarBlockInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PillarBlockInterfaceTransactorRaw struct {
	Contract *PillarBlockInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPillarBlockInterface creates a new instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterface(address common.Address, backend bind.ContractBackend) (*PillarBlockInterface, error) {
	contract, err := bindPillarBlockInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterface{PillarBlockInterfaceCaller: PillarBlockInterfaceCaller{contract: contract}, PillarBlockInterfaceTransactor: PillarBlockInterfaceTransactor{contract: contract}, PillarBlockInterfaceFilterer: PillarBlockInterfaceFilterer{contract: contract}}, nil
}

// NewPillarBlockInterfaceCaller creates a new read-only instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PillarBlockInterfaceCaller, error) {
	contract, err := bindPillarBlockInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterfaceCaller{contract: contract}, nil
}

// NewPillarBlockInterfaceTransactor creates a new write-only instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PillarBlockInterfaceTransactor, error) {
	contract, err := bindPillarBlockInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterfaceTransactor{contract: contract}, nil
}

// NewPillarBlockInterfaceFilterer creates a new log filterer instance of PillarBlockInterface, bound to a specific deployed contract.
func NewPillarBlockInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PillarBlockInterfaceFilterer, error) {
	contract, err := bindPillarBlockInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PillarBlockInterfaceFilterer{contract: contract}, nil
}

// bindPillarBlockInterface binds a generic wrapper to an already deployed contract.
func bindPillarBlockInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PillarBlockInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PillarBlockInterface *PillarBlockInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PillarBlockInterface.Contract.PillarBlockInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PillarBlockInterface *PillarBlockInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.PillarBlockInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PillarBlockInterface *PillarBlockInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.PillarBlockInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PillarBlockInterface *PillarBlockInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PillarBlockInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PillarBlockInterface *PillarBlockInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PillarBlockInterface *PillarBlockInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PillarBlockInterface.Contract.contract.Transact(opts, method, params...)
}

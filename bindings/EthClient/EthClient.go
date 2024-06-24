// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EthClient

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

// EthClientMetaData contains all meta data concerning the EthClient contract.
var EthClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeRootKeyByEpoch\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeRootsMappingPosition\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"client\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBeaconLightClient\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethBridgeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizedBridgeRoot\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_client\",\"type\":\"address\",\"internalType\":\"contractBeaconLightClient\"},{\"name\":\"_eth_bridge_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processBridgeRoot\",\"inputs\":[{\"name\":\"account_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"storage_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgeRootProcessed\",\"inputs\":[{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidBridgeRoot\",\"inputs\":[{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// EthClientABI is the input ABI used to generate the binding from.
// Deprecated: Use EthClientMetaData.ABI instead.
var EthClientABI = EthClientMetaData.ABI

// EthClient is an auto generated Go binding around an Ethereum contract.
type EthClient struct {
	EthClientCaller     // Read-only binding to the contract
	EthClientTransactor // Write-only binding to the contract
	EthClientFilterer   // Log filterer for contract events
}

// EthClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthClientSession struct {
	Contract     *EthClient        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthClientCallerSession struct {
	Contract *EthClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthClientTransactorSession struct {
	Contract     *EthClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthClientRaw struct {
	Contract *EthClient // Generic contract binding to access the raw methods on
}

// EthClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthClientCallerRaw struct {
	Contract *EthClientCaller // Generic read-only contract binding to access the raw methods on
}

// EthClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthClientTransactorRaw struct {
	Contract *EthClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthClient creates a new instance of EthClient, bound to a specific deployed contract.
func NewEthClient(address common.Address, backend bind.ContractBackend) (*EthClient, error) {
	contract, err := bindEthClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthClient{EthClientCaller: EthClientCaller{contract: contract}, EthClientTransactor: EthClientTransactor{contract: contract}, EthClientFilterer: EthClientFilterer{contract: contract}}, nil
}

// NewEthClientCaller creates a new read-only instance of EthClient, bound to a specific deployed contract.
func NewEthClientCaller(address common.Address, caller bind.ContractCaller) (*EthClientCaller, error) {
	contract, err := bindEthClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthClientCaller{contract: contract}, nil
}

// NewEthClientTransactor creates a new write-only instance of EthClient, bound to a specific deployed contract.
func NewEthClientTransactor(address common.Address, transactor bind.ContractTransactor) (*EthClientTransactor, error) {
	contract, err := bindEthClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthClientTransactor{contract: contract}, nil
}

// NewEthClientFilterer creates a new log filterer instance of EthClient, bound to a specific deployed contract.
func NewEthClientFilterer(address common.Address, filterer bind.ContractFilterer) (*EthClientFilterer, error) {
	contract, err := bindEthClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthClientFilterer{contract: contract}, nil
}

// bindEthClient binds a generic wrapper to an already deployed contract.
func bindEthClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClient *EthClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClient.Contract.EthClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClient *EthClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClient.Contract.EthClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClient *EthClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClient.Contract.EthClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthClient *EthClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthClient *EthClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthClient *EthClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthClient.Contract.contract.Transact(opts, method, params...)
}

// BridgeRootKeyByEpoch is a free data retrieval call binding the contract method 0xbdac44d1.
//
// Solidity: function bridgeRootKeyByEpoch(uint256 epoch) view returns(bytes32)
func (_EthClient *EthClientCaller) BridgeRootKeyByEpoch(opts *bind.CallOpts, epoch *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "bridgeRootKeyByEpoch", epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BridgeRootKeyByEpoch is a free data retrieval call binding the contract method 0xbdac44d1.
//
// Solidity: function bridgeRootKeyByEpoch(uint256 epoch) view returns(bytes32)
func (_EthClient *EthClientSession) BridgeRootKeyByEpoch(epoch *big.Int) ([32]byte, error) {
	return _EthClient.Contract.BridgeRootKeyByEpoch(&_EthClient.CallOpts, epoch)
}

// BridgeRootKeyByEpoch is a free data retrieval call binding the contract method 0xbdac44d1.
//
// Solidity: function bridgeRootKeyByEpoch(uint256 epoch) view returns(bytes32)
func (_EthClient *EthClientCallerSession) BridgeRootKeyByEpoch(epoch *big.Int) ([32]byte, error) {
	return _EthClient.Contract.BridgeRootKeyByEpoch(&_EthClient.CallOpts, epoch)
}

// BridgeRootsMappingPosition is a free data retrieval call binding the contract method 0xd4544d91.
//
// Solidity: function bridgeRootsMappingPosition() view returns(bytes32)
func (_EthClient *EthClientCaller) BridgeRootsMappingPosition(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "bridgeRootsMappingPosition")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BridgeRootsMappingPosition is a free data retrieval call binding the contract method 0xd4544d91.
//
// Solidity: function bridgeRootsMappingPosition() view returns(bytes32)
func (_EthClient *EthClientSession) BridgeRootsMappingPosition() ([32]byte, error) {
	return _EthClient.Contract.BridgeRootsMappingPosition(&_EthClient.CallOpts)
}

// BridgeRootsMappingPosition is a free data retrieval call binding the contract method 0xd4544d91.
//
// Solidity: function bridgeRootsMappingPosition() view returns(bytes32)
func (_EthClient *EthClientCallerSession) BridgeRootsMappingPosition() ([32]byte, error) {
	return _EthClient.Contract.BridgeRootsMappingPosition(&_EthClient.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_EthClient *EthClientCaller) Client(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "client")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_EthClient *EthClientSession) Client() (common.Address, error) {
	return _EthClient.Contract.Client(&_EthClient.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_EthClient *EthClientCallerSession) Client() (common.Address, error) {
	return _EthClient.Contract.Client(&_EthClient.CallOpts)
}

// EthBridgeAddress is a free data retrieval call binding the contract method 0xf7b4fb57.
//
// Solidity: function ethBridgeAddress() view returns(address)
func (_EthClient *EthClientCaller) EthBridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "ethBridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthBridgeAddress is a free data retrieval call binding the contract method 0xf7b4fb57.
//
// Solidity: function ethBridgeAddress() view returns(address)
func (_EthClient *EthClientSession) EthBridgeAddress() (common.Address, error) {
	return _EthClient.Contract.EthBridgeAddress(&_EthClient.CallOpts)
}

// EthBridgeAddress is a free data retrieval call binding the contract method 0xf7b4fb57.
//
// Solidity: function ethBridgeAddress() view returns(address)
func (_EthClient *EthClientCallerSession) EthBridgeAddress() (common.Address, error) {
	return _EthClient.Contract.EthBridgeAddress(&_EthClient.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_EthClient *EthClientCaller) GetFinalizedBridgeRoot(opts *bind.CallOpts, epoch *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "getFinalizedBridgeRoot", epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_EthClient *EthClientSession) GetFinalizedBridgeRoot(epoch *big.Int) ([32]byte, error) {
	return _EthClient.Contract.GetFinalizedBridgeRoot(&_EthClient.CallOpts, epoch)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_EthClient *EthClientCallerSession) GetFinalizedBridgeRoot(epoch *big.Int) ([32]byte, error) {
	return _EthClient.Contract.GetFinalizedBridgeRoot(&_EthClient.CallOpts, epoch)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_EthClient *EthClientCaller) GetMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "getMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_EthClient *EthClientSession) GetMerkleRoot() ([32]byte, error) {
	return _EthClient.Contract.GetMerkleRoot(&_EthClient.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_EthClient *EthClientCallerSession) GetMerkleRoot() ([32]byte, error) {
	return _EthClient.Contract.GetMerkleRoot(&_EthClient.CallOpts)
}

// LastEpoch is a free data retrieval call binding the contract method 0x06a4c983.
//
// Solidity: function lastEpoch() view returns(uint256)
func (_EthClient *EthClientCaller) LastEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "lastEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEpoch is a free data retrieval call binding the contract method 0x06a4c983.
//
// Solidity: function lastEpoch() view returns(uint256)
func (_EthClient *EthClientSession) LastEpoch() (*big.Int, error) {
	return _EthClient.Contract.LastEpoch(&_EthClient.CallOpts)
}

// LastEpoch is a free data retrieval call binding the contract method 0x06a4c983.
//
// Solidity: function lastEpoch() view returns(uint256)
func (_EthClient *EthClientCallerSession) LastEpoch() (*big.Int, error) {
	return _EthClient.Contract.LastEpoch(&_EthClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClient *EthClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClient *EthClientSession) Owner() (common.Address, error) {
	return _EthClient.Contract.Owner(&_EthClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthClient *EthClientCallerSession) Owner() (common.Address, error) {
	return _EthClient.Contract.Owner(&_EthClient.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _client, address _eth_bridge_address) returns()
func (_EthClient *EthClientTransactor) Initialize(opts *bind.TransactOpts, _client common.Address, _eth_bridge_address common.Address) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "initialize", _client, _eth_bridge_address)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _client, address _eth_bridge_address) returns()
func (_EthClient *EthClientSession) Initialize(_client common.Address, _eth_bridge_address common.Address) (*types.Transaction, error) {
	return _EthClient.Contract.Initialize(&_EthClient.TransactOpts, _client, _eth_bridge_address)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _client, address _eth_bridge_address) returns()
func (_EthClient *EthClientTransactorSession) Initialize(_client common.Address, _eth_bridge_address common.Address) (*types.Transaction, error) {
	return _EthClient.Contract.Initialize(&_EthClient.TransactOpts, _client, _eth_bridge_address)
}

// ProcessBridgeRoot is a paid mutator transaction binding the contract method 0x99b7ff1b.
//
// Solidity: function processBridgeRoot(bytes[] account_proof, bytes[] storage_proof) returns()
func (_EthClient *EthClientTransactor) ProcessBridgeRoot(opts *bind.TransactOpts, account_proof [][]byte, storage_proof [][]byte) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "processBridgeRoot", account_proof, storage_proof)
}

// ProcessBridgeRoot is a paid mutator transaction binding the contract method 0x99b7ff1b.
//
// Solidity: function processBridgeRoot(bytes[] account_proof, bytes[] storage_proof) returns()
func (_EthClient *EthClientSession) ProcessBridgeRoot(account_proof [][]byte, storage_proof [][]byte) (*types.Transaction, error) {
	return _EthClient.Contract.ProcessBridgeRoot(&_EthClient.TransactOpts, account_proof, storage_proof)
}

// ProcessBridgeRoot is a paid mutator transaction binding the contract method 0x99b7ff1b.
//
// Solidity: function processBridgeRoot(bytes[] account_proof, bytes[] storage_proof) returns()
func (_EthClient *EthClientTransactorSession) ProcessBridgeRoot(account_proof [][]byte, storage_proof [][]byte) (*types.Transaction, error) {
	return _EthClient.Contract.ProcessBridgeRoot(&_EthClient.TransactOpts, account_proof, storage_proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthClient *EthClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthClient *EthClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthClient.Contract.RenounceOwnership(&_EthClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthClient *EthClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthClient.Contract.RenounceOwnership(&_EthClient.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthClient *EthClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthClient *EthClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthClient.Contract.TransferOwnership(&_EthClient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthClient *EthClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthClient.Contract.TransferOwnership(&_EthClient.TransactOpts, newOwner)
}

// EthClientBridgeRootProcessedIterator is returned from FilterBridgeRootProcessed and is used to iterate over the raw logs and unpacked data for BridgeRootProcessed events raised by the EthClient contract.
type EthClientBridgeRootProcessedIterator struct {
	Event *EthClientBridgeRootProcessed // Event containing the contract specifics and raw log

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
func (it *EthClientBridgeRootProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClientBridgeRootProcessed)
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
		it.Event = new(EthClientBridgeRootProcessed)
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
func (it *EthClientBridgeRootProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClientBridgeRootProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClientBridgeRootProcessed represents a BridgeRootProcessed event raised by the EthClient contract.
type EthClientBridgeRootProcessed struct {
	BridgeRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgeRootProcessed is a free log retrieval operation binding the contract event 0xa7b0442eda88383df55bc140c0db1af8e39d1ad2c5183a7e44aae7bbf91842a3.
//
// Solidity: event BridgeRootProcessed(bytes32 indexed bridgeRoot)
func (_EthClient *EthClientFilterer) FilterBridgeRootProcessed(opts *bind.FilterOpts, bridgeRoot [][32]byte) (*EthClientBridgeRootProcessedIterator, error) {

	var bridgeRootRule []interface{}
	for _, bridgeRootItem := range bridgeRoot {
		bridgeRootRule = append(bridgeRootRule, bridgeRootItem)
	}

	logs, sub, err := _EthClient.contract.FilterLogs(opts, "BridgeRootProcessed", bridgeRootRule)
	if err != nil {
		return nil, err
	}
	return &EthClientBridgeRootProcessedIterator{contract: _EthClient.contract, event: "BridgeRootProcessed", logs: logs, sub: sub}, nil
}

// WatchBridgeRootProcessed is a free log subscription operation binding the contract event 0xa7b0442eda88383df55bc140c0db1af8e39d1ad2c5183a7e44aae7bbf91842a3.
//
// Solidity: event BridgeRootProcessed(bytes32 indexed bridgeRoot)
func (_EthClient *EthClientFilterer) WatchBridgeRootProcessed(opts *bind.WatchOpts, sink chan<- *EthClientBridgeRootProcessed, bridgeRoot [][32]byte) (event.Subscription, error) {

	var bridgeRootRule []interface{}
	for _, bridgeRootItem := range bridgeRoot {
		bridgeRootRule = append(bridgeRootRule, bridgeRootItem)
	}

	logs, sub, err := _EthClient.contract.WatchLogs(opts, "BridgeRootProcessed", bridgeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClientBridgeRootProcessed)
				if err := _EthClient.contract.UnpackLog(event, "BridgeRootProcessed", log); err != nil {
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

// ParseBridgeRootProcessed is a log parse operation binding the contract event 0xa7b0442eda88383df55bc140c0db1af8e39d1ad2c5183a7e44aae7bbf91842a3.
//
// Solidity: event BridgeRootProcessed(bytes32 indexed bridgeRoot)
func (_EthClient *EthClientFilterer) ParseBridgeRootProcessed(log types.Log) (*EthClientBridgeRootProcessed, error) {
	event := new(EthClientBridgeRootProcessed)
	if err := _EthClient.contract.UnpackLog(event, "BridgeRootProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EthClient contract.
type EthClientInitializedIterator struct {
	Event *EthClientInitialized // Event containing the contract specifics and raw log

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
func (it *EthClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClientInitialized)
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
		it.Event = new(EthClientInitialized)
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
func (it *EthClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClientInitialized represents a Initialized event raised by the EthClient contract.
type EthClientInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EthClient *EthClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*EthClientInitializedIterator, error) {

	logs, sub, err := _EthClient.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EthClientInitializedIterator{contract: _EthClient.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EthClient *EthClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EthClientInitialized) (event.Subscription, error) {

	logs, sub, err := _EthClient.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClientInitialized)
				if err := _EthClient.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EthClient *EthClientFilterer) ParseInitialized(log types.Log) (*EthClientInitialized, error) {
	event := new(EthClientInitialized)
	if err := _EthClient.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthClient contract.
type EthClientOwnershipTransferredIterator struct {
	Event *EthClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthClientOwnershipTransferred)
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
		it.Event = new(EthClientOwnershipTransferred)
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
func (it *EthClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthClientOwnershipTransferred represents a OwnershipTransferred event raised by the EthClient contract.
type EthClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthClient *EthClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthClient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthClientOwnershipTransferredIterator{contract: _EthClient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthClient *EthClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthClient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthClientOwnershipTransferred)
				if err := _EthClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthClient *EthClientFilterer) ParseOwnershipTransferred(log types.Log) (*EthClientOwnershipTransferred, error) {
	event := new(EthClientOwnershipTransferred)
	if err := _EthClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

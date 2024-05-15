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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_client\",\"type\":\"address\",\"internalType\":\"contractBeaconLightClient\"},{\"name\":\"_eth_bridge_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeRootKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"client\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBeaconLightClient\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethBridgeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizedBridgeRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processBridgeRoot\",\"inputs\":[{\"name\":\"account_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"storage_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"refundAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidBridgeRoot\",\"inputs\":[{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
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

// BridgeRootKey is a free data retrieval call binding the contract method 0x9760948a.
//
// Solidity: function bridgeRootKey() view returns(bytes32)
func (_EthClient *EthClientCaller) BridgeRootKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "bridgeRootKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BridgeRootKey is a free data retrieval call binding the contract method 0x9760948a.
//
// Solidity: function bridgeRootKey() view returns(bytes32)
func (_EthClient *EthClientSession) BridgeRootKey() ([32]byte, error) {
	return _EthClient.Contract.BridgeRootKey(&_EthClient.CallOpts)
}

// BridgeRootKey is a free data retrieval call binding the contract method 0x9760948a.
//
// Solidity: function bridgeRootKey() view returns(bytes32)
func (_EthClient *EthClientCallerSession) BridgeRootKey() ([32]byte, error) {
	return _EthClient.Contract.BridgeRootKey(&_EthClient.CallOpts)
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

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_EthClient *EthClientCaller) GetFinalizedBridgeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "getFinalizedBridgeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_EthClient *EthClientSession) GetFinalizedBridgeRoot() ([32]byte, error) {
	return _EthClient.Contract.GetFinalizedBridgeRoot(&_EthClient.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_EthClient *EthClientCallerSession) GetFinalizedBridgeRoot() ([32]byte, error) {
	return _EthClient.Contract.GetFinalizedBridgeRoot(&_EthClient.CallOpts)
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

// RefundAmount is a free data retrieval call binding the contract method 0xad33513f.
//
// Solidity: function refundAmount() view returns(uint256)
func (_EthClient *EthClientCaller) RefundAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "refundAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundAmount is a free data retrieval call binding the contract method 0xad33513f.
//
// Solidity: function refundAmount() view returns(uint256)
func (_EthClient *EthClientSession) RefundAmount() (*big.Int, error) {
	return _EthClient.Contract.RefundAmount(&_EthClient.CallOpts)
}

// RefundAmount is a free data retrieval call binding the contract method 0xad33513f.
//
// Solidity: function refundAmount() view returns(uint256)
func (_EthClient *EthClientCallerSession) RefundAmount() (*big.Int, error) {
	return _EthClient.Contract.RefundAmount(&_EthClient.CallOpts)
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

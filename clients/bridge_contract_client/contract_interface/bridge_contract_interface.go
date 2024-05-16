// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge_contract_interface

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

// SharedStructsBridgeState is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsBridgeState struct {
	Epoch  *big.Int
	States []SharedStructsStateWithAddress
}

// SharedStructsContractStateHash is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsContractStateHash struct {
	ContractAddress common.Address
	StateHash       [32]byte
}

// SharedStructsStateWithAddress is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsStateWithAddress struct {
	ContractAddress common.Address
	State           []byte
}

// SharedStructsStateWithProof is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsStateWithProof struct {
	State       SharedStructsBridgeState
	StateHashes []SharedStructsContractStateHash
}

// BridgeContractInterfaceMetaData contains all meta data concerning the BridgeContractInterface contract.
var BridgeContractInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStateHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidStateHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizationInterval\",\"type\":\"uint256\"}],\"name\":\"NotEnoughBlocksPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"name\":\"NotSuccessiveEpochs\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"}],\"name\":\"StateNotMatchingBridgeRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedContractAddress\",\"type\":\"address\"}],\"name\":\"UnmatchingContractAddresses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"UnregisteredContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressCannotBeRegistered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"ConnectorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"}],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"connector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"StateApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"appliedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"internalType\":\"structSharedStructs.StateWithAddress[]\",\"name\":\"states\",\"type\":\"tuple[]\"}],\"internalType\":\"structSharedStructs.BridgeState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"internalType\":\"structSharedStructs.ContractStateHash[]\",\"name\":\"state_hashes\",\"type\":\"tuple[]\"}],\"internalType\":\"structSharedStructs.StateWithProof\",\"name\":\"state_with_proof\",\"type\":\"tuple\"}],\"name\":\"applyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"connectors\",\"outputs\":[{\"internalType\":\"contractIBridgeConnector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizationInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateWithProof\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"internalType\":\"structSharedStructs.StateWithAddress[]\",\"name\":\"states\",\"type\":\"tuple[]\"}],\"internalType\":\"structSharedStructs.BridgeState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"internalType\":\"structSharedStructs.ContractStateHash[]\",\"name\":\"state_hashes\",\"type\":\"tuple[]\"}],\"internalType\":\"structSharedStructs.StateWithProof\",\"name\":\"ret\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lightClient\",\"outputs\":[{\"internalType\":\"contractIBridgeLightClient\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"localAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridgeConnector\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"registerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registeredTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_finalizationInterval\",\"type\":\"uint256\"}],\"name\":\"setFinalizationInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shouldFinalizeEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BridgeContractInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeContractInterfaceMetaData.ABI instead.
var BridgeContractInterfaceABI = BridgeContractInterfaceMetaData.ABI

// BridgeContractInterface is an auto generated Go binding around an Ethereum contract.
type BridgeContractInterface struct {
	BridgeContractInterfaceCaller     // Read-only binding to the contract
	BridgeContractInterfaceTransactor // Write-only binding to the contract
	BridgeContractInterfaceFilterer   // Log filterer for contract events
}

// BridgeContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeContractInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeContractInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeContractInterfaceSession struct {
	Contract     *BridgeContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgeContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeContractInterfaceCallerSession struct {
	Contract *BridgeContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BridgeContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeContractInterfaceTransactorSession struct {
	Contract     *BridgeContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BridgeContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeContractInterfaceRaw struct {
	Contract *BridgeContractInterface // Generic contract binding to access the raw methods on
}

// BridgeContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeContractInterfaceCallerRaw struct {
	Contract *BridgeContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeContractInterfaceTransactorRaw struct {
	Contract *BridgeContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeContractInterface creates a new instance of BridgeContractInterface, bound to a specific deployed contract.
func NewBridgeContractInterface(address common.Address, backend bind.ContractBackend) (*BridgeContractInterface, error) {
	contract, err := bindBridgeContractInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterface{BridgeContractInterfaceCaller: BridgeContractInterfaceCaller{contract: contract}, BridgeContractInterfaceTransactor: BridgeContractInterfaceTransactor{contract: contract}, BridgeContractInterfaceFilterer: BridgeContractInterfaceFilterer{contract: contract}}, nil
}

// NewBridgeContractInterfaceCaller creates a new read-only instance of BridgeContractInterface, bound to a specific deployed contract.
func NewBridgeContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*BridgeContractInterfaceCaller, error) {
	contract, err := bindBridgeContractInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceCaller{contract: contract}, nil
}

// NewBridgeContractInterfaceTransactor creates a new write-only instance of BridgeContractInterface, bound to a specific deployed contract.
func NewBridgeContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeContractInterfaceTransactor, error) {
	contract, err := bindBridgeContractInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceTransactor{contract: contract}, nil
}

// NewBridgeContractInterfaceFilterer creates a new log filterer instance of BridgeContractInterface, bound to a specific deployed contract.
func NewBridgeContractInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeContractInterfaceFilterer, error) {
	contract, err := bindBridgeContractInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceFilterer{contract: contract}, nil
}

// bindBridgeContractInterface binds a generic wrapper to an already deployed contract.
func bindBridgeContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeContractInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeContractInterface *BridgeContractInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeContractInterface.Contract.BridgeContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeContractInterface *BridgeContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.BridgeContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeContractInterface *BridgeContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.BridgeContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeContractInterface *BridgeContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeContractInterface *BridgeContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeContractInterface *BridgeContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeContractInterface *BridgeContractInterfaceSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeContractInterface.Contract.UPGRADEINTERFACEVERSION(&_BridgeContractInterface.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeContractInterface.Contract.UPGRADEINTERFACEVERSION(&_BridgeContractInterface.CallOpts)
}

// AppliedEpoch is a free data retrieval call binding the contract method 0x35add093.
//
// Solidity: function appliedEpoch() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) AppliedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "appliedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppliedEpoch is a free data retrieval call binding the contract method 0x35add093.
//
// Solidity: function appliedEpoch() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceSession) AppliedEpoch() (*big.Int, error) {
	return _BridgeContractInterface.Contract.AppliedEpoch(&_BridgeContractInterface.CallOpts)
}

// AppliedEpoch is a free data retrieval call binding the contract method 0x35add093.
//
// Solidity: function appliedEpoch() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) AppliedEpoch() (*big.Int, error) {
	return _BridgeContractInterface.Contract.AppliedEpoch(&_BridgeContractInterface.CallOpts)
}

// BridgeRoot is a free data retrieval call binding the contract method 0x177a0c2c.
//
// Solidity: function bridgeRoot() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) BridgeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "bridgeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BridgeRoot is a free data retrieval call binding the contract method 0x177a0c2c.
//
// Solidity: function bridgeRoot() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceSession) BridgeRoot() ([32]byte, error) {
	return _BridgeContractInterface.Contract.BridgeRoot(&_BridgeContractInterface.CallOpts)
}

// BridgeRoot is a free data retrieval call binding the contract method 0x177a0c2c.
//
// Solidity: function bridgeRoot() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) BridgeRoot() ([32]byte, error) {
	return _BridgeContractInterface.Contract.BridgeRoot(&_BridgeContractInterface.CallOpts)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) Connectors(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "connectors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceSession) Connectors(arg0 common.Address) (common.Address, error) {
	return _BridgeContractInterface.Contract.Connectors(&_BridgeContractInterface.CallOpts, arg0)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) Connectors(arg0 common.Address) (common.Address, error) {
	return _BridgeContractInterface.Contract.Connectors(&_BridgeContractInterface.CallOpts, arg0)
}

// FinalizationInterval is a free data retrieval call binding the contract method 0xbb04c6fc.
//
// Solidity: function finalizationInterval() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) FinalizationInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "finalizationInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizationInterval is a free data retrieval call binding the contract method 0xbb04c6fc.
//
// Solidity: function finalizationInterval() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceSession) FinalizationInterval() (*big.Int, error) {
	return _BridgeContractInterface.Contract.FinalizationInterval(&_BridgeContractInterface.CallOpts)
}

// FinalizationInterval is a free data retrieval call binding the contract method 0xbb04c6fc.
//
// Solidity: function finalizationInterval() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) FinalizationInterval() (*big.Int, error) {
	return _BridgeContractInterface.Contract.FinalizationInterval(&_BridgeContractInterface.CallOpts)
}

// FinalizedEpoch is a free data retrieval call binding the contract method 0x6bfa7398.
//
// Solidity: function finalizedEpoch() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) FinalizedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "finalizedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizedEpoch is a free data retrieval call binding the contract method 0x6bfa7398.
//
// Solidity: function finalizedEpoch() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceSession) FinalizedEpoch() (*big.Int, error) {
	return _BridgeContractInterface.Contract.FinalizedEpoch(&_BridgeContractInterface.CallOpts)
}

// FinalizedEpoch is a free data retrieval call binding the contract method 0x6bfa7398.
//
// Solidity: function finalizedEpoch() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) FinalizedEpoch() (*big.Int, error) {
	return _BridgeContractInterface.Contract.FinalizedEpoch(&_BridgeContractInterface.CallOpts)
}

// GetBridgeRoot is a free data retrieval call binding the contract method 0x695a253f.
//
// Solidity: function getBridgeRoot() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) GetBridgeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "getBridgeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBridgeRoot is a free data retrieval call binding the contract method 0x695a253f.
//
// Solidity: function getBridgeRoot() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceSession) GetBridgeRoot() ([32]byte, error) {
	return _BridgeContractInterface.Contract.GetBridgeRoot(&_BridgeContractInterface.CallOpts)
}

// GetBridgeRoot is a free data retrieval call binding the contract method 0x695a253f.
//
// Solidity: function getBridgeRoot() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) GetBridgeRoot() ([32]byte, error) {
	return _BridgeContractInterface.Contract.GetBridgeRoot(&_BridgeContractInterface.CallOpts)
}

// GetStateWithProof is a free data retrieval call binding the contract method 0xfe65d463.
//
// Solidity: function getStateWithProof() view returns(((uint256,(address,bytes)[]),(address,bytes32)[]) ret)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) GetStateWithProof(opts *bind.CallOpts) (SharedStructsStateWithProof, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "getStateWithProof")

	if err != nil {
		return *new(SharedStructsStateWithProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsStateWithProof)).(*SharedStructsStateWithProof)

	return out0, err

}

// GetStateWithProof is a free data retrieval call binding the contract method 0xfe65d463.
//
// Solidity: function getStateWithProof() view returns(((uint256,(address,bytes)[]),(address,bytes32)[]) ret)
func (_BridgeContractInterface *BridgeContractInterfaceSession) GetStateWithProof() (SharedStructsStateWithProof, error) {
	return _BridgeContractInterface.Contract.GetStateWithProof(&_BridgeContractInterface.CallOpts)
}

// GetStateWithProof is a free data retrieval call binding the contract method 0xfe65d463.
//
// Solidity: function getStateWithProof() view returns(((uint256,(address,bytes)[]),(address,bytes32)[]) ret)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) GetStateWithProof() (SharedStructsStateWithProof, error) {
	return _BridgeContractInterface.Contract.GetStateWithProof(&_BridgeContractInterface.CallOpts)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0xae1da0b5.
//
// Solidity: function lastFinalizedBlock() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) LastFinalizedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "lastFinalizedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0xae1da0b5.
//
// Solidity: function lastFinalizedBlock() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceSession) LastFinalizedBlock() (*big.Int, error) {
	return _BridgeContractInterface.Contract.LastFinalizedBlock(&_BridgeContractInterface.CallOpts)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0xae1da0b5.
//
// Solidity: function lastFinalizedBlock() view returns(uint256)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) LastFinalizedBlock() (*big.Int, error) {
	return _BridgeContractInterface.Contract.LastFinalizedBlock(&_BridgeContractInterface.CallOpts)
}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) LightClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "lightClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceSession) LightClient() (common.Address, error) {
	return _BridgeContractInterface.Contract.LightClient(&_BridgeContractInterface.CallOpts)
}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) LightClient() (common.Address, error) {
	return _BridgeContractInterface.Contract.LightClient(&_BridgeContractInterface.CallOpts)
}

// LocalAddress is a free data retrieval call binding the contract method 0x76081bd5.
//
// Solidity: function localAddress(address ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) LocalAddress(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "localAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalAddress is a free data retrieval call binding the contract method 0x76081bd5.
//
// Solidity: function localAddress(address ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceSession) LocalAddress(arg0 common.Address) (common.Address, error) {
	return _BridgeContractInterface.Contract.LocalAddress(&_BridgeContractInterface.CallOpts, arg0)
}

// LocalAddress is a free data retrieval call binding the contract method 0x76081bd5.
//
// Solidity: function localAddress(address ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) LocalAddress(arg0 common.Address) (common.Address, error) {
	return _BridgeContractInterface.Contract.LocalAddress(&_BridgeContractInterface.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceSession) Owner() (common.Address, error) {
	return _BridgeContractInterface.Contract.Owner(&_BridgeContractInterface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) Owner() (common.Address, error) {
	return _BridgeContractInterface.Contract.Owner(&_BridgeContractInterface.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeContractInterface.Contract.ProxiableUUID(&_BridgeContractInterface.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeContractInterface.Contract.ProxiableUUID(&_BridgeContractInterface.CallOpts)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x45466616.
//
// Solidity: function registeredTokens() view returns(address[])
func (_BridgeContractInterface *BridgeContractInterfaceCaller) RegisteredTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "registeredTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RegisteredTokens is a free data retrieval call binding the contract method 0x45466616.
//
// Solidity: function registeredTokens() view returns(address[])
func (_BridgeContractInterface *BridgeContractInterfaceSession) RegisteredTokens() ([]common.Address, error) {
	return _BridgeContractInterface.Contract.RegisteredTokens(&_BridgeContractInterface.CallOpts)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x45466616.
//
// Solidity: function registeredTokens() view returns(address[])
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) RegisteredTokens() ([]common.Address, error) {
	return _BridgeContractInterface.Contract.RegisteredTokens(&_BridgeContractInterface.CallOpts)
}

// ShouldFinalizeEpoch is a free data retrieval call binding the contract method 0x567041cd.
//
// Solidity: function shouldFinalizeEpoch() view returns(bool)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) ShouldFinalizeEpoch(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "shouldFinalizeEpoch")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldFinalizeEpoch is a free data retrieval call binding the contract method 0x567041cd.
//
// Solidity: function shouldFinalizeEpoch() view returns(bool)
func (_BridgeContractInterface *BridgeContractInterfaceSession) ShouldFinalizeEpoch() (bool, error) {
	return _BridgeContractInterface.Contract.ShouldFinalizeEpoch(&_BridgeContractInterface.CallOpts)
}

// ShouldFinalizeEpoch is a free data retrieval call binding the contract method 0x567041cd.
//
// Solidity: function shouldFinalizeEpoch() view returns(bool)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) ShouldFinalizeEpoch() (bool, error) {
	return _BridgeContractInterface.Contract.ShouldFinalizeEpoch(&_BridgeContractInterface.CallOpts)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCaller) TokenAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeContractInterface.contract.Call(opts, &out, "tokenAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceSession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _BridgeContractInterface.Contract.TokenAddresses(&_BridgeContractInterface.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_BridgeContractInterface *BridgeContractInterfaceCallerSession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _BridgeContractInterface.Contract.TokenAddresses(&_BridgeContractInterface.CallOpts, arg0)
}

// ApplyState is a paid mutator transaction binding the contract method 0x6cd50a67.
//
// Solidity: function applyState(((uint256,(address,bytes)[]),(address,bytes32)[]) state_with_proof) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactor) ApplyState(opts *bind.TransactOpts, state_with_proof SharedStructsStateWithProof) (*types.Transaction, error) {
	return _BridgeContractInterface.contract.Transact(opts, "applyState", state_with_proof)
}

// ApplyState is a paid mutator transaction binding the contract method 0x6cd50a67.
//
// Solidity: function applyState(((uint256,(address,bytes)[]),(address,bytes32)[]) state_with_proof) returns()
func (_BridgeContractInterface *BridgeContractInterfaceSession) ApplyState(state_with_proof SharedStructsStateWithProof) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.ApplyState(&_BridgeContractInterface.TransactOpts, state_with_proof)
}

// ApplyState is a paid mutator transaction binding the contract method 0x6cd50a67.
//
// Solidity: function applyState(((uint256,(address,bytes)[]),(address,bytes32)[]) state_with_proof) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactorSession) ApplyState(state_with_proof SharedStructsStateWithProof) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.ApplyState(&_BridgeContractInterface.TransactOpts, state_with_proof)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x82ae9ef7.
//
// Solidity: function finalizeEpoch() returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactor) FinalizeEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeContractInterface.contract.Transact(opts, "finalizeEpoch")
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x82ae9ef7.
//
// Solidity: function finalizeEpoch() returns()
func (_BridgeContractInterface *BridgeContractInterfaceSession) FinalizeEpoch() (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.FinalizeEpoch(&_BridgeContractInterface.TransactOpts)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x82ae9ef7.
//
// Solidity: function finalizeEpoch() returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactorSession) FinalizeEpoch() (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.FinalizeEpoch(&_BridgeContractInterface.TransactOpts)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x22a5dde4.
//
// Solidity: function registerContract(address connector) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactor) RegisterContract(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _BridgeContractInterface.contract.Transact(opts, "registerContract", connector)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x22a5dde4.
//
// Solidity: function registerContract(address connector) returns()
func (_BridgeContractInterface *BridgeContractInterfaceSession) RegisterContract(connector common.Address) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.RegisterContract(&_BridgeContractInterface.TransactOpts, connector)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x22a5dde4.
//
// Solidity: function registerContract(address connector) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactorSession) RegisterContract(connector common.Address) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.RegisterContract(&_BridgeContractInterface.TransactOpts, connector)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeContractInterface.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeContractInterface *BridgeContractInterfaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.RenounceOwnership(&_BridgeContractInterface.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.RenounceOwnership(&_BridgeContractInterface.TransactOpts)
}

// SetFinalizationInterval is a paid mutator transaction binding the contract method 0x8f084f0f.
//
// Solidity: function setFinalizationInterval(uint256 _finalizationInterval) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactor) SetFinalizationInterval(opts *bind.TransactOpts, _finalizationInterval *big.Int) (*types.Transaction, error) {
	return _BridgeContractInterface.contract.Transact(opts, "setFinalizationInterval", _finalizationInterval)
}

// SetFinalizationInterval is a paid mutator transaction binding the contract method 0x8f084f0f.
//
// Solidity: function setFinalizationInterval(uint256 _finalizationInterval) returns()
func (_BridgeContractInterface *BridgeContractInterfaceSession) SetFinalizationInterval(_finalizationInterval *big.Int) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.SetFinalizationInterval(&_BridgeContractInterface.TransactOpts, _finalizationInterval)
}

// SetFinalizationInterval is a paid mutator transaction binding the contract method 0x8f084f0f.
//
// Solidity: function setFinalizationInterval(uint256 _finalizationInterval) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactorSession) SetFinalizationInterval(_finalizationInterval *big.Int) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.SetFinalizationInterval(&_BridgeContractInterface.TransactOpts, _finalizationInterval)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeContractInterface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeContractInterface *BridgeContractInterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.TransferOwnership(&_BridgeContractInterface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.TransferOwnership(&_BridgeContractInterface.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeContractInterface.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeContractInterface *BridgeContractInterfaceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.UpgradeToAndCall(&_BridgeContractInterface.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeContractInterface *BridgeContractInterfaceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeContractInterface.Contract.UpgradeToAndCall(&_BridgeContractInterface.TransactOpts, newImplementation, data)
}

// BridgeContractInterfaceConnectorRegisteredIterator is returned from FilterConnectorRegistered and is used to iterate over the raw logs and unpacked data for ConnectorRegistered events raised by the BridgeContractInterface contract.
type BridgeContractInterfaceConnectorRegisteredIterator struct {
	Event *BridgeContractInterfaceConnectorRegistered // Event containing the contract specifics and raw log

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
func (it *BridgeContractInterfaceConnectorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeContractInterfaceConnectorRegistered)
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
		it.Event = new(BridgeContractInterfaceConnectorRegistered)
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
func (it *BridgeContractInterfaceConnectorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractInterfaceConnectorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeContractInterfaceConnectorRegistered represents a ConnectorRegistered event raised by the BridgeContractInterface contract.
type BridgeContractInterfaceConnectorRegistered struct {
	Connector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConnectorRegistered is a free log retrieval operation binding the contract event 0x069152cbb42907841c181d8a585e40e80b53f5b80110a3bd7308b035b9e57ab7.
//
// Solidity: event ConnectorRegistered(address indexed connector)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) FilterConnectorRegistered(opts *bind.FilterOpts, connector []common.Address) (*BridgeContractInterfaceConnectorRegisteredIterator, error) {

	var connectorRule []interface{}
	for _, connectorItem := range connector {
		connectorRule = append(connectorRule, connectorItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.FilterLogs(opts, "ConnectorRegistered", connectorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceConnectorRegisteredIterator{contract: _BridgeContractInterface.contract, event: "ConnectorRegistered", logs: logs, sub: sub}, nil
}

// WatchConnectorRegistered is a free log subscription operation binding the contract event 0x069152cbb42907841c181d8a585e40e80b53f5b80110a3bd7308b035b9e57ab7.
//
// Solidity: event ConnectorRegistered(address indexed connector)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) WatchConnectorRegistered(opts *bind.WatchOpts, sink chan<- *BridgeContractInterfaceConnectorRegistered, connector []common.Address) (event.Subscription, error) {

	var connectorRule []interface{}
	for _, connectorItem := range connector {
		connectorRule = append(connectorRule, connectorItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.WatchLogs(opts, "ConnectorRegistered", connectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeContractInterfaceConnectorRegistered)
				if err := _BridgeContractInterface.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
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

// ParseConnectorRegistered is a log parse operation binding the contract event 0x069152cbb42907841c181d8a585e40e80b53f5b80110a3bd7308b035b9e57ab7.
//
// Solidity: event ConnectorRegistered(address indexed connector)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) ParseConnectorRegistered(log types.Log) (*BridgeContractInterfaceConnectorRegistered, error) {
	event := new(BridgeContractInterfaceConnectorRegistered)
	if err := _BridgeContractInterface.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeContractInterfaceFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the BridgeContractInterface contract.
type BridgeContractInterfaceFinalizedIterator struct {
	Event *BridgeContractInterfaceFinalized // Event containing the contract specifics and raw log

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
func (it *BridgeContractInterfaceFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeContractInterfaceFinalized)
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
		it.Event = new(BridgeContractInterfaceFinalized)
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
func (it *BridgeContractInterfaceFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractInterfaceFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeContractInterfaceFinalized represents a Finalized event raised by the BridgeContractInterface contract.
type BridgeContractInterfaceFinalized struct {
	Epoch      *big.Int
	BridgeRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0xa05a0e9561eff1f01a29e7a680d5957bb7312e5766a8da1f494b6d6ac18031f4.
//
// Solidity: event Finalized(uint256 indexed epoch, bytes32 bridgeRoot)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) FilterFinalized(opts *bind.FilterOpts, epoch []*big.Int) (*BridgeContractInterfaceFinalizedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.FilterLogs(opts, "Finalized", epochRule)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceFinalizedIterator{contract: _BridgeContractInterface.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0xa05a0e9561eff1f01a29e7a680d5957bb7312e5766a8da1f494b6d6ac18031f4.
//
// Solidity: event Finalized(uint256 indexed epoch, bytes32 bridgeRoot)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *BridgeContractInterfaceFinalized, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.WatchLogs(opts, "Finalized", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeContractInterfaceFinalized)
				if err := _BridgeContractInterface.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// ParseFinalized is a log parse operation binding the contract event 0xa05a0e9561eff1f01a29e7a680d5957bb7312e5766a8da1f494b6d6ac18031f4.
//
// Solidity: event Finalized(uint256 indexed epoch, bytes32 bridgeRoot)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) ParseFinalized(log types.Log) (*BridgeContractInterfaceFinalized, error) {
	event := new(BridgeContractInterfaceFinalized)
	if err := _BridgeContractInterface.contract.UnpackLog(event, "Finalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeContractInterfaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeContractInterface contract.
type BridgeContractInterfaceInitializedIterator struct {
	Event *BridgeContractInterfaceInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeContractInterfaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeContractInterfaceInitialized)
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
		it.Event = new(BridgeContractInterfaceInitialized)
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
func (it *BridgeContractInterfaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractInterfaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeContractInterfaceInitialized represents a Initialized event raised by the BridgeContractInterface contract.
type BridgeContractInterfaceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeContractInterfaceInitializedIterator, error) {

	logs, sub, err := _BridgeContractInterface.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceInitializedIterator{contract: _BridgeContractInterface.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeContractInterfaceInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeContractInterface.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeContractInterfaceInitialized)
				if err := _BridgeContractInterface.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) ParseInitialized(log types.Log) (*BridgeContractInterfaceInitialized, error) {
	event := new(BridgeContractInterfaceInitialized)
	if err := _BridgeContractInterface.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeContractInterfaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeContractInterface contract.
type BridgeContractInterfaceOwnershipTransferredIterator struct {
	Event *BridgeContractInterfaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeContractInterfaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeContractInterfaceOwnershipTransferred)
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
		it.Event = new(BridgeContractInterfaceOwnershipTransferred)
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
func (it *BridgeContractInterfaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractInterfaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeContractInterfaceOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeContractInterface contract.
type BridgeContractInterfaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeContractInterfaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceOwnershipTransferredIterator{contract: _BridgeContractInterface.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeContractInterfaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeContractInterfaceOwnershipTransferred)
				if err := _BridgeContractInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeContractInterfaceOwnershipTransferred, error) {
	event := new(BridgeContractInterfaceOwnershipTransferred)
	if err := _BridgeContractInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeContractInterfaceStateAppliedIterator is returned from FilterStateApplied and is used to iterate over the raw logs and unpacked data for StateApplied events raised by the BridgeContractInterface contract.
type BridgeContractInterfaceStateAppliedIterator struct {
	Event *BridgeContractInterfaceStateApplied // Event containing the contract specifics and raw log

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
func (it *BridgeContractInterfaceStateAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeContractInterfaceStateApplied)
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
		it.Event = new(BridgeContractInterfaceStateApplied)
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
func (it *BridgeContractInterfaceStateAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractInterfaceStateAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeContractInterfaceStateApplied represents a StateApplied event raised by the BridgeContractInterface contract.
type BridgeContractInterfaceStateApplied struct {
	State     common.Hash
	Receiver  common.Address
	Connector common.Address
	Refund    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateApplied is a free log retrieval operation binding the contract event 0xb3f87aed1075b6cfb4baab631c807ace0ef9b56b0a7cda905c99bce86fea6768.
//
// Solidity: event StateApplied(bytes indexed state, address indexed receiver, address indexed connector, uint256 refund)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) FilterStateApplied(opts *bind.FilterOpts, state [][]byte, receiver []common.Address, connector []common.Address) (*BridgeContractInterfaceStateAppliedIterator, error) {

	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var connectorRule []interface{}
	for _, connectorItem := range connector {
		connectorRule = append(connectorRule, connectorItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.FilterLogs(opts, "StateApplied", stateRule, receiverRule, connectorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceStateAppliedIterator{contract: _BridgeContractInterface.contract, event: "StateApplied", logs: logs, sub: sub}, nil
}

// WatchStateApplied is a free log subscription operation binding the contract event 0xb3f87aed1075b6cfb4baab631c807ace0ef9b56b0a7cda905c99bce86fea6768.
//
// Solidity: event StateApplied(bytes indexed state, address indexed receiver, address indexed connector, uint256 refund)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) WatchStateApplied(opts *bind.WatchOpts, sink chan<- *BridgeContractInterfaceStateApplied, state [][]byte, receiver []common.Address, connector []common.Address) (event.Subscription, error) {

	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var connectorRule []interface{}
	for _, connectorItem := range connector {
		connectorRule = append(connectorRule, connectorItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.WatchLogs(opts, "StateApplied", stateRule, receiverRule, connectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeContractInterfaceStateApplied)
				if err := _BridgeContractInterface.contract.UnpackLog(event, "StateApplied", log); err != nil {
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

// ParseStateApplied is a log parse operation binding the contract event 0xb3f87aed1075b6cfb4baab631c807ace0ef9b56b0a7cda905c99bce86fea6768.
//
// Solidity: event StateApplied(bytes indexed state, address indexed receiver, address indexed connector, uint256 refund)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) ParseStateApplied(log types.Log) (*BridgeContractInterfaceStateApplied, error) {
	event := new(BridgeContractInterfaceStateApplied)
	if err := _BridgeContractInterface.contract.UnpackLog(event, "StateApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeContractInterfaceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BridgeContractInterface contract.
type BridgeContractInterfaceUpgradedIterator struct {
	Event *BridgeContractInterfaceUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeContractInterfaceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeContractInterfaceUpgraded)
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
		it.Event = new(BridgeContractInterfaceUpgraded)
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
func (it *BridgeContractInterfaceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeContractInterfaceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeContractInterfaceUpgraded represents a Upgraded event raised by the BridgeContractInterface contract.
type BridgeContractInterfaceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeContractInterfaceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeContractInterfaceUpgradedIterator{contract: _BridgeContractInterface.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeContractInterfaceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeContractInterface.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeContractInterfaceUpgraded)
				if err := _BridgeContractInterface.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeContractInterface *BridgeContractInterfaceFilterer) ParseUpgraded(log types.Log) (*BridgeContractInterfaceUpgraded, error) {
	event := new(BridgeContractInterfaceUpgraded)
	if err := _BridgeContractInterface.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BridgeBase

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

// BridgeBaseMetaData contains all meta data concerning the BridgeBase contract.
var BridgeBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"appliedEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"applyState\",\"inputs\":[{\"name\":\"state_with_proof\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.StateWithProof\",\"components\":[{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.BridgeState\",\"components\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"states\",\"type\":\"tuple[]\",\"internalType\":\"structSharedStructs.StateWithAddress[]\",\"components\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"state_hashes\",\"type\":\"tuple[]\",\"internalType\":\"structSharedStructs.ContractStateHash[]\",\"components\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stateHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"connectors\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBridgeConnector\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizationInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeEpoch\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizedEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBridgeRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateWithProof\",\"inputs\":[],\"outputs\":[{\"name\":\"ret\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.StateWithProof\",\"components\":[{\"name\":\"state\",\"type\":\"tuple\",\"internalType\":\"structSharedStructs.BridgeState\",\"components\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"states\",\"type\":\"tuple[]\",\"internalType\":\"structSharedStructs.StateWithAddress[]\",\"components\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"state_hashes\",\"type\":\"tuple[]\",\"internalType\":\"structSharedStructs.ContractStateHash[]\",\"components\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stateHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastFinalizedBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lightClient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBridgeLightClient\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localAddress\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerContract\",\"inputs\":[{\"name\":\"connector\",\"type\":\"address\",\"internalType\":\"contractIBridgeConnector\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registeredTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalizationInterval\",\"inputs\":[{\"name\":\"_finalizationInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenAddresses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ConnectorRegistered\",\"inputs\":[{\"name\":\"connector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token_source\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token_destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Finalized\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ConnectorAlreadyRegistered\",\"inputs\":[{\"name\":\"connector\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoStateToFinalize\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllStatesApplied\",\"inputs\":[{\"name\":\"processed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughBlocksPassed\",\"inputs\":[{\"name\":\"lastFinalizedBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSuccessiveEpochs\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nextEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StateNotMatchingBridgeRoot\",\"inputs\":[{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressCannotBeRegistered\",\"inputs\":[]}]",
}

// BridgeBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeBaseMetaData.ABI instead.
var BridgeBaseABI = BridgeBaseMetaData.ABI

// BridgeBase is an auto generated Go binding around an Ethereum contract.
type BridgeBase struct {
	BridgeBaseCaller     // Read-only binding to the contract
	BridgeBaseTransactor // Write-only binding to the contract
	BridgeBaseFilterer   // Log filterer for contract events
}

// BridgeBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeBaseSession struct {
	Contract     *BridgeBase       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeBaseCallerSession struct {
	Contract *BridgeBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeBaseTransactorSession struct {
	Contract     *BridgeBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeBaseRaw struct {
	Contract *BridgeBase // Generic contract binding to access the raw methods on
}

// BridgeBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeBaseCallerRaw struct {
	Contract *BridgeBaseCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeBaseTransactorRaw struct {
	Contract *BridgeBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeBase creates a new instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBase(address common.Address, backend bind.ContractBackend) (*BridgeBase, error) {
	contract, err := bindBridgeBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeBase{BridgeBaseCaller: BridgeBaseCaller{contract: contract}, BridgeBaseTransactor: BridgeBaseTransactor{contract: contract}, BridgeBaseFilterer: BridgeBaseFilterer{contract: contract}}, nil
}

// NewBridgeBaseCaller creates a new read-only instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseCaller(address common.Address, caller bind.ContractCaller) (*BridgeBaseCaller, error) {
	contract, err := bindBridgeBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseCaller{contract: contract}, nil
}

// NewBridgeBaseTransactor creates a new write-only instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeBaseTransactor, error) {
	contract, err := bindBridgeBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseTransactor{contract: contract}, nil
}

// NewBridgeBaseFilterer creates a new log filterer instance of BridgeBase, bound to a specific deployed contract.
func NewBridgeBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeBaseFilterer, error) {
	contract, err := bindBridgeBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseFilterer{contract: contract}, nil
}

// bindBridgeBase binds a generic wrapper to an already deployed contract.
func bindBridgeBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBase *BridgeBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBase.Contract.BridgeBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBase *BridgeBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.Contract.BridgeBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBase *BridgeBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBase.Contract.BridgeBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeBase *BridgeBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeBase *BridgeBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeBase *BridgeBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeBase.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeBase *BridgeBaseCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeBase *BridgeBaseSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeBase.Contract.UPGRADEINTERFACEVERSION(&_BridgeBase.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BridgeBase *BridgeBaseCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BridgeBase.Contract.UPGRADEINTERFACEVERSION(&_BridgeBase.CallOpts)
}

// AppliedEpoch is a free data retrieval call binding the contract method 0x35add093.
//
// Solidity: function appliedEpoch() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) AppliedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "appliedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppliedEpoch is a free data retrieval call binding the contract method 0x35add093.
//
// Solidity: function appliedEpoch() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) AppliedEpoch() (*big.Int, error) {
	return _BridgeBase.Contract.AppliedEpoch(&_BridgeBase.CallOpts)
}

// AppliedEpoch is a free data retrieval call binding the contract method 0x35add093.
//
// Solidity: function appliedEpoch() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) AppliedEpoch() (*big.Int, error) {
	return _BridgeBase.Contract.AppliedEpoch(&_BridgeBase.CallOpts)
}

// BridgeRoot is a free data retrieval call binding the contract method 0x177a0c2c.
//
// Solidity: function bridgeRoot() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) BridgeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "bridgeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BridgeRoot is a free data retrieval call binding the contract method 0x177a0c2c.
//
// Solidity: function bridgeRoot() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) BridgeRoot() ([32]byte, error) {
	return _BridgeBase.Contract.BridgeRoot(&_BridgeBase.CallOpts)
}

// BridgeRoot is a free data retrieval call binding the contract method 0x177a0c2c.
//
// Solidity: function bridgeRoot() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) BridgeRoot() ([32]byte, error) {
	return _BridgeBase.Contract.BridgeRoot(&_BridgeBase.CallOpts)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address ) view returns(address)
func (_BridgeBase *BridgeBaseCaller) Connectors(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "connectors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address ) view returns(address)
func (_BridgeBase *BridgeBaseSession) Connectors(arg0 common.Address) (common.Address, error) {
	return _BridgeBase.Contract.Connectors(&_BridgeBase.CallOpts, arg0)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address ) view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) Connectors(arg0 common.Address) (common.Address, error) {
	return _BridgeBase.Contract.Connectors(&_BridgeBase.CallOpts, arg0)
}

// FinalizationInterval is a free data retrieval call binding the contract method 0xbb04c6fc.
//
// Solidity: function finalizationInterval() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) FinalizationInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "finalizationInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizationInterval is a free data retrieval call binding the contract method 0xbb04c6fc.
//
// Solidity: function finalizationInterval() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) FinalizationInterval() (*big.Int, error) {
	return _BridgeBase.Contract.FinalizationInterval(&_BridgeBase.CallOpts)
}

// FinalizationInterval is a free data retrieval call binding the contract method 0xbb04c6fc.
//
// Solidity: function finalizationInterval() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) FinalizationInterval() (*big.Int, error) {
	return _BridgeBase.Contract.FinalizationInterval(&_BridgeBase.CallOpts)
}

// FinalizedEpoch is a free data retrieval call binding the contract method 0x6bfa7398.
//
// Solidity: function finalizedEpoch() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) FinalizedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "finalizedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizedEpoch is a free data retrieval call binding the contract method 0x6bfa7398.
//
// Solidity: function finalizedEpoch() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) FinalizedEpoch() (*big.Int, error) {
	return _BridgeBase.Contract.FinalizedEpoch(&_BridgeBase.CallOpts)
}

// FinalizedEpoch is a free data retrieval call binding the contract method 0x6bfa7398.
//
// Solidity: function finalizedEpoch() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) FinalizedEpoch() (*big.Int, error) {
	return _BridgeBase.Contract.FinalizedEpoch(&_BridgeBase.CallOpts)
}

// GetBridgeRoot is a free data retrieval call binding the contract method 0x695a253f.
//
// Solidity: function getBridgeRoot() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) GetBridgeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "getBridgeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBridgeRoot is a free data retrieval call binding the contract method 0x695a253f.
//
// Solidity: function getBridgeRoot() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) GetBridgeRoot() ([32]byte, error) {
	return _BridgeBase.Contract.GetBridgeRoot(&_BridgeBase.CallOpts)
}

// GetBridgeRoot is a free data retrieval call binding the contract method 0x695a253f.
//
// Solidity: function getBridgeRoot() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) GetBridgeRoot() ([32]byte, error) {
	return _BridgeBase.Contract.GetBridgeRoot(&_BridgeBase.CallOpts)
}

// GetStateWithProof is a free data retrieval call binding the contract method 0xfe65d463.
//
// Solidity: function getStateWithProof() view returns(((uint256,(address,bytes)[]),(address,bytes32)[]) ret)
func (_BridgeBase *BridgeBaseCaller) GetStateWithProof(opts *bind.CallOpts) (SharedStructsStateWithProof, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "getStateWithProof")

	if err != nil {
		return *new(SharedStructsStateWithProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsStateWithProof)).(*SharedStructsStateWithProof)

	return out0, err

}

// GetStateWithProof is a free data retrieval call binding the contract method 0xfe65d463.
//
// Solidity: function getStateWithProof() view returns(((uint256,(address,bytes)[]),(address,bytes32)[]) ret)
func (_BridgeBase *BridgeBaseSession) GetStateWithProof() (SharedStructsStateWithProof, error) {
	return _BridgeBase.Contract.GetStateWithProof(&_BridgeBase.CallOpts)
}

// GetStateWithProof is a free data retrieval call binding the contract method 0xfe65d463.
//
// Solidity: function getStateWithProof() view returns(((uint256,(address,bytes)[]),(address,bytes32)[]) ret)
func (_BridgeBase *BridgeBaseCallerSession) GetStateWithProof() (SharedStructsStateWithProof, error) {
	return _BridgeBase.Contract.GetStateWithProof(&_BridgeBase.CallOpts)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0xae1da0b5.
//
// Solidity: function lastFinalizedBlock() view returns(uint256)
func (_BridgeBase *BridgeBaseCaller) LastFinalizedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "lastFinalizedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0xae1da0b5.
//
// Solidity: function lastFinalizedBlock() view returns(uint256)
func (_BridgeBase *BridgeBaseSession) LastFinalizedBlock() (*big.Int, error) {
	return _BridgeBase.Contract.LastFinalizedBlock(&_BridgeBase.CallOpts)
}

// LastFinalizedBlock is a free data retrieval call binding the contract method 0xae1da0b5.
//
// Solidity: function lastFinalizedBlock() view returns(uint256)
func (_BridgeBase *BridgeBaseCallerSession) LastFinalizedBlock() (*big.Int, error) {
	return _BridgeBase.Contract.LastFinalizedBlock(&_BridgeBase.CallOpts)
}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_BridgeBase *BridgeBaseCaller) LightClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "lightClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_BridgeBase *BridgeBaseSession) LightClient() (common.Address, error) {
	return _BridgeBase.Contract.LightClient(&_BridgeBase.CallOpts)
}

// LightClient is a free data retrieval call binding the contract method 0xb5700e68.
//
// Solidity: function lightClient() view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) LightClient() (common.Address, error) {
	return _BridgeBase.Contract.LightClient(&_BridgeBase.CallOpts)
}

// LocalAddress is a free data retrieval call binding the contract method 0x76081bd5.
//
// Solidity: function localAddress(address ) view returns(address)
func (_BridgeBase *BridgeBaseCaller) LocalAddress(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "localAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalAddress is a free data retrieval call binding the contract method 0x76081bd5.
//
// Solidity: function localAddress(address ) view returns(address)
func (_BridgeBase *BridgeBaseSession) LocalAddress(arg0 common.Address) (common.Address, error) {
	return _BridgeBase.Contract.LocalAddress(&_BridgeBase.CallOpts, arg0)
}

// LocalAddress is a free data retrieval call binding the contract method 0x76081bd5.
//
// Solidity: function localAddress(address ) view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) LocalAddress(arg0 common.Address) (common.Address, error) {
	return _BridgeBase.Contract.LocalAddress(&_BridgeBase.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBase *BridgeBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBase *BridgeBaseSession) Owner() (common.Address, error) {
	return _BridgeBase.Contract.Owner(&_BridgeBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) Owner() (common.Address, error) {
	return _BridgeBase.Contract.Owner(&_BridgeBase.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeBase *BridgeBaseCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeBase *BridgeBaseSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeBase.Contract.ProxiableUUID(&_BridgeBase.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BridgeBase *BridgeBaseCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BridgeBase.Contract.ProxiableUUID(&_BridgeBase.CallOpts)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x45466616.
//
// Solidity: function registeredTokens() view returns(address[])
func (_BridgeBase *BridgeBaseCaller) RegisteredTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "registeredTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RegisteredTokens is a free data retrieval call binding the contract method 0x45466616.
//
// Solidity: function registeredTokens() view returns(address[])
func (_BridgeBase *BridgeBaseSession) RegisteredTokens() ([]common.Address, error) {
	return _BridgeBase.Contract.RegisteredTokens(&_BridgeBase.CallOpts)
}

// RegisteredTokens is a free data retrieval call binding the contract method 0x45466616.
//
// Solidity: function registeredTokens() view returns(address[])
func (_BridgeBase *BridgeBaseCallerSession) RegisteredTokens() ([]common.Address, error) {
	return _BridgeBase.Contract.RegisteredTokens(&_BridgeBase.CallOpts)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_BridgeBase *BridgeBaseCaller) TokenAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeBase.contract.Call(opts, &out, "tokenAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_BridgeBase *BridgeBaseSession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _BridgeBase.Contract.TokenAddresses(&_BridgeBase.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0xe5df8b84.
//
// Solidity: function tokenAddresses(uint256 ) view returns(address)
func (_BridgeBase *BridgeBaseCallerSession) TokenAddresses(arg0 *big.Int) (common.Address, error) {
	return _BridgeBase.Contract.TokenAddresses(&_BridgeBase.CallOpts, arg0)
}

// ApplyState is a paid mutator transaction binding the contract method 0x6cd50a67.
//
// Solidity: function applyState(((uint256,(address,bytes)[]),(address,bytes32)[]) state_with_proof) returns()
func (_BridgeBase *BridgeBaseTransactor) ApplyState(opts *bind.TransactOpts, state_with_proof SharedStructsStateWithProof) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "applyState", state_with_proof)
}

// ApplyState is a paid mutator transaction binding the contract method 0x6cd50a67.
//
// Solidity: function applyState(((uint256,(address,bytes)[]),(address,bytes32)[]) state_with_proof) returns()
func (_BridgeBase *BridgeBaseSession) ApplyState(state_with_proof SharedStructsStateWithProof) (*types.Transaction, error) {
	return _BridgeBase.Contract.ApplyState(&_BridgeBase.TransactOpts, state_with_proof)
}

// ApplyState is a paid mutator transaction binding the contract method 0x6cd50a67.
//
// Solidity: function applyState(((uint256,(address,bytes)[]),(address,bytes32)[]) state_with_proof) returns()
func (_BridgeBase *BridgeBaseTransactorSession) ApplyState(state_with_proof SharedStructsStateWithProof) (*types.Transaction, error) {
	return _BridgeBase.Contract.ApplyState(&_BridgeBase.TransactOpts, state_with_proof)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x82ae9ef7.
//
// Solidity: function finalizeEpoch() returns()
func (_BridgeBase *BridgeBaseTransactor) FinalizeEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "finalizeEpoch")
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x82ae9ef7.
//
// Solidity: function finalizeEpoch() returns()
func (_BridgeBase *BridgeBaseSession) FinalizeEpoch() (*types.Transaction, error) {
	return _BridgeBase.Contract.FinalizeEpoch(&_BridgeBase.TransactOpts)
}

// FinalizeEpoch is a paid mutator transaction binding the contract method 0x82ae9ef7.
//
// Solidity: function finalizeEpoch() returns()
func (_BridgeBase *BridgeBaseTransactorSession) FinalizeEpoch() (*types.Transaction, error) {
	return _BridgeBase.Contract.FinalizeEpoch(&_BridgeBase.TransactOpts)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x22a5dde4.
//
// Solidity: function registerContract(address connector) returns()
func (_BridgeBase *BridgeBaseTransactor) RegisterContract(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "registerContract", connector)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x22a5dde4.
//
// Solidity: function registerContract(address connector) returns()
func (_BridgeBase *BridgeBaseSession) RegisterContract(connector common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RegisterContract(&_BridgeBase.TransactOpts, connector)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x22a5dde4.
//
// Solidity: function registerContract(address connector) returns()
func (_BridgeBase *BridgeBaseTransactorSession) RegisterContract(connector common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.RegisterContract(&_BridgeBase.TransactOpts, connector)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBase *BridgeBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBase *BridgeBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBase.Contract.RenounceOwnership(&_BridgeBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeBase *BridgeBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeBase.Contract.RenounceOwnership(&_BridgeBase.TransactOpts)
}

// SetFinalizationInterval is a paid mutator transaction binding the contract method 0x8f084f0f.
//
// Solidity: function setFinalizationInterval(uint256 _finalizationInterval) returns()
func (_BridgeBase *BridgeBaseTransactor) SetFinalizationInterval(opts *bind.TransactOpts, _finalizationInterval *big.Int) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "setFinalizationInterval", _finalizationInterval)
}

// SetFinalizationInterval is a paid mutator transaction binding the contract method 0x8f084f0f.
//
// Solidity: function setFinalizationInterval(uint256 _finalizationInterval) returns()
func (_BridgeBase *BridgeBaseSession) SetFinalizationInterval(_finalizationInterval *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetFinalizationInterval(&_BridgeBase.TransactOpts, _finalizationInterval)
}

// SetFinalizationInterval is a paid mutator transaction binding the contract method 0x8f084f0f.
//
// Solidity: function setFinalizationInterval(uint256 _finalizationInterval) returns()
func (_BridgeBase *BridgeBaseTransactorSession) SetFinalizationInterval(_finalizationInterval *big.Int) (*types.Transaction, error) {
	return _BridgeBase.Contract.SetFinalizationInterval(&_BridgeBase.TransactOpts, _finalizationInterval)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBase *BridgeBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBase *BridgeBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.TransferOwnership(&_BridgeBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeBase *BridgeBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeBase.Contract.TransferOwnership(&_BridgeBase.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeBase *BridgeBaseTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeBase.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeBase *BridgeBaseSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeBase.Contract.UpgradeToAndCall(&_BridgeBase.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BridgeBase *BridgeBaseTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BridgeBase.Contract.UpgradeToAndCall(&_BridgeBase.TransactOpts, newImplementation, data)
}

// BridgeBaseConnectorRegisteredIterator is returned from FilterConnectorRegistered and is used to iterate over the raw logs and unpacked data for ConnectorRegistered events raised by the BridgeBase contract.
type BridgeBaseConnectorRegisteredIterator struct {
	Event *BridgeBaseConnectorRegistered // Event containing the contract specifics and raw log

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
func (it *BridgeBaseConnectorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseConnectorRegistered)
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
		it.Event = new(BridgeBaseConnectorRegistered)
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
func (it *BridgeBaseConnectorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseConnectorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseConnectorRegistered represents a ConnectorRegistered event raised by the BridgeBase contract.
type BridgeBaseConnectorRegistered struct {
	Connector        common.Address
	TokenSource      common.Address
	TokenDestination common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConnectorRegistered is a free log retrieval operation binding the contract event 0x2d481d4ec464c9eef253e96f6662dd4420b22159ddc54ab42d443e0c8604f34e.
//
// Solidity: event ConnectorRegistered(address indexed connector, address indexed token_source, address indexed token_destination)
func (_BridgeBase *BridgeBaseFilterer) FilterConnectorRegistered(opts *bind.FilterOpts, connector []common.Address, token_source []common.Address, token_destination []common.Address) (*BridgeBaseConnectorRegisteredIterator, error) {

	var connectorRule []interface{}
	for _, connectorItem := range connector {
		connectorRule = append(connectorRule, connectorItem)
	}
	var token_sourceRule []interface{}
	for _, token_sourceItem := range token_source {
		token_sourceRule = append(token_sourceRule, token_sourceItem)
	}
	var token_destinationRule []interface{}
	for _, token_destinationItem := range token_destination {
		token_destinationRule = append(token_destinationRule, token_destinationItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "ConnectorRegistered", connectorRule, token_sourceRule, token_destinationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseConnectorRegisteredIterator{contract: _BridgeBase.contract, event: "ConnectorRegistered", logs: logs, sub: sub}, nil
}

// WatchConnectorRegistered is a free log subscription operation binding the contract event 0x2d481d4ec464c9eef253e96f6662dd4420b22159ddc54ab42d443e0c8604f34e.
//
// Solidity: event ConnectorRegistered(address indexed connector, address indexed token_source, address indexed token_destination)
func (_BridgeBase *BridgeBaseFilterer) WatchConnectorRegistered(opts *bind.WatchOpts, sink chan<- *BridgeBaseConnectorRegistered, connector []common.Address, token_source []common.Address, token_destination []common.Address) (event.Subscription, error) {

	var connectorRule []interface{}
	for _, connectorItem := range connector {
		connectorRule = append(connectorRule, connectorItem)
	}
	var token_sourceRule []interface{}
	for _, token_sourceItem := range token_source {
		token_sourceRule = append(token_sourceRule, token_sourceItem)
	}
	var token_destinationRule []interface{}
	for _, token_destinationItem := range token_destination {
		token_destinationRule = append(token_destinationRule, token_destinationItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "ConnectorRegistered", connectorRule, token_sourceRule, token_destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseConnectorRegistered)
				if err := _BridgeBase.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
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

// ParseConnectorRegistered is a log parse operation binding the contract event 0x2d481d4ec464c9eef253e96f6662dd4420b22159ddc54ab42d443e0c8604f34e.
//
// Solidity: event ConnectorRegistered(address indexed connector, address indexed token_source, address indexed token_destination)
func (_BridgeBase *BridgeBaseFilterer) ParseConnectorRegistered(log types.Log) (*BridgeBaseConnectorRegistered, error) {
	event := new(BridgeBaseConnectorRegistered)
	if err := _BridgeBase.contract.UnpackLog(event, "ConnectorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the BridgeBase contract.
type BridgeBaseFinalizedIterator struct {
	Event *BridgeBaseFinalized // Event containing the contract specifics and raw log

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
func (it *BridgeBaseFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseFinalized)
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
		it.Event = new(BridgeBaseFinalized)
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
func (it *BridgeBaseFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseFinalized represents a Finalized event raised by the BridgeBase contract.
type BridgeBaseFinalized struct {
	Epoch      *big.Int
	BridgeRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0xa05a0e9561eff1f01a29e7a680d5957bb7312e5766a8da1f494b6d6ac18031f4.
//
// Solidity: event Finalized(uint256 indexed epoch, bytes32 bridgeRoot)
func (_BridgeBase *BridgeBaseFilterer) FilterFinalized(opts *bind.FilterOpts, epoch []*big.Int) (*BridgeBaseFinalizedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Finalized", epochRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseFinalizedIterator{contract: _BridgeBase.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0xa05a0e9561eff1f01a29e7a680d5957bb7312e5766a8da1f494b6d6ac18031f4.
//
// Solidity: event Finalized(uint256 indexed epoch, bytes32 bridgeRoot)
func (_BridgeBase *BridgeBaseFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *BridgeBaseFinalized, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Finalized", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseFinalized)
				if err := _BridgeBase.contract.UnpackLog(event, "Finalized", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseFinalized(log types.Log) (*BridgeBaseFinalized, error) {
	event := new(BridgeBaseFinalized)
	if err := _BridgeBase.contract.UnpackLog(event, "Finalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeBase contract.
type BridgeBaseInitializedIterator struct {
	Event *BridgeBaseInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseInitialized)
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
		it.Event = new(BridgeBaseInitialized)
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
func (it *BridgeBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseInitialized represents a Initialized event raised by the BridgeBase contract.
type BridgeBaseInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeBase *BridgeBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeBaseInitializedIterator, error) {

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeBaseInitializedIterator{contract: _BridgeBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BridgeBase *BridgeBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseInitialized)
				if err := _BridgeBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseInitialized(log types.Log) (*BridgeBaseInitialized, error) {
	event := new(BridgeBaseInitialized)
	if err := _BridgeBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeBase contract.
type BridgeBaseOwnershipTransferredIterator struct {
	Event *BridgeBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseOwnershipTransferred)
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
		it.Event = new(BridgeBaseOwnershipTransferred)
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
func (it *BridgeBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeBase contract.
type BridgeBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBase *BridgeBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseOwnershipTransferredIterator{contract: _BridgeBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeBase *BridgeBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseOwnershipTransferred)
				if err := _BridgeBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeBaseOwnershipTransferred, error) {
	event := new(BridgeBaseOwnershipTransferred)
	if err := _BridgeBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBaseUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BridgeBase contract.
type BridgeBaseUpgradedIterator struct {
	Event *BridgeBaseUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeBaseUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBaseUpgraded)
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
		it.Event = new(BridgeBaseUpgraded)
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
func (it *BridgeBaseUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBaseUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBaseUpgraded represents a Upgraded event raised by the BridgeBase contract.
type BridgeBaseUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeBase *BridgeBaseFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeBaseUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeBase.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBaseUpgradedIterator{contract: _BridgeBase.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BridgeBase *BridgeBaseFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeBaseUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BridgeBase.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBaseUpgraded)
				if err := _BridgeBase.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BridgeBase *BridgeBaseFilterer) ParseUpgraded(log types.Log) (*BridgeBaseUpgraded, error) {
	event := new(BridgeBaseUpgraded)
	if err := _BridgeBase.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

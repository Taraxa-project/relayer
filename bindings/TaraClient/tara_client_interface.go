// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tara_client_interface

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

// CompactSignature is an auto generated low-level Go binding around an user-defined struct.
type CompactSignature struct {
	R  [32]byte
	Vs [32]byte
}

// PillarBlockFinalizationData is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockFinalizationData struct {
	Period     *big.Int
	StateRoot  [32]byte
	PrevHash   [32]byte
	BridgeRoot [32]byte
	Epoch      *big.Int
}

// PillarBlockFinalizedBlock is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockFinalizedBlock struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}

// PillarBlockVoteCountChange is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockVoteCountChange struct {
	Validator common.Address
	Change    int32
}

// PillarBlockWithChanges is an auto generated low-level Go binding around an user-defined struct.
type PillarBlockWithChanges struct {
	Block            PillarBlockFinalizationData
	ValidatorChanges []PillarBlockVoteCountChange
}

// TaraClientInterfaceMetaData contains all meta data concerning the TaraClientInterface contract.
var TaraClientInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeBlocks\",\"inputs\":[{\"name\":\"blocks\",\"type\":\"tuple[]\",\"internalType\":\"structPillarBlock.WithChanges[]\",\"components\":[{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"validatorChanges\",\"type\":\"tuple[]\",\"internalType\":\"structPillarBlock.VoteCountChange[]\",\"components\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"change\",\"type\":\"int32\",\"internalType\":\"int32\"}]}]},{\"name\":\"lastBlockSigs\",\"type\":\"tuple[]\",\"internalType\":\"structCompactSignature[]\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalized\",\"inputs\":[],\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"finalizedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizedBridgeRoots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizedBlock\",\"components\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"finalizedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizedBridgeRoot\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignaturesWeight\",\"inputs\":[{\"name\":\"h\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structCompactSignature[]\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_pillarBlockInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pillarBlockInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"threshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorVoteCounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlockFinalized\",\"inputs\":[{\"name\":\"finalized\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPillarBlock.FinalizedBlock\",\"components\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"finalizedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdChanged\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorWeightChanged\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"HashesNotMatching\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidBlockInterval\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// TaraClientInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use TaraClientInterfaceMetaData.ABI instead.
var TaraClientInterfaceABI = TaraClientInterfaceMetaData.ABI

// TaraClientInterface is an auto generated Go binding around an Ethereum contract.
type TaraClientInterface struct {
	TaraClientInterfaceCaller     // Read-only binding to the contract
	TaraClientInterfaceTransactor // Write-only binding to the contract
	TaraClientInterfaceFilterer   // Log filterer for contract events
}

// TaraClientInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaraClientInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaraClientInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaraClientInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaraClientInterfaceSession struct {
	Contract     *TaraClientInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TaraClientInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaraClientInterfaceCallerSession struct {
	Contract *TaraClientInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TaraClientInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaraClientInterfaceTransactorSession struct {
	Contract     *TaraClientInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TaraClientInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaraClientInterfaceRaw struct {
	Contract *TaraClientInterface // Generic contract binding to access the raw methods on
}

// TaraClientInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaraClientInterfaceCallerRaw struct {
	Contract *TaraClientInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// TaraClientInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaraClientInterfaceTransactorRaw struct {
	Contract *TaraClientInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaraClientInterface creates a new instance of TaraClientInterface, bound to a specific deployed contract.
func NewTaraClientInterface(address common.Address, backend bind.ContractBackend) (*TaraClientInterface, error) {
	contract, err := bindTaraClientInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaraClientInterface{TaraClientInterfaceCaller: TaraClientInterfaceCaller{contract: contract}, TaraClientInterfaceTransactor: TaraClientInterfaceTransactor{contract: contract}, TaraClientInterfaceFilterer: TaraClientInterfaceFilterer{contract: contract}}, nil
}

// NewTaraClientInterfaceCaller creates a new read-only instance of TaraClientInterface, bound to a specific deployed contract.
func NewTaraClientInterfaceCaller(address common.Address, caller bind.ContractCaller) (*TaraClientInterfaceCaller, error) {
	contract, err := bindTaraClientInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceCaller{contract: contract}, nil
}

// NewTaraClientInterfaceTransactor creates a new write-only instance of TaraClientInterface, bound to a specific deployed contract.
func NewTaraClientInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TaraClientInterfaceTransactor, error) {
	contract, err := bindTaraClientInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceTransactor{contract: contract}, nil
}

// NewTaraClientInterfaceFilterer creates a new log filterer instance of TaraClientInterface, bound to a specific deployed contract.
func NewTaraClientInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TaraClientInterfaceFilterer, error) {
	contract, err := bindTaraClientInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceFilterer{contract: contract}, nil
}

// bindTaraClientInterface binds a generic wrapper to an already deployed contract.
func bindTaraClientInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaraClientInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClientInterface *TaraClientInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClientInterface.Contract.TaraClientInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClientInterface *TaraClientInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.TaraClientInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClientInterface *TaraClientInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.TaraClientInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClientInterface *TaraClientInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClientInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClientInterface *TaraClientInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClientInterface *TaraClientInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.contract.Transact(opts, method, params...)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32,uint256) block, uint256 finalizedAt)
func (_TaraClientInterface *TaraClientInterfaceCaller) Finalized(opts *bind.CallOpts) (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "finalized")

	outstruct := new(struct {
		BlockHash   [32]byte
		Block       PillarBlockFinalizationData
		FinalizedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Block = *abi.ConvertType(out[1], new(PillarBlockFinalizationData)).(*PillarBlockFinalizationData)
	outstruct.FinalizedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32,uint256) block, uint256 finalizedAt)
func (_TaraClientInterface *TaraClientInterfaceSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClientInterface.Contract.Finalized(&_TaraClientInterface.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32,uint256) block, uint256 finalizedAt)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClientInterface.Contract.Finalized(&_TaraClientInterface.CallOpts)
}

// FinalizedBridgeRoots is a free data retrieval call binding the contract method 0x93ea1203.
//
// Solidity: function finalizedBridgeRoots(uint256 ) view returns(bytes32)
func (_TaraClientInterface *TaraClientInterfaceCaller) FinalizedBridgeRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "finalizedBridgeRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedBridgeRoots is a free data retrieval call binding the contract method 0x93ea1203.
//
// Solidity: function finalizedBridgeRoots(uint256 ) view returns(bytes32)
func (_TaraClientInterface *TaraClientInterfaceSession) FinalizedBridgeRoots(arg0 *big.Int) ([32]byte, error) {
	return _TaraClientInterface.Contract.FinalizedBridgeRoots(&_TaraClientInterface.CallOpts, arg0)
}

// FinalizedBridgeRoots is a free data retrieval call binding the contract method 0x93ea1203.
//
// Solidity: function finalizedBridgeRoots(uint256 ) view returns(bytes32)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) FinalizedBridgeRoots(arg0 *big.Int) ([32]byte, error) {
	return _TaraClientInterface.Contract.FinalizedBridgeRoots(&_TaraClientInterface.CallOpts, arg0)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256))
func (_TaraClientInterface *TaraClientInterfaceCaller) GetFinalized(opts *bind.CallOpts) (PillarBlockFinalizedBlock, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "getFinalized")

	if err != nil {
		return *new(PillarBlockFinalizedBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(PillarBlockFinalizedBlock)).(*PillarBlockFinalizedBlock)

	return out0, err

}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256))
func (_TaraClientInterface *TaraClientInterfaceSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClientInterface.Contract.GetFinalized(&_TaraClientInterface.CallOpts)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256))
func (_TaraClientInterface *TaraClientInterfaceCallerSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClientInterface.Contract.GetFinalized(&_TaraClientInterface.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_TaraClientInterface *TaraClientInterfaceCaller) GetFinalizedBridgeRoot(opts *bind.CallOpts, epoch *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "getFinalizedBridgeRoot", epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_TaraClientInterface *TaraClientInterfaceSession) GetFinalizedBridgeRoot(epoch *big.Int) ([32]byte, error) {
	return _TaraClientInterface.Contract.GetFinalizedBridgeRoot(&_TaraClientInterface.CallOpts, epoch)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) GetFinalizedBridgeRoot(epoch *big.Int) ([32]byte, error) {
	return _TaraClientInterface.Contract.GetFinalizedBridgeRoot(&_TaraClientInterface.CallOpts, epoch)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientInterface *TaraClientInterfaceCaller) GetSignaturesWeight(opts *bind.CallOpts, h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "getSignaturesWeight", h, signatures)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientInterface *TaraClientInterfaceSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientInterface.Contract.GetSignaturesWeight(&_TaraClientInterface.CallOpts, h, signatures)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientInterface.Contract.GetSignaturesWeight(&_TaraClientInterface.CallOpts, h, signatures)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClientInterface *TaraClientInterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClientInterface *TaraClientInterfaceSession) Owner() (common.Address, error) {
	return _TaraClientInterface.Contract.Owner(&_TaraClientInterface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) Owner() (common.Address, error) {
	return _TaraClientInterface.Contract.Owner(&_TaraClientInterface.CallOpts)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCaller) PillarBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "pillarBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClientInterface.Contract.PillarBlockInterval(&_TaraClientInterface.CallOpts)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClientInterface.Contract.PillarBlockInterval(&_TaraClientInterface.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceSession) Threshold() (*big.Int, error) {
	return _TaraClientInterface.Contract.Threshold(&_TaraClientInterface.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) Threshold() (*big.Int, error) {
	return _TaraClientInterface.Contract.Threshold(&_TaraClientInterface.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceSession) TotalWeight() (*big.Int, error) {
	return _TaraClientInterface.Contract.TotalWeight(&_TaraClientInterface.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) TotalWeight() (*big.Int, error) {
	return _TaraClientInterface.Contract.TotalWeight(&_TaraClientInterface.CallOpts)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCaller) ValidatorVoteCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientInterface.contract.Call(opts, &out, "validatorVoteCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClientInterface.Contract.ValidatorVoteCounts(&_TaraClientInterface.CallOpts, arg0)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientInterface *TaraClientInterfaceCallerSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClientInterface.Contract.ValidatorVoteCounts(&_TaraClientInterface.CallOpts, arg0)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5d0d5734.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32,uint256),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactor) FinalizeBlocks(opts *bind.TransactOpts, blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientInterface.contract.Transact(opts, "finalizeBlocks", blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5d0d5734.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32,uint256),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientInterface *TaraClientInterfaceSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.FinalizeBlocks(&_TaraClientInterface.TransactOpts, blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5d0d5734.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32,uint256),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactorSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.FinalizeBlocks(&_TaraClientInterface.TransactOpts, blocks, lastBlockSigs)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _threshold, uint256 _pillarBlockInterval) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactor) Initialize(opts *bind.TransactOpts, _threshold *big.Int, _pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClientInterface.contract.Transact(opts, "initialize", _threshold, _pillarBlockInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _threshold, uint256 _pillarBlockInterval) returns()
func (_TaraClientInterface *TaraClientInterfaceSession) Initialize(_threshold *big.Int, _pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.Initialize(&_TaraClientInterface.TransactOpts, _threshold, _pillarBlockInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _threshold, uint256 _pillarBlockInterval) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactorSession) Initialize(_threshold *big.Int, _pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.Initialize(&_TaraClientInterface.TransactOpts, _threshold, _pillarBlockInterval)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClientInterface *TaraClientInterfaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientInterface.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClientInterface *TaraClientInterfaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaraClientInterface.Contract.RenounceOwnership(&_TaraClientInterface.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClientInterface *TaraClientInterfaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaraClientInterface.Contract.RenounceOwnership(&_TaraClientInterface.TransactOpts)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactor) SetThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientInterface.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientInterface *TaraClientInterfaceSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.SetThreshold(&_TaraClientInterface.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactorSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.SetThreshold(&_TaraClientInterface.TransactOpts, _threshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaraClientInterface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClientInterface *TaraClientInterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.TransferOwnership(&_TaraClientInterface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClientInterface *TaraClientInterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaraClientInterface.Contract.TransferOwnership(&_TaraClientInterface.TransactOpts, newOwner)
}

// TaraClientInterfaceBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the TaraClientInterface contract.
type TaraClientInterfaceBlockFinalizedIterator struct {
	Event *TaraClientInterfaceBlockFinalized // Event containing the contract specifics and raw log

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
func (it *TaraClientInterfaceBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientInterfaceBlockFinalized)
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
		it.Event = new(TaraClientInterfaceBlockFinalized)
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
func (it *TaraClientInterfaceBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientInterfaceBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientInterfaceBlockFinalized represents a BlockFinalized event raised by the TaraClientInterface contract.
type TaraClientInterfaceBlockFinalized struct {
	Finalized PillarBlockFinalizedBlock
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xc3af94e02b408059c0f32325f41a0d0384b419789b788f11d9d3ec59fb55cd45.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256) finalized)
func (_TaraClientInterface *TaraClientInterfaceFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*TaraClientInterfaceBlockFinalizedIterator, error) {

	logs, sub, err := _TaraClientInterface.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceBlockFinalizedIterator{contract: _TaraClientInterface.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xc3af94e02b408059c0f32325f41a0d0384b419789b788f11d9d3ec59fb55cd45.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256) finalized)
func (_TaraClientInterface *TaraClientInterfaceFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *TaraClientInterfaceBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _TaraClientInterface.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientInterfaceBlockFinalized)
				if err := _TaraClientInterface.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xc3af94e02b408059c0f32325f41a0d0384b419789b788f11d9d3ec59fb55cd45.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256) finalized)
func (_TaraClientInterface *TaraClientInterfaceFilterer) ParseBlockFinalized(log types.Log) (*TaraClientInterfaceBlockFinalized, error) {
	event := new(TaraClientInterfaceBlockFinalized)
	if err := _TaraClientInterface.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientInterfaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaraClientInterface contract.
type TaraClientInterfaceInitializedIterator struct {
	Event *TaraClientInterfaceInitialized // Event containing the contract specifics and raw log

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
func (it *TaraClientInterfaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientInterfaceInitialized)
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
		it.Event = new(TaraClientInterfaceInitialized)
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
func (it *TaraClientInterfaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientInterfaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientInterfaceInitialized represents a Initialized event raised by the TaraClientInterface contract.
type TaraClientInterfaceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaraClientInterface *TaraClientInterfaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaraClientInterfaceInitializedIterator, error) {

	logs, sub, err := _TaraClientInterface.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceInitializedIterator{contract: _TaraClientInterface.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaraClientInterface *TaraClientInterfaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaraClientInterfaceInitialized) (event.Subscription, error) {

	logs, sub, err := _TaraClientInterface.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientInterfaceInitialized)
				if err := _TaraClientInterface.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaraClientInterface *TaraClientInterfaceFilterer) ParseInitialized(log types.Log) (*TaraClientInterfaceInitialized, error) {
	event := new(TaraClientInterfaceInitialized)
	if err := _TaraClientInterface.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientInterfaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaraClientInterface contract.
type TaraClientInterfaceOwnershipTransferredIterator struct {
	Event *TaraClientInterfaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaraClientInterfaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientInterfaceOwnershipTransferred)
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
		it.Event = new(TaraClientInterfaceOwnershipTransferred)
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
func (it *TaraClientInterfaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientInterfaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientInterfaceOwnershipTransferred represents a OwnershipTransferred event raised by the TaraClientInterface contract.
type TaraClientInterfaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaraClientInterface *TaraClientInterfaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaraClientInterfaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaraClientInterface.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceOwnershipTransferredIterator{contract: _TaraClientInterface.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaraClientInterface *TaraClientInterfaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaraClientInterfaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaraClientInterface.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientInterfaceOwnershipTransferred)
				if err := _TaraClientInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaraClientInterface *TaraClientInterfaceFilterer) ParseOwnershipTransferred(log types.Log) (*TaraClientInterfaceOwnershipTransferred, error) {
	event := new(TaraClientInterfaceOwnershipTransferred)
	if err := _TaraClientInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientInterfaceThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the TaraClientInterface contract.
type TaraClientInterfaceThresholdChangedIterator struct {
	Event *TaraClientInterfaceThresholdChanged // Event containing the contract specifics and raw log

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
func (it *TaraClientInterfaceThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientInterfaceThresholdChanged)
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
		it.Event = new(TaraClientInterfaceThresholdChanged)
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
func (it *TaraClientInterfaceThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientInterfaceThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientInterfaceThresholdChanged represents a ThresholdChanged event raised by the TaraClientInterface contract.
type TaraClientInterfaceThresholdChanged struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_TaraClientInterface *TaraClientInterfaceFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*TaraClientInterfaceThresholdChangedIterator, error) {

	logs, sub, err := _TaraClientInterface.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceThresholdChangedIterator{contract: _TaraClientInterface.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_TaraClientInterface *TaraClientInterfaceFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *TaraClientInterfaceThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _TaraClientInterface.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientInterfaceThresholdChanged)
				if err := _TaraClientInterface.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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

// ParseThresholdChanged is a log parse operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_TaraClientInterface *TaraClientInterfaceFilterer) ParseThresholdChanged(log types.Log) (*TaraClientInterfaceThresholdChanged, error) {
	event := new(TaraClientInterfaceThresholdChanged)
	if err := _TaraClientInterface.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientInterfaceValidatorWeightChangedIterator is returned from FilterValidatorWeightChanged and is used to iterate over the raw logs and unpacked data for ValidatorWeightChanged events raised by the TaraClientInterface contract.
type TaraClientInterfaceValidatorWeightChangedIterator struct {
	Event *TaraClientInterfaceValidatorWeightChanged // Event containing the contract specifics and raw log

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
func (it *TaraClientInterfaceValidatorWeightChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientInterfaceValidatorWeightChanged)
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
		it.Event = new(TaraClientInterfaceValidatorWeightChanged)
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
func (it *TaraClientInterfaceValidatorWeightChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientInterfaceValidatorWeightChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientInterfaceValidatorWeightChanged represents a ValidatorWeightChanged event raised by the TaraClientInterface contract.
type TaraClientInterfaceValidatorWeightChanged struct {
	Validator common.Address
	Weight    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightChanged is a free log retrieval operation binding the contract event 0xe3490611629a41809c531d61c4adf71d3e682bfe292318e23f7227f1570d921b.
//
// Solidity: event ValidatorWeightChanged(address indexed validator, uint256 weight)
func (_TaraClientInterface *TaraClientInterfaceFilterer) FilterValidatorWeightChanged(opts *bind.FilterOpts, validator []common.Address) (*TaraClientInterfaceValidatorWeightChangedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraClientInterface.contract.FilterLogs(opts, "ValidatorWeightChanged", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraClientInterfaceValidatorWeightChangedIterator{contract: _TaraClientInterface.contract, event: "ValidatorWeightChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightChanged is a free log subscription operation binding the contract event 0xe3490611629a41809c531d61c4adf71d3e682bfe292318e23f7227f1570d921b.
//
// Solidity: event ValidatorWeightChanged(address indexed validator, uint256 weight)
func (_TaraClientInterface *TaraClientInterfaceFilterer) WatchValidatorWeightChanged(opts *bind.WatchOpts, sink chan<- *TaraClientInterfaceValidatorWeightChanged, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraClientInterface.contract.WatchLogs(opts, "ValidatorWeightChanged", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientInterfaceValidatorWeightChanged)
				if err := _TaraClientInterface.contract.UnpackLog(event, "ValidatorWeightChanged", log); err != nil {
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

// ParseValidatorWeightChanged is a log parse operation binding the contract event 0xe3490611629a41809c531d61c4adf71d3e682bfe292318e23f7227f1570d921b.
//
// Solidity: event ValidatorWeightChanged(address indexed validator, uint256 weight)
func (_TaraClientInterface *TaraClientInterfaceFilterer) ParseValidatorWeightChanged(log types.Log) (*TaraClientInterfaceValidatorWeightChanged, error) {
	event := new(TaraClientInterfaceValidatorWeightChanged)
	if err := _TaraClientInterface.contract.UnpackLog(event, "ValidatorWeightChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

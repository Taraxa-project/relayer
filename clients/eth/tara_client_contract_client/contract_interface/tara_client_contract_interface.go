// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tara_client_contract_interface

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
	BridgeRoot [32]byte
	PrevHash   [32]byte
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

// TaraClientContractInterfaceMetaData contains all meta data concerning the TaraClientContractInterface contract.
var TaraClientContractInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"HashesNotMatching\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InvalidBlockInterval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"finalizedAt\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPillarBlock.FinalizedBlock\",\"name\":\"finalized\",\"type\":\"tuple\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pillarBlockInterval\",\"type\":\"uint256\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ValidatorWeightChanged\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"change\",\"type\":\"int32\"}],\"internalType\":\"structPillarBlock.VoteCountChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"internalType\":\"structPillarBlock.WithChanges[]\",\"name\":\"blocks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature[]\",\"name\":\"lastBlockSigs\",\"type\":\"tuple[]\"}],\"name\":\"finalizeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"finalizedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalized\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bridgeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"}],\"internalType\":\"structPillarBlock.FinalizationData\",\"name\":\"block\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"finalizedAt\",\"type\":\"uint256\"}],\"internalType\":\"structPillarBlock.FinalizedBlock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalizedBridgeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"internalType\":\"structCompactSignature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"getSignaturesWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pillarBlockInterval\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pillarBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"change\",\"type\":\"int32\"}],\"internalType\":\"structPillarBlock.VoteCountChange[]\",\"name\":\"validatorChanges\",\"type\":\"tuple[]\"}],\"name\":\"processValidatorChanges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorVoteCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TaraClientContractInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use TaraClientContractInterfaceMetaData.ABI instead.
var TaraClientContractInterfaceABI = TaraClientContractInterfaceMetaData.ABI

// TaraClientContractInterface is an auto generated Go binding around an Ethereum contract.
type TaraClientContractInterface struct {
	TaraClientContractInterfaceCaller     // Read-only binding to the contract
	TaraClientContractInterfaceTransactor // Write-only binding to the contract
	TaraClientContractInterfaceFilterer   // Log filterer for contract events
}

// TaraClientContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientContractInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaraClientContractInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaraClientContractInterfaceSession struct {
	Contract     *TaraClientContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TaraClientContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaraClientContractInterfaceCallerSession struct {
	Contract *TaraClientContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// TaraClientContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaraClientContractInterfaceTransactorSession struct {
	Contract     *TaraClientContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// TaraClientContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaraClientContractInterfaceRaw struct {
	Contract *TaraClientContractInterface // Generic contract binding to access the raw methods on
}

// TaraClientContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceCallerRaw struct {
	Contract *TaraClientContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// TaraClientContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaraClientContractInterfaceTransactorRaw struct {
	Contract *TaraClientContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaraClientContractInterface creates a new instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterface(address common.Address, backend bind.ContractBackend) (*TaraClientContractInterface, error) {
	contract, err := bindTaraClientContractInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterface{TaraClientContractInterfaceCaller: TaraClientContractInterfaceCaller{contract: contract}, TaraClientContractInterfaceTransactor: TaraClientContractInterfaceTransactor{contract: contract}, TaraClientContractInterfaceFilterer: TaraClientContractInterfaceFilterer{contract: contract}}, nil
}

// NewTaraClientContractInterfaceCaller creates a new read-only instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*TaraClientContractInterfaceCaller, error) {
	contract, err := bindTaraClientContractInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceCaller{contract: contract}, nil
}

// NewTaraClientContractInterfaceTransactor creates a new write-only instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TaraClientContractInterfaceTransactor, error) {
	contract, err := bindTaraClientContractInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceTransactor{contract: contract}, nil
}

// NewTaraClientContractInterfaceFilterer creates a new log filterer instance of TaraClientContractInterface, bound to a specific deployed contract.
func NewTaraClientContractInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TaraClientContractInterfaceFilterer, error) {
	contract, err := bindTaraClientContractInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceFilterer{contract: contract}, nil
}

// bindTaraClientContractInterface binds a generic wrapper to an already deployed contract.
func bindTaraClientContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaraClientContractInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClientContractInterface *TaraClientContractInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClientContractInterface.Contract.TaraClientContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClientContractInterface *TaraClientContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.TaraClientContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClientContractInterface *TaraClientContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.TaraClientContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClientContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.contract.Transact(opts, method, params...)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32) block, uint256 finalizedAt)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) Finalized(opts *bind.CallOpts) (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "finalized")

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
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32) block, uint256 finalizedAt)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClientContractInterface.Contract.Finalized(&_TaraClientContractInterface.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32) block, uint256 finalizedAt)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClientContractInterface.Contract.Finalized(&_TaraClientContractInterface.CallOpts)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32),uint256))
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) GetFinalized(opts *bind.CallOpts) (PillarBlockFinalizedBlock, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "getFinalized")

	if err != nil {
		return *new(PillarBlockFinalizedBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(PillarBlockFinalizedBlock)).(*PillarBlockFinalizedBlock)

	return out0, err

}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32),uint256))
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClientContractInterface.Contract.GetFinalized(&_TaraClientContractInterface.CallOpts)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32),uint256))
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClientContractInterface.Contract.GetFinalized(&_TaraClientContractInterface.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) GetFinalizedBridgeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "getFinalizedBridgeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetFinalizedBridgeRoot() ([32]byte, error) {
	return _TaraClientContractInterface.Contract.GetFinalizedBridgeRoot(&_TaraClientContractInterface.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0x5a93a4bb.
//
// Solidity: function getFinalizedBridgeRoot() view returns(bytes32)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetFinalizedBridgeRoot() ([32]byte, error) {
	return _TaraClientContractInterface.Contract.GetFinalizedBridgeRoot(&_TaraClientContractInterface.CallOpts)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) GetSignaturesWeight(opts *bind.CallOpts, h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "getSignaturesWeight", h, signatures)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.GetSignaturesWeight(&_TaraClientContractInterface.CallOpts, h, signatures)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.GetSignaturesWeight(&_TaraClientContractInterface.CallOpts, h, signatures)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) Owner() (common.Address, error) {
	return _TaraClientContractInterface.Contract.Owner(&_TaraClientContractInterface.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) Owner() (common.Address, error) {
	return _TaraClientContractInterface.Contract.Owner(&_TaraClientContractInterface.CallOpts)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) PillarBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "pillarBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.PillarBlockInterval(&_TaraClientContractInterface.CallOpts)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.PillarBlockInterval(&_TaraClientContractInterface.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) Threshold() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.Threshold(&_TaraClientContractInterface.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) Threshold() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.Threshold(&_TaraClientContractInterface.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) TotalWeight() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.TotalWeight(&_TaraClientContractInterface.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) TotalWeight() (*big.Int, error) {
	return _TaraClientContractInterface.Contract.TotalWeight(&_TaraClientContractInterface.CallOpts)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCaller) ValidatorVoteCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaraClientContractInterface.contract.Call(opts, &out, "validatorVoteCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.ValidatorVoteCounts(&_TaraClientContractInterface.CallOpts, arg0)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClientContractInterface *TaraClientContractInterfaceCallerSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClientContractInterface.Contract.ValidatorVoteCounts(&_TaraClientContractInterface.CallOpts, arg0)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x508a0785.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) FinalizeBlocks(opts *bind.TransactOpts, blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "finalizeBlocks", blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x508a0785.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizeBlocks(&_TaraClientContractInterface.TransactOpts, blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x508a0785.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.FinalizeBlocks(&_TaraClientContractInterface.TransactOpts, blocks, lastBlockSigs)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _threshold, uint256 _pillarBlockInterval) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) Initialize(opts *bind.TransactOpts, _threshold *big.Int, _pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "initialize", _threshold, _pillarBlockInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _threshold, uint256 _pillarBlockInterval) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) Initialize(_threshold *big.Int, _pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.Initialize(&_TaraClientContractInterface.TransactOpts, _threshold, _pillarBlockInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _threshold, uint256 _pillarBlockInterval) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) Initialize(_threshold *big.Int, _pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.Initialize(&_TaraClientContractInterface.TransactOpts, _threshold, _pillarBlockInterval)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0x8e3aa66d.
//
// Solidity: function processValidatorChanges((address,int32)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) ProcessValidatorChanges(opts *bind.TransactOpts, validatorChanges []PillarBlockVoteCountChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "processValidatorChanges", validatorChanges)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0x8e3aa66d.
//
// Solidity: function processValidatorChanges((address,int32)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) ProcessValidatorChanges(validatorChanges []PillarBlockVoteCountChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.ProcessValidatorChanges(&_TaraClientContractInterface.TransactOpts, validatorChanges)
}

// ProcessValidatorChanges is a paid mutator transaction binding the contract method 0x8e3aa66d.
//
// Solidity: function processValidatorChanges((address,int32)[] validatorChanges) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) ProcessValidatorChanges(validatorChanges []PillarBlockVoteCountChange) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.ProcessValidatorChanges(&_TaraClientContractInterface.TransactOpts, validatorChanges)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.RenounceOwnership(&_TaraClientContractInterface.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.RenounceOwnership(&_TaraClientContractInterface.TransactOpts)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) SetThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.SetThreshold(&_TaraClientContractInterface.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.SetThreshold(&_TaraClientContractInterface.TransactOpts, _threshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaraClientContractInterface.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.TransferOwnership(&_TaraClientContractInterface.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClientContractInterface *TaraClientContractInterfaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaraClientContractInterface.Contract.TransferOwnership(&_TaraClientContractInterface.TransactOpts, newOwner)
}

// TaraClientContractInterfaceBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceBlockFinalizedIterator struct {
	Event *TaraClientContractInterfaceBlockFinalized // Event containing the contract specifics and raw log

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
func (it *TaraClientContractInterfaceBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientContractInterfaceBlockFinalized)
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
		it.Event = new(TaraClientContractInterfaceBlockFinalized)
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
func (it *TaraClientContractInterfaceBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientContractInterfaceBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientContractInterfaceBlockFinalized represents a BlockFinalized event raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceBlockFinalized struct {
	Finalized PillarBlockFinalizedBlock
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xfe3488e1f4169f43b6b09a6d3e43e05d5d4932a5baf9d378b4aac9550939c44b.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32),uint256) finalized)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*TaraClientContractInterfaceBlockFinalizedIterator, error) {

	logs, sub, err := _TaraClientContractInterface.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceBlockFinalizedIterator{contract: _TaraClientContractInterface.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xfe3488e1f4169f43b6b09a6d3e43e05d5d4932a5baf9d378b4aac9550939c44b.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32),uint256) finalized)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *TaraClientContractInterfaceBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _TaraClientContractInterface.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientContractInterfaceBlockFinalized)
				if err := _TaraClientContractInterface.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// ParseBlockFinalized is a log parse operation binding the contract event 0xfe3488e1f4169f43b6b09a6d3e43e05d5d4932a5baf9d378b4aac9550939c44b.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32),uint256) finalized)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) ParseBlockFinalized(log types.Log) (*TaraClientContractInterfaceBlockFinalized, error) {
	event := new(TaraClientContractInterfaceBlockFinalized)
	if err := _TaraClientContractInterface.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientContractInterfaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceInitializedIterator struct {
	Event *TaraClientContractInterfaceInitialized // Event containing the contract specifics and raw log

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
func (it *TaraClientContractInterfaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientContractInterfaceInitialized)
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
		it.Event = new(TaraClientContractInterfaceInitialized)
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
func (it *TaraClientContractInterfaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientContractInterfaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientContractInterfaceInitialized represents a Initialized event raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceInitialized struct {
	Threshold           *big.Int
	PillarBlockInterval *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x997896709b4e932ee42750e2e14ed1f7ab6c60bd911ff567aa2485cf0b20051b.
//
// Solidity: event Initialized(uint256 threshold, uint256 pillarBlockInterval)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaraClientContractInterfaceInitializedIterator, error) {

	logs, sub, err := _TaraClientContractInterface.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceInitializedIterator{contract: _TaraClientContractInterface.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x997896709b4e932ee42750e2e14ed1f7ab6c60bd911ff567aa2485cf0b20051b.
//
// Solidity: event Initialized(uint256 threshold, uint256 pillarBlockInterval)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaraClientContractInterfaceInitialized) (event.Subscription, error) {

	logs, sub, err := _TaraClientContractInterface.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientContractInterfaceInitialized)
				if err := _TaraClientContractInterface.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x997896709b4e932ee42750e2e14ed1f7ab6c60bd911ff567aa2485cf0b20051b.
//
// Solidity: event Initialized(uint256 threshold, uint256 pillarBlockInterval)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) ParseInitialized(log types.Log) (*TaraClientContractInterfaceInitialized, error) {
	event := new(TaraClientContractInterfaceInitialized)
	if err := _TaraClientContractInterface.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientContractInterfaceInitialized0Iterator is returned from FilterInitialized0 and is used to iterate over the raw logs and unpacked data for Initialized0 events raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceInitialized0Iterator struct {
	Event *TaraClientContractInterfaceInitialized0 // Event containing the contract specifics and raw log

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
func (it *TaraClientContractInterfaceInitialized0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientContractInterfaceInitialized0)
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
		it.Event = new(TaraClientContractInterfaceInitialized0)
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
func (it *TaraClientContractInterfaceInitialized0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientContractInterfaceInitialized0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientContractInterfaceInitialized0 represents a Initialized0 event raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceInitialized0 struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized0 is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) FilterInitialized0(opts *bind.FilterOpts) (*TaraClientContractInterfaceInitialized0Iterator, error) {

	logs, sub, err := _TaraClientContractInterface.contract.FilterLogs(opts, "Initialized0")
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceInitialized0Iterator{contract: _TaraClientContractInterface.contract, event: "Initialized0", logs: logs, sub: sub}, nil
}

// WatchInitialized0 is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) WatchInitialized0(opts *bind.WatchOpts, sink chan<- *TaraClientContractInterfaceInitialized0) (event.Subscription, error) {

	logs, sub, err := _TaraClientContractInterface.contract.WatchLogs(opts, "Initialized0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientContractInterfaceInitialized0)
				if err := _TaraClientContractInterface.contract.UnpackLog(event, "Initialized0", log); err != nil {
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

// ParseInitialized0 is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) ParseInitialized0(log types.Log) (*TaraClientContractInterfaceInitialized0, error) {
	event := new(TaraClientContractInterfaceInitialized0)
	if err := _TaraClientContractInterface.contract.UnpackLog(event, "Initialized0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientContractInterfaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceOwnershipTransferredIterator struct {
	Event *TaraClientContractInterfaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaraClientContractInterfaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientContractInterfaceOwnershipTransferred)
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
		it.Event = new(TaraClientContractInterfaceOwnershipTransferred)
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
func (it *TaraClientContractInterfaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientContractInterfaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientContractInterfaceOwnershipTransferred represents a OwnershipTransferred event raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaraClientContractInterfaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaraClientContractInterface.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceOwnershipTransferredIterator{contract: _TaraClientContractInterface.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaraClientContractInterfaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaraClientContractInterface.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientContractInterfaceOwnershipTransferred)
				if err := _TaraClientContractInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) ParseOwnershipTransferred(log types.Log) (*TaraClientContractInterfaceOwnershipTransferred, error) {
	event := new(TaraClientContractInterfaceOwnershipTransferred)
	if err := _TaraClientContractInterface.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientContractInterfaceThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceThresholdChangedIterator struct {
	Event *TaraClientContractInterfaceThresholdChanged // Event containing the contract specifics and raw log

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
func (it *TaraClientContractInterfaceThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientContractInterfaceThresholdChanged)
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
		it.Event = new(TaraClientContractInterfaceThresholdChanged)
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
func (it *TaraClientContractInterfaceThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientContractInterfaceThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientContractInterfaceThresholdChanged represents a ThresholdChanged event raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceThresholdChanged struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*TaraClientContractInterfaceThresholdChangedIterator, error) {

	logs, sub, err := _TaraClientContractInterface.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceThresholdChangedIterator{contract: _TaraClientContractInterface.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *TaraClientContractInterfaceThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _TaraClientContractInterface.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientContractInterfaceThresholdChanged)
				if err := _TaraClientContractInterface.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) ParseThresholdChanged(log types.Log) (*TaraClientContractInterfaceThresholdChanged, error) {
	event := new(TaraClientContractInterfaceThresholdChanged)
	if err := _TaraClientContractInterface.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientContractInterfaceValidatorWeightChangedIterator is returned from FilterValidatorWeightChanged and is used to iterate over the raw logs and unpacked data for ValidatorWeightChanged events raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceValidatorWeightChangedIterator struct {
	Event *TaraClientContractInterfaceValidatorWeightChanged // Event containing the contract specifics and raw log

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
func (it *TaraClientContractInterfaceValidatorWeightChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientContractInterfaceValidatorWeightChanged)
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
		it.Event = new(TaraClientContractInterfaceValidatorWeightChanged)
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
func (it *TaraClientContractInterfaceValidatorWeightChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientContractInterfaceValidatorWeightChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientContractInterfaceValidatorWeightChanged represents a ValidatorWeightChanged event raised by the TaraClientContractInterface contract.
type TaraClientContractInterfaceValidatorWeightChanged struct {
	Validator common.Address
	Weight    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorWeightChanged is a free log retrieval operation binding the contract event 0xe3490611629a41809c531d61c4adf71d3e682bfe292318e23f7227f1570d921b.
//
// Solidity: event ValidatorWeightChanged(address indexed validator, uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) FilterValidatorWeightChanged(opts *bind.FilterOpts, validator []common.Address) (*TaraClientContractInterfaceValidatorWeightChangedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraClientContractInterface.contract.FilterLogs(opts, "ValidatorWeightChanged", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraClientContractInterfaceValidatorWeightChangedIterator{contract: _TaraClientContractInterface.contract, event: "ValidatorWeightChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorWeightChanged is a free log subscription operation binding the contract event 0xe3490611629a41809c531d61c4adf71d3e682bfe292318e23f7227f1570d921b.
//
// Solidity: event ValidatorWeightChanged(address indexed validator, uint256 weight)
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) WatchValidatorWeightChanged(opts *bind.WatchOpts, sink chan<- *TaraClientContractInterfaceValidatorWeightChanged, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraClientContractInterface.contract.WatchLogs(opts, "ValidatorWeightChanged", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientContractInterfaceValidatorWeightChanged)
				if err := _TaraClientContractInterface.contract.UnpackLog(event, "ValidatorWeightChanged", log); err != nil {
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
func (_TaraClientContractInterface *TaraClientContractInterfaceFilterer) ParseValidatorWeightChanged(log types.Log) (*TaraClientContractInterfaceValidatorWeightChanged, error) {
	event := new(TaraClientContractInterfaceValidatorWeightChanged)
	if err := _TaraClientContractInterface.contract.UnpackLog(event, "ValidatorWeightChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

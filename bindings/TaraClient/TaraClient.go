// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TaraClient

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

// TaraClientMetaData contains all meta data concerning the TaraClient contract.
var TaraClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeBlocks\",\"inputs\":[{\"name\":\"blocks\",\"type\":\"tuple[]\",\"internalType\":\"structPillarBlock.WithChanges[]\",\"components\":[{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"validatorChanges\",\"type\":\"tuple[]\",\"internalType\":\"structPillarBlock.VoteCountChange[]\",\"components\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"change\",\"type\":\"int32\",\"internalType\":\"int32\"}]}]},{\"name\":\"lastBlockSigs\",\"type\":\"tuple[]\",\"internalType\":\"structCompactSignature[]\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalized\",\"inputs\":[],\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"finalizedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizedBridgeRoots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizedBlock\",\"components\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"finalizedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizedBridgeRoot\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSignaturesWeight\",\"inputs\":[{\"name\":\"h\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structCompactSignature[]\",\"components\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pillarBlockInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pillarBlockInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validatorVoteCounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BlockFinalized\",\"inputs\":[{\"name\":\"finalized\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPillarBlock.FinalizedBlock\",\"components\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block\",\"type\":\"tuple\",\"internalType\":\"structPillarBlock.FinalizationData\",\"components\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"finalizedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"HashesNotMatching\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidBlockInterval\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// TaraClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TaraClientMetaData.ABI instead.
var TaraClientABI = TaraClientMetaData.ABI

// TaraClient is an auto generated Go binding around an Ethereum contract.
type TaraClient struct {
	TaraClientCaller     // Read-only binding to the contract
	TaraClientTransactor // Write-only binding to the contract
	TaraClientFilterer   // Log filterer for contract events
}

// TaraClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaraClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaraClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaraClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaraClientSession struct {
	Contract     *TaraClient       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaraClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaraClientCallerSession struct {
	Contract *TaraClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TaraClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaraClientTransactorSession struct {
	Contract     *TaraClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TaraClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaraClientRaw struct {
	Contract *TaraClient // Generic contract binding to access the raw methods on
}

// TaraClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaraClientCallerRaw struct {
	Contract *TaraClientCaller // Generic read-only contract binding to access the raw methods on
}

// TaraClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaraClientTransactorRaw struct {
	Contract *TaraClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaraClient creates a new instance of TaraClient, bound to a specific deployed contract.
func NewTaraClient(address common.Address, backend bind.ContractBackend) (*TaraClient, error) {
	contract, err := bindTaraClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaraClient{TaraClientCaller: TaraClientCaller{contract: contract}, TaraClientTransactor: TaraClientTransactor{contract: contract}, TaraClientFilterer: TaraClientFilterer{contract: contract}}, nil
}

// NewTaraClientCaller creates a new read-only instance of TaraClient, bound to a specific deployed contract.
func NewTaraClientCaller(address common.Address, caller bind.ContractCaller) (*TaraClientCaller, error) {
	contract, err := bindTaraClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientCaller{contract: contract}, nil
}

// NewTaraClientTransactor creates a new write-only instance of TaraClient, bound to a specific deployed contract.
func NewTaraClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TaraClientTransactor, error) {
	contract, err := bindTaraClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaraClientTransactor{contract: contract}, nil
}

// NewTaraClientFilterer creates a new log filterer instance of TaraClient, bound to a specific deployed contract.
func NewTaraClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TaraClientFilterer, error) {
	contract, err := bindTaraClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaraClientFilterer{contract: contract}, nil
}

// bindTaraClient binds a generic wrapper to an already deployed contract.
func bindTaraClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaraClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClient *TaraClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClient.Contract.TaraClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClient *TaraClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClient.Contract.TaraClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClient *TaraClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClient.Contract.TaraClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraClient *TaraClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraClient *TaraClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraClient *TaraClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraClient.Contract.contract.Transact(opts, method, params...)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32,uint256) block, uint256 finalizedAt)
func (_TaraClient *TaraClientCaller) Finalized(opts *bind.CallOpts) (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "finalized")

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
func (_TaraClient *TaraClientSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClient.Contract.Finalized(&_TaraClient.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bytes32 blockHash, (uint256,bytes32,bytes32,bytes32,uint256) block, uint256 finalizedAt)
func (_TaraClient *TaraClientCallerSession) Finalized() (struct {
	BlockHash   [32]byte
	Block       PillarBlockFinalizationData
	FinalizedAt *big.Int
}, error) {
	return _TaraClient.Contract.Finalized(&_TaraClient.CallOpts)
}

// FinalizedBridgeRoots is a free data retrieval call binding the contract method 0x93ea1203.
//
// Solidity: function finalizedBridgeRoots(uint256 ) view returns(bytes32)
func (_TaraClient *TaraClientCaller) FinalizedBridgeRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "finalizedBridgeRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedBridgeRoots is a free data retrieval call binding the contract method 0x93ea1203.
//
// Solidity: function finalizedBridgeRoots(uint256 ) view returns(bytes32)
func (_TaraClient *TaraClientSession) FinalizedBridgeRoots(arg0 *big.Int) ([32]byte, error) {
	return _TaraClient.Contract.FinalizedBridgeRoots(&_TaraClient.CallOpts, arg0)
}

// FinalizedBridgeRoots is a free data retrieval call binding the contract method 0x93ea1203.
//
// Solidity: function finalizedBridgeRoots(uint256 ) view returns(bytes32)
func (_TaraClient *TaraClientCallerSession) FinalizedBridgeRoots(arg0 *big.Int) ([32]byte, error) {
	return _TaraClient.Contract.FinalizedBridgeRoots(&_TaraClient.CallOpts, arg0)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256))
func (_TaraClient *TaraClientCaller) GetFinalized(opts *bind.CallOpts) (PillarBlockFinalizedBlock, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "getFinalized")

	if err != nil {
		return *new(PillarBlockFinalizedBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(PillarBlockFinalizedBlock)).(*PillarBlockFinalizedBlock)

	return out0, err

}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256))
func (_TaraClient *TaraClientSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClient.Contract.GetFinalized(&_TaraClient.CallOpts)
}

// GetFinalized is a free data retrieval call binding the contract method 0x6b28a5e4.
//
// Solidity: function getFinalized() view returns((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256))
func (_TaraClient *TaraClientCallerSession) GetFinalized() (PillarBlockFinalizedBlock, error) {
	return _TaraClient.Contract.GetFinalized(&_TaraClient.CallOpts)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_TaraClient *TaraClientCaller) GetFinalizedBridgeRoot(opts *bind.CallOpts, epoch *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "getFinalizedBridgeRoot", epoch)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_TaraClient *TaraClientSession) GetFinalizedBridgeRoot(epoch *big.Int) ([32]byte, error) {
	return _TaraClient.Contract.GetFinalizedBridgeRoot(&_TaraClient.CallOpts, epoch)
}

// GetFinalizedBridgeRoot is a free data retrieval call binding the contract method 0xaa2bb43d.
//
// Solidity: function getFinalizedBridgeRoot(uint256 epoch) view returns(bytes32)
func (_TaraClient *TaraClientCallerSession) GetFinalizedBridgeRoot(epoch *big.Int) ([32]byte, error) {
	return _TaraClient.Contract.GetFinalizedBridgeRoot(&_TaraClient.CallOpts, epoch)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClient *TaraClientCaller) GetSignaturesWeight(opts *bind.CallOpts, h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "getSignaturesWeight", h, signatures)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClient *TaraClientSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClient.Contract.GetSignaturesWeight(&_TaraClient.CallOpts, h, signatures)
}

// GetSignaturesWeight is a free data retrieval call binding the contract method 0x97f8e2f0.
//
// Solidity: function getSignaturesWeight(bytes32 h, (bytes32,bytes32)[] signatures) view returns(uint256 weight)
func (_TaraClient *TaraClientCallerSession) GetSignaturesWeight(h [32]byte, signatures []CompactSignature) (*big.Int, error) {
	return _TaraClient.Contract.GetSignaturesWeight(&_TaraClient.CallOpts, h, signatures)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClient *TaraClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClient *TaraClientSession) Owner() (common.Address, error) {
	return _TaraClient.Contract.Owner(&_TaraClient.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaraClient *TaraClientCallerSession) Owner() (common.Address, error) {
	return _TaraClient.Contract.Owner(&_TaraClient.CallOpts)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClient *TaraClientCaller) PillarBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "pillarBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClient *TaraClientSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClient.Contract.PillarBlockInterval(&_TaraClient.CallOpts)
}

// PillarBlockInterval is a free data retrieval call binding the contract method 0x6ed0afdd.
//
// Solidity: function pillarBlockInterval() view returns(uint256)
func (_TaraClient *TaraClientCallerSession) PillarBlockInterval() (*big.Int, error) {
	return _TaraClient.Contract.PillarBlockInterval(&_TaraClient.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClient *TaraClientCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClient *TaraClientSession) TotalWeight() (*big.Int, error) {
	return _TaraClient.Contract.TotalWeight(&_TaraClient.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_TaraClient *TaraClientCallerSession) TotalWeight() (*big.Int, error) {
	return _TaraClient.Contract.TotalWeight(&_TaraClient.CallOpts)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClient *TaraClientCaller) ValidatorVoteCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TaraClient.contract.Call(opts, &out, "validatorVoteCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClient *TaraClientSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClient.Contract.ValidatorVoteCounts(&_TaraClient.CallOpts, arg0)
}

// ValidatorVoteCounts is a free data retrieval call binding the contract method 0x76a6124e.
//
// Solidity: function validatorVoteCounts(address ) view returns(uint256)
func (_TaraClient *TaraClientCallerSession) ValidatorVoteCounts(arg0 common.Address) (*big.Int, error) {
	return _TaraClient.Contract.ValidatorVoteCounts(&_TaraClient.CallOpts, arg0)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5d0d5734.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32,uint256),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClient *TaraClientTransactor) FinalizeBlocks(opts *bind.TransactOpts, blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClient.contract.Transact(opts, "finalizeBlocks", blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5d0d5734.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32,uint256),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClient *TaraClientSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClient.Contract.FinalizeBlocks(&_TaraClient.TransactOpts, blocks, lastBlockSigs)
}

// FinalizeBlocks is a paid mutator transaction binding the contract method 0x5d0d5734.
//
// Solidity: function finalizeBlocks(((uint256,bytes32,bytes32,bytes32,uint256),(address,int32)[])[] blocks, (bytes32,bytes32)[] lastBlockSigs) returns()
func (_TaraClient *TaraClientTransactorSession) FinalizeBlocks(blocks []PillarBlockWithChanges, lastBlockSigs []CompactSignature) (*types.Transaction, error) {
	return _TaraClient.Contract.FinalizeBlocks(&_TaraClient.TransactOpts, blocks, lastBlockSigs)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _pillarBlockInterval) returns()
func (_TaraClient *TaraClientTransactor) Initialize(opts *bind.TransactOpts, _pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClient.contract.Transact(opts, "initialize", _pillarBlockInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _pillarBlockInterval) returns()
func (_TaraClient *TaraClientSession) Initialize(_pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClient.Contract.Initialize(&_TaraClient.TransactOpts, _pillarBlockInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _pillarBlockInterval) returns()
func (_TaraClient *TaraClientTransactorSession) Initialize(_pillarBlockInterval *big.Int) (*types.Transaction, error) {
	return _TaraClient.Contract.Initialize(&_TaraClient.TransactOpts, _pillarBlockInterval)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClient *TaraClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraClient.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClient *TaraClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaraClient.Contract.RenounceOwnership(&_TaraClient.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaraClient *TaraClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaraClient.Contract.RenounceOwnership(&_TaraClient.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClient *TaraClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaraClient.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClient *TaraClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaraClient.Contract.TransferOwnership(&_TaraClient.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaraClient *TaraClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaraClient.Contract.TransferOwnership(&_TaraClient.TransactOpts, newOwner)
}

// TaraClientBlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the TaraClient contract.
type TaraClientBlockFinalizedIterator struct {
	Event *TaraClientBlockFinalized // Event containing the contract specifics and raw log

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
func (it *TaraClientBlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientBlockFinalized)
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
		it.Event = new(TaraClientBlockFinalized)
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
func (it *TaraClientBlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientBlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientBlockFinalized represents a BlockFinalized event raised by the TaraClient contract.
type TaraClientBlockFinalized struct {
	Finalized PillarBlockFinalizedBlock
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0xc3af94e02b408059c0f32325f41a0d0384b419789b788f11d9d3ec59fb55cd45.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256) finalized)
func (_TaraClient *TaraClientFilterer) FilterBlockFinalized(opts *bind.FilterOpts) (*TaraClientBlockFinalizedIterator, error) {

	logs, sub, err := _TaraClient.contract.FilterLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return &TaraClientBlockFinalizedIterator{contract: _TaraClient.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0xc3af94e02b408059c0f32325f41a0d0384b419789b788f11d9d3ec59fb55cd45.
//
// Solidity: event BlockFinalized((bytes32,(uint256,bytes32,bytes32,bytes32,uint256),uint256) finalized)
func (_TaraClient *TaraClientFilterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *TaraClientBlockFinalized) (event.Subscription, error) {

	logs, sub, err := _TaraClient.contract.WatchLogs(opts, "BlockFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientBlockFinalized)
				if err := _TaraClient.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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
func (_TaraClient *TaraClientFilterer) ParseBlockFinalized(log types.Log) (*TaraClientBlockFinalized, error) {
	event := new(TaraClientBlockFinalized)
	if err := _TaraClient.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaraClient contract.
type TaraClientInitializedIterator struct {
	Event *TaraClientInitialized // Event containing the contract specifics and raw log

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
func (it *TaraClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientInitialized)
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
		it.Event = new(TaraClientInitialized)
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
func (it *TaraClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientInitialized represents a Initialized event raised by the TaraClient contract.
type TaraClientInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaraClient *TaraClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaraClientInitializedIterator, error) {

	logs, sub, err := _TaraClient.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaraClientInitializedIterator{contract: _TaraClient.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TaraClient *TaraClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaraClientInitialized) (event.Subscription, error) {

	logs, sub, err := _TaraClient.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientInitialized)
				if err := _TaraClient.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaraClient *TaraClientFilterer) ParseInitialized(log types.Log) (*TaraClientInitialized, error) {
	event := new(TaraClientInitialized)
	if err := _TaraClient.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaraClient contract.
type TaraClientOwnershipTransferredIterator struct {
	Event *TaraClientOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TaraClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraClientOwnershipTransferred)
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
		it.Event = new(TaraClientOwnershipTransferred)
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
func (it *TaraClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraClientOwnershipTransferred represents a OwnershipTransferred event raised by the TaraClient contract.
type TaraClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaraClient *TaraClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaraClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaraClient.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaraClientOwnershipTransferredIterator{contract: _TaraClient.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaraClient *TaraClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaraClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaraClient.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraClientOwnershipTransferred)
				if err := _TaraClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TaraClient *TaraClientFilterer) ParseOwnershipTransferred(log types.Log) (*TaraClientOwnershipTransferred, error) {
	event := new(TaraClientOwnershipTransferred)
	if err := _TaraClient.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

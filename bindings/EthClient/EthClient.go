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

// BeaconChainBeaconBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainBeaconBlockHeader struct {
	Slot          uint64
	ProposerIndex uint64
	ParentRoot    [32]byte
	StateRoot     [32]byte
	BodyRoot      [32]byte
}

// BeaconChainExecutionPayloadHeader is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainExecutionPayloadHeader struct {
	ParentHash       [32]byte
	FeeRecipient     common.Address
	StateRoot        [32]byte
	ReceiptsRoot     [32]byte
	LogsBloom        [32]byte
	PrevRandao       [32]byte
	BlockNumber      uint64
	GasLimit         uint64
	GasUsed          uint64
	Timestamp        uint64
	ExtraData        [32]byte
	BaseFeePerGas    *big.Int
	BlockHash        [32]byte
	TransactionsRoot [32]byte
	WithdrawalsRoot  [32]byte
	BlobGasUsed      uint64
	ExcessBlobGas    uint64
}

// BeaconChainLightClientHeader is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainLightClientHeader struct {
	Beacon          BeaconChainBeaconBlockHeader
	Execution       BeaconChainExecutionPayloadHeader
	ExecutionBranch [][32]byte
}

// BeaconChainSyncCommittee is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainSyncCommittee struct {
	Pubkeys         [512][]byte
	AggregatePubkey []byte
}

// BeaconLightClientUpdateFinalizedHeaderUpdate is an auto generated low-level Go binding around an user-defined struct.
type BeaconLightClientUpdateFinalizedHeaderUpdate struct {
	AttestedHeader         BeaconChainLightClientHeader
	SignatureSyncCommittee BeaconChainSyncCommittee
	FinalizedHeader        BeaconChainLightClientHeader
	FinalityBranch         [][32]byte
	SyncAggregate          BeaconLightClientUpdateSyncAggregate
	ForkVersion            [4]byte
	SignatureSlot          uint64
}

// BeaconLightClientUpdateSyncAggregate is an auto generated low-level Go binding around an user-defined struct.
type BeaconLightClientUpdateSyncAggregate struct {
	SyncCommitteeBits      [2][32]byte
	SyncCommitteeSignature []byte
}

// BeaconLightClientUpdateSyncCommitteePeriodUpdate is an auto generated low-level Go binding around an user-defined struct.
type BeaconLightClientUpdateSyncCommitteePeriodUpdate struct {
	NextSyncCommittee       BeaconChainSyncCommittee
	NextSyncCommitteeBranch [][32]byte
}

// EthClientMetaData contains all meta data concerning the EthClient contract.
var EthClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beaconClient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBeaconLightClient\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeRootKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethBridgeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinalizedBridgeRoot\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"import_next_sync_committee\",\"inputs\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.FinalizedHeaderUpdate\",\"components\":[{\"name\":\"attested_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blob_gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excess_blob_gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"signature_sync_committee\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.SyncCommittee\",\"components\":[{\"name\":\"pubkeys\",\"type\":\"bytes[512]\",\"internalType\":\"bytes[512]\"},{\"name\":\"aggregate_pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finalized_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blob_gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excess_blob_gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"finality_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"sync_aggregate\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.SyncAggregate\",\"components\":[{\"name\":\"sync_committee_bits\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"sync_committee_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"fork_version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"signature_slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sc_update\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.SyncCommitteePeriodUpdate\",\"components\":[{\"name\":\"next_sync_committee\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.SyncCommittee\",\"components\":[{\"name\":\"pubkeys\",\"type\":\"bytes[512]\",\"internalType\":\"bytes[512]\"},{\"name\":\"aggregate_pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"next_sync_committee_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_client\",\"type\":\"address\",\"internalType\":\"contractBeaconLightClient\"},{\"name\":\"_eth_bridge_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processHeaderWithProofs\",\"inputs\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.FinalizedHeaderUpdate\",\"components\":[{\"name\":\"attested_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blob_gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excess_blob_gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"signature_sync_committee\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.SyncCommittee\",\"components\":[{\"name\":\"pubkeys\",\"type\":\"bytes[512]\",\"internalType\":\"bytes[512]\"},{\"name\":\"aggregate_pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finalized_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blob_gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excess_blob_gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"finality_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"sync_aggregate\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.SyncAggregate\",\"components\":[{\"name\":\"sync_committee_bits\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"sync_committee_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"fork_version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"signature_slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"account_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"epoch_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"root_proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sync_committee_roots\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeBeaconClient\",\"inputs\":[{\"name\":\"_client\",\"type\":\"address\",\"internalType\":\"contractBeaconLightClient\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgeRootProcessed\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bridgeRoot\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSuccessiveEpochs\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nextEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
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

// BeaconClient is a free data retrieval call binding the contract method 0xe6527ab7.
//
// Solidity: function beaconClient() view returns(address)
func (_EthClient *EthClientCaller) BeaconClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "beaconClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconClient is a free data retrieval call binding the contract method 0xe6527ab7.
//
// Solidity: function beaconClient() view returns(address)
func (_EthClient *EthClientSession) BeaconClient() (common.Address, error) {
	return _EthClient.Contract.BeaconClient(&_EthClient.CallOpts)
}

// BeaconClient is a free data retrieval call binding the contract method 0xe6527ab7.
//
// Solidity: function beaconClient() view returns(address)
func (_EthClient *EthClientCallerSession) BeaconClient() (common.Address, error) {
	return _EthClient.Contract.BeaconClient(&_EthClient.CallOpts)
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

// EpochKey is a free data retrieval call binding the contract method 0xf6c029b3.
//
// Solidity: function epochKey() view returns(bytes32)
func (_EthClient *EthClientCaller) EpochKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "epochKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EpochKey is a free data retrieval call binding the contract method 0xf6c029b3.
//
// Solidity: function epochKey() view returns(bytes32)
func (_EthClient *EthClientSession) EpochKey() ([32]byte, error) {
	return _EthClient.Contract.EpochKey(&_EthClient.CallOpts)
}

// EpochKey is a free data retrieval call binding the contract method 0xf6c029b3.
//
// Solidity: function epochKey() view returns(bytes32)
func (_EthClient *EthClientCallerSession) EpochKey() ([32]byte, error) {
	return _EthClient.Contract.EpochKey(&_EthClient.CallOpts)
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

// Slot is a free data retrieval call binding the contract method 0x1a88bc66.
//
// Solidity: function slot() view returns(uint64)
func (_EthClient *EthClientCaller) Slot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "slot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Slot is a free data retrieval call binding the contract method 0x1a88bc66.
//
// Solidity: function slot() view returns(uint64)
func (_EthClient *EthClientSession) Slot() (uint64, error) {
	return _EthClient.Contract.Slot(&_EthClient.CallOpts)
}

// Slot is a free data retrieval call binding the contract method 0x1a88bc66.
//
// Solidity: function slot() view returns(uint64)
func (_EthClient *EthClientCallerSession) Slot() (uint64, error) {
	return _EthClient.Contract.Slot(&_EthClient.CallOpts)
}

// SyncCommitteeRoots is a free data retrieval call binding the contract method 0xcb4cb856.
//
// Solidity: function sync_committee_roots(uint64 period) view returns(bytes32)
func (_EthClient *EthClientCaller) SyncCommitteeRoots(opts *bind.CallOpts, period uint64) ([32]byte, error) {
	var out []interface{}
	err := _EthClient.contract.Call(opts, &out, "sync_committee_roots", period)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SyncCommitteeRoots is a free data retrieval call binding the contract method 0xcb4cb856.
//
// Solidity: function sync_committee_roots(uint64 period) view returns(bytes32)
func (_EthClient *EthClientSession) SyncCommitteeRoots(period uint64) ([32]byte, error) {
	return _EthClient.Contract.SyncCommitteeRoots(&_EthClient.CallOpts, period)
}

// SyncCommitteeRoots is a free data retrieval call binding the contract method 0xcb4cb856.
//
// Solidity: function sync_committee_roots(uint64 period) view returns(bytes32)
func (_EthClient *EthClientCallerSession) SyncCommitteeRoots(period uint64) ([32]byte, error) {
	return _EthClient.Contract.SyncCommitteeRoots(&_EthClient.CallOpts, period)
}

// ImportNextSyncCommittee is a paid mutator transaction binding the contract method 0x474f6535.
//
// Solidity: function import_next_sync_committee((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header, ((bytes[512],bytes),bytes32[]) sc_update) returns()
func (_EthClient *EthClientTransactor) ImportNextSyncCommittee(opts *bind.TransactOpts, header BeaconLightClientUpdateFinalizedHeaderUpdate, sc_update BeaconLightClientUpdateSyncCommitteePeriodUpdate) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "import_next_sync_committee", header, sc_update)
}

// ImportNextSyncCommittee is a paid mutator transaction binding the contract method 0x474f6535.
//
// Solidity: function import_next_sync_committee((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header, ((bytes[512],bytes),bytes32[]) sc_update) returns()
func (_EthClient *EthClientSession) ImportNextSyncCommittee(header BeaconLightClientUpdateFinalizedHeaderUpdate, sc_update BeaconLightClientUpdateSyncCommitteePeriodUpdate) (*types.Transaction, error) {
	return _EthClient.Contract.ImportNextSyncCommittee(&_EthClient.TransactOpts, header, sc_update)
}

// ImportNextSyncCommittee is a paid mutator transaction binding the contract method 0x474f6535.
//
// Solidity: function import_next_sync_committee((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header, ((bytes[512],bytes),bytes32[]) sc_update) returns()
func (_EthClient *EthClientTransactorSession) ImportNextSyncCommittee(header BeaconLightClientUpdateFinalizedHeaderUpdate, sc_update BeaconLightClientUpdateSyncCommitteePeriodUpdate) (*types.Transaction, error) {
	return _EthClient.Contract.ImportNextSyncCommittee(&_EthClient.TransactOpts, header, sc_update)
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

// ProcessHeaderWithProofs is a paid mutator transaction binding the contract method 0x94620bc0.
//
// Solidity: function processHeaderWithProofs((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header, bytes[] account_proof, bytes[] epoch_proof, bytes[] root_proof) returns()
func (_EthClient *EthClientTransactor) ProcessHeaderWithProofs(opts *bind.TransactOpts, header BeaconLightClientUpdateFinalizedHeaderUpdate, account_proof [][]byte, epoch_proof [][]byte, root_proof [][]byte) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "processHeaderWithProofs", header, account_proof, epoch_proof, root_proof)
}

// ProcessHeaderWithProofs is a paid mutator transaction binding the contract method 0x94620bc0.
//
// Solidity: function processHeaderWithProofs((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header, bytes[] account_proof, bytes[] epoch_proof, bytes[] root_proof) returns()
func (_EthClient *EthClientSession) ProcessHeaderWithProofs(header BeaconLightClientUpdateFinalizedHeaderUpdate, account_proof [][]byte, epoch_proof [][]byte, root_proof [][]byte) (*types.Transaction, error) {
	return _EthClient.Contract.ProcessHeaderWithProofs(&_EthClient.TransactOpts, header, account_proof, epoch_proof, root_proof)
}

// ProcessHeaderWithProofs is a paid mutator transaction binding the contract method 0x94620bc0.
//
// Solidity: function processHeaderWithProofs((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,uint64,uint64),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header, bytes[] account_proof, bytes[] epoch_proof, bytes[] root_proof) returns()
func (_EthClient *EthClientTransactorSession) ProcessHeaderWithProofs(header BeaconLightClientUpdateFinalizedHeaderUpdate, account_proof [][]byte, epoch_proof [][]byte, root_proof [][]byte) (*types.Transaction, error) {
	return _EthClient.Contract.ProcessHeaderWithProofs(&_EthClient.TransactOpts, header, account_proof, epoch_proof, root_proof)
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

// UpgradeBeaconClient is a paid mutator transaction binding the contract method 0xb88ebb1f.
//
// Solidity: function upgradeBeaconClient(address _client) returns()
func (_EthClient *EthClientTransactor) UpgradeBeaconClient(opts *bind.TransactOpts, _client common.Address) (*types.Transaction, error) {
	return _EthClient.contract.Transact(opts, "upgradeBeaconClient", _client)
}

// UpgradeBeaconClient is a paid mutator transaction binding the contract method 0xb88ebb1f.
//
// Solidity: function upgradeBeaconClient(address _client) returns()
func (_EthClient *EthClientSession) UpgradeBeaconClient(_client common.Address) (*types.Transaction, error) {
	return _EthClient.Contract.UpgradeBeaconClient(&_EthClient.TransactOpts, _client)
}

// UpgradeBeaconClient is a paid mutator transaction binding the contract method 0xb88ebb1f.
//
// Solidity: function upgradeBeaconClient(address _client) returns()
func (_EthClient *EthClientTransactorSession) UpgradeBeaconClient(_client common.Address) (*types.Transaction, error) {
	return _EthClient.Contract.UpgradeBeaconClient(&_EthClient.TransactOpts, _client)
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
	Epoch      *big.Int
	BridgeRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgeRootProcessed is a free log retrieval operation binding the contract event 0xe8e1c85b1ad1942881c66620874e433d51385c12cb4833318b8d13af7370a9b7.
//
// Solidity: event BridgeRootProcessed(uint256 indexed epoch, bytes32 indexed bridgeRoot)
func (_EthClient *EthClientFilterer) FilterBridgeRootProcessed(opts *bind.FilterOpts, epoch []*big.Int, bridgeRoot [][32]byte) (*EthClientBridgeRootProcessedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var bridgeRootRule []interface{}
	for _, bridgeRootItem := range bridgeRoot {
		bridgeRootRule = append(bridgeRootRule, bridgeRootItem)
	}

	logs, sub, err := _EthClient.contract.FilterLogs(opts, "BridgeRootProcessed", epochRule, bridgeRootRule)
	if err != nil {
		return nil, err
	}
	return &EthClientBridgeRootProcessedIterator{contract: _EthClient.contract, event: "BridgeRootProcessed", logs: logs, sub: sub}, nil
}

// WatchBridgeRootProcessed is a free log subscription operation binding the contract event 0xe8e1c85b1ad1942881c66620874e433d51385c12cb4833318b8d13af7370a9b7.
//
// Solidity: event BridgeRootProcessed(uint256 indexed epoch, bytes32 indexed bridgeRoot)
func (_EthClient *EthClientFilterer) WatchBridgeRootProcessed(opts *bind.WatchOpts, sink chan<- *EthClientBridgeRootProcessed, epoch []*big.Int, bridgeRoot [][32]byte) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var bridgeRootRule []interface{}
	for _, bridgeRootItem := range bridgeRoot {
		bridgeRootRule = append(bridgeRootRule, bridgeRootItem)
	}

	logs, sub, err := _EthClient.contract.WatchLogs(opts, "BridgeRootProcessed", epochRule, bridgeRootRule)
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

// ParseBridgeRootProcessed is a log parse operation binding the contract event 0xe8e1c85b1ad1942881c66620874e433d51385c12cb4833318b8d13af7370a9b7.
//
// Solidity: event BridgeRootProcessed(uint256 indexed epoch, bytes32 indexed bridgeRoot)
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BeaconLightClient

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

// BeaconLightClientMetaData contains all meta data concerning the BeaconLightClient contract.
var BeaconLightClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_block_number\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_merkle_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_current_sync_committee_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_genesis_validators_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GENESIS_VALIDATORS_ROOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"block_number\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"import_finalized_header\",\"inputs\":[{\"name\":\"update\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.FinalizedHeaderUpdate\",\"components\":[{\"name\":\"attested_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"signature_sync_committee\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.SyncCommittee\",\"components\":[{\"name\":\"pubkeys\",\"type\":\"bytes[512]\",\"internalType\":\"bytes[512]\"},{\"name\":\"aggregate_pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finalized_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"finality_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"sync_aggregate\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.SyncAggregate\",\"components\":[{\"name\":\"sync_committee_bits\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"sync_committee_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"fork_version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"signature_slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"import_next_sync_committee\",\"inputs\":[{\"name\":\"header_update\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.FinalizedHeaderUpdate\",\"components\":[{\"name\":\"attested_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"signature_sync_committee\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.SyncCommittee\",\"components\":[{\"name\":\"pubkeys\",\"type\":\"bytes[512]\",\"internalType\":\"bytes[512]\"},{\"name\":\"aggregate_pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finalized_header\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.LightClientHeader\",\"components\":[{\"name\":\"beacon\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.ExecutionPayloadHeader\",\"components\":[{\"name\":\"parent_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receipts_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"logs_bloom\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prev_randao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"block_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gas_used\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extra_data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"base_fee_per_gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"block_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactions_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawals_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"execution_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"name\":\"finality_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"sync_aggregate\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.SyncAggregate\",\"components\":[{\"name\":\"sync_committee_bits\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"sync_committee_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"fork_version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"signature_slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sc_update\",\"type\":\"tuple\",\"internalType\":\"structBeaconLightClientUpdate.SyncCommitteePeriodUpdate\",\"components\":[{\"name\":\"next_sync_committee\",\"type\":\"tuple\",\"internalType\":\"structBeaconChain.SyncCommittee\",\"components\":[{\"name\":\"pubkeys\",\"type\":\"bytes[512]\",\"internalType\":\"bytes[512]\"},{\"name\":\"aggregate_pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"next_sync_committee_branch\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"merkle_root\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sync_committee_roots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FinalizedExecutionPayloadHeaderImported\",\"inputs\":[{\"name\":\"block_number\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalizedHeaderImported\",\"inputs\":[{\"name\":\"finalized_header\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBeaconChain.BeaconBlockHeader\",\"components\":[{\"name\":\"slot\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposer_index\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parent_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"body_root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NextSyncCommitteeImported\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sync_committee_root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// BeaconLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use BeaconLightClientMetaData.ABI instead.
var BeaconLightClientABI = BeaconLightClientMetaData.ABI

// BeaconLightClient is an auto generated Go binding around an Ethereum contract.
type BeaconLightClient struct {
	BeaconLightClientCaller     // Read-only binding to the contract
	BeaconLightClientTransactor // Write-only binding to the contract
	BeaconLightClientFilterer   // Log filterer for contract events
}

// BeaconLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeaconLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeaconLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeaconLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeaconLightClientSession struct {
	Contract     *BeaconLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BeaconLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeaconLightClientCallerSession struct {
	Contract *BeaconLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BeaconLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeaconLightClientTransactorSession struct {
	Contract     *BeaconLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BeaconLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeaconLightClientRaw struct {
	Contract *BeaconLightClient // Generic contract binding to access the raw methods on
}

// BeaconLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeaconLightClientCallerRaw struct {
	Contract *BeaconLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// BeaconLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeaconLightClientTransactorRaw struct {
	Contract *BeaconLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeaconLightClient creates a new instance of BeaconLightClient, bound to a specific deployed contract.
func NewBeaconLightClient(address common.Address, backend bind.ContractBackend) (*BeaconLightClient, error) {
	contract, err := bindBeaconLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconLightClient{BeaconLightClientCaller: BeaconLightClientCaller{contract: contract}, BeaconLightClientTransactor: BeaconLightClientTransactor{contract: contract}, BeaconLightClientFilterer: BeaconLightClientFilterer{contract: contract}}, nil
}

// NewBeaconLightClientCaller creates a new read-only instance of BeaconLightClient, bound to a specific deployed contract.
func NewBeaconLightClientCaller(address common.Address, caller bind.ContractCaller) (*BeaconLightClientCaller, error) {
	contract, err := bindBeaconLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconLightClientCaller{contract: contract}, nil
}

// NewBeaconLightClientTransactor creates a new write-only instance of BeaconLightClient, bound to a specific deployed contract.
func NewBeaconLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconLightClientTransactor, error) {
	contract, err := bindBeaconLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconLightClientTransactor{contract: contract}, nil
}

// NewBeaconLightClientFilterer creates a new log filterer instance of BeaconLightClient, bound to a specific deployed contract.
func NewBeaconLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconLightClientFilterer, error) {
	contract, err := bindBeaconLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconLightClientFilterer{contract: contract}, nil
}

// bindBeaconLightClient binds a generic wrapper to an already deployed contract.
func bindBeaconLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeaconLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconLightClient *BeaconLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconLightClient.Contract.BeaconLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconLightClient *BeaconLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.BeaconLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconLightClient *BeaconLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.BeaconLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconLightClient *BeaconLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconLightClient *BeaconLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconLightClient *BeaconLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.contract.Transact(opts, method, params...)
}

// GENESISVALIDATORSROOT is a free data retrieval call binding the contract method 0xa8769acb.
//
// Solidity: function GENESIS_VALIDATORS_ROOT() view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientCaller) GENESISVALIDATORSROOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BeaconLightClient.contract.Call(opts, &out, "GENESIS_VALIDATORS_ROOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GENESISVALIDATORSROOT is a free data retrieval call binding the contract method 0xa8769acb.
//
// Solidity: function GENESIS_VALIDATORS_ROOT() view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientSession) GENESISVALIDATORSROOT() ([32]byte, error) {
	return _BeaconLightClient.Contract.GENESISVALIDATORSROOT(&_BeaconLightClient.CallOpts)
}

// GENESISVALIDATORSROOT is a free data retrieval call binding the contract method 0xa8769acb.
//
// Solidity: function GENESIS_VALIDATORS_ROOT() view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientCallerSession) GENESISVALIDATORSROOT() ([32]byte, error) {
	return _BeaconLightClient.Contract.GENESISVALIDATORSROOT(&_BeaconLightClient.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x25a58b56.
//
// Solidity: function block_number() view returns(uint256)
func (_BeaconLightClient *BeaconLightClientCaller) BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeaconLightClient.contract.Call(opts, &out, "block_number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0x25a58b56.
//
// Solidity: function block_number() view returns(uint256)
func (_BeaconLightClient *BeaconLightClientSession) BlockNumber() (*big.Int, error) {
	return _BeaconLightClient.Contract.BlockNumber(&_BeaconLightClient.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x25a58b56.
//
// Solidity: function block_number() view returns(uint256)
func (_BeaconLightClient *BeaconLightClientCallerSession) BlockNumber() (*big.Int, error) {
	return _BeaconLightClient.Contract.BlockNumber(&_BeaconLightClient.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0xfd5e8efe.
//
// Solidity: function merkle_root() view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BeaconLightClient.contract.Call(opts, &out, "merkle_root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0xfd5e8efe.
//
// Solidity: function merkle_root() view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientSession) MerkleRoot() ([32]byte, error) {
	return _BeaconLightClient.Contract.MerkleRoot(&_BeaconLightClient.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0xfd5e8efe.
//
// Solidity: function merkle_root() view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientCallerSession) MerkleRoot() ([32]byte, error) {
	return _BeaconLightClient.Contract.MerkleRoot(&_BeaconLightClient.CallOpts)
}

// Slot is a free data retrieval call binding the contract method 0x1a88bc66.
//
// Solidity: function slot() view returns(uint64)
func (_BeaconLightClient *BeaconLightClientCaller) Slot(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BeaconLightClient.contract.Call(opts, &out, "slot")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Slot is a free data retrieval call binding the contract method 0x1a88bc66.
//
// Solidity: function slot() view returns(uint64)
func (_BeaconLightClient *BeaconLightClientSession) Slot() (uint64, error) {
	return _BeaconLightClient.Contract.Slot(&_BeaconLightClient.CallOpts)
}

// Slot is a free data retrieval call binding the contract method 0x1a88bc66.
//
// Solidity: function slot() view returns(uint64)
func (_BeaconLightClient *BeaconLightClientCallerSession) Slot() (uint64, error) {
	return _BeaconLightClient.Contract.Slot(&_BeaconLightClient.CallOpts)
}

// SyncCommitteeRoots is a free data retrieval call binding the contract method 0xcb4cb856.
//
// Solidity: function sync_committee_roots(uint64 ) view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientCaller) SyncCommitteeRoots(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _BeaconLightClient.contract.Call(opts, &out, "sync_committee_roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SyncCommitteeRoots is a free data retrieval call binding the contract method 0xcb4cb856.
//
// Solidity: function sync_committee_roots(uint64 ) view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientSession) SyncCommitteeRoots(arg0 uint64) ([32]byte, error) {
	return _BeaconLightClient.Contract.SyncCommitteeRoots(&_BeaconLightClient.CallOpts, arg0)
}

// SyncCommitteeRoots is a free data retrieval call binding the contract method 0xcb4cb856.
//
// Solidity: function sync_committee_roots(uint64 ) view returns(bytes32)
func (_BeaconLightClient *BeaconLightClientCallerSession) SyncCommitteeRoots(arg0 uint64) ([32]byte, error) {
	return _BeaconLightClient.Contract.SyncCommitteeRoots(&_BeaconLightClient.CallOpts, arg0)
}

// ImportFinalizedHeader is a paid mutator transaction binding the contract method 0x1666fc3f.
//
// Solidity: function import_finalized_header((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) update) returns()
func (_BeaconLightClient *BeaconLightClientTransactor) ImportFinalizedHeader(opts *bind.TransactOpts, update BeaconLightClientUpdateFinalizedHeaderUpdate) (*types.Transaction, error) {
	return _BeaconLightClient.contract.Transact(opts, "import_finalized_header", update)
}

// ImportFinalizedHeader is a paid mutator transaction binding the contract method 0x1666fc3f.
//
// Solidity: function import_finalized_header((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) update) returns()
func (_BeaconLightClient *BeaconLightClientSession) ImportFinalizedHeader(update BeaconLightClientUpdateFinalizedHeaderUpdate) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.ImportFinalizedHeader(&_BeaconLightClient.TransactOpts, update)
}

// ImportFinalizedHeader is a paid mutator transaction binding the contract method 0x1666fc3f.
//
// Solidity: function import_finalized_header((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) update) returns()
func (_BeaconLightClient *BeaconLightClientTransactorSession) ImportFinalizedHeader(update BeaconLightClientUpdateFinalizedHeaderUpdate) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.ImportFinalizedHeader(&_BeaconLightClient.TransactOpts, update)
}

// ImportNextSyncCommittee is a paid mutator transaction binding the contract method 0x2491ee63.
//
// Solidity: function import_next_sync_committee((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header_update, ((bytes[512],bytes),bytes32[]) sc_update) returns()
func (_BeaconLightClient *BeaconLightClientTransactor) ImportNextSyncCommittee(opts *bind.TransactOpts, header_update BeaconLightClientUpdateFinalizedHeaderUpdate, sc_update BeaconLightClientUpdateSyncCommitteePeriodUpdate) (*types.Transaction, error) {
	return _BeaconLightClient.contract.Transact(opts, "import_next_sync_committee", header_update, sc_update)
}

// ImportNextSyncCommittee is a paid mutator transaction binding the contract method 0x2491ee63.
//
// Solidity: function import_next_sync_committee((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header_update, ((bytes[512],bytes),bytes32[]) sc_update) returns()
func (_BeaconLightClient *BeaconLightClientSession) ImportNextSyncCommittee(header_update BeaconLightClientUpdateFinalizedHeaderUpdate, sc_update BeaconLightClientUpdateSyncCommitteePeriodUpdate) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.ImportNextSyncCommittee(&_BeaconLightClient.TransactOpts, header_update, sc_update)
}

// ImportNextSyncCommittee is a paid mutator transaction binding the contract method 0x2491ee63.
//
// Solidity: function import_next_sync_committee((((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),(bytes[512],bytes),((uint64,uint64,bytes32,bytes32,bytes32),(bytes32,address,bytes32,bytes32,bytes32,bytes32,uint64,uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32),bytes32[]),bytes32[],(bytes32[2],bytes),bytes4,uint64) header_update, ((bytes[512],bytes),bytes32[]) sc_update) returns()
func (_BeaconLightClient *BeaconLightClientTransactorSession) ImportNextSyncCommittee(header_update BeaconLightClientUpdateFinalizedHeaderUpdate, sc_update BeaconLightClientUpdateSyncCommitteePeriodUpdate) (*types.Transaction, error) {
	return _BeaconLightClient.Contract.ImportNextSyncCommittee(&_BeaconLightClient.TransactOpts, header_update, sc_update)
}

// BeaconLightClientFinalizedExecutionPayloadHeaderImportedIterator is returned from FilterFinalizedExecutionPayloadHeaderImported and is used to iterate over the raw logs and unpacked data for FinalizedExecutionPayloadHeaderImported events raised by the BeaconLightClient contract.
type BeaconLightClientFinalizedExecutionPayloadHeaderImportedIterator struct {
	Event *BeaconLightClientFinalizedExecutionPayloadHeaderImported // Event containing the contract specifics and raw log

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
func (it *BeaconLightClientFinalizedExecutionPayloadHeaderImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconLightClientFinalizedExecutionPayloadHeaderImported)
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
		it.Event = new(BeaconLightClientFinalizedExecutionPayloadHeaderImported)
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
func (it *BeaconLightClientFinalizedExecutionPayloadHeaderImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconLightClientFinalizedExecutionPayloadHeaderImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconLightClientFinalizedExecutionPayloadHeaderImported represents a FinalizedExecutionPayloadHeaderImported event raised by the BeaconLightClient contract.
type BeaconLightClientFinalizedExecutionPayloadHeaderImported struct {
	BlockNumber *big.Int
	StateRoot   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalizedExecutionPayloadHeaderImported is a free log retrieval operation binding the contract event 0x098143735ba648c9551e1ad4e3b286472943aa05f93510276fad78b342b29398.
//
// Solidity: event FinalizedExecutionPayloadHeaderImported(uint256 block_number, bytes32 state_root)
func (_BeaconLightClient *BeaconLightClientFilterer) FilterFinalizedExecutionPayloadHeaderImported(opts *bind.FilterOpts) (*BeaconLightClientFinalizedExecutionPayloadHeaderImportedIterator, error) {

	logs, sub, err := _BeaconLightClient.contract.FilterLogs(opts, "FinalizedExecutionPayloadHeaderImported")
	if err != nil {
		return nil, err
	}
	return &BeaconLightClientFinalizedExecutionPayloadHeaderImportedIterator{contract: _BeaconLightClient.contract, event: "FinalizedExecutionPayloadHeaderImported", logs: logs, sub: sub}, nil
}

// WatchFinalizedExecutionPayloadHeaderImported is a free log subscription operation binding the contract event 0x098143735ba648c9551e1ad4e3b286472943aa05f93510276fad78b342b29398.
//
// Solidity: event FinalizedExecutionPayloadHeaderImported(uint256 block_number, bytes32 state_root)
func (_BeaconLightClient *BeaconLightClientFilterer) WatchFinalizedExecutionPayloadHeaderImported(opts *bind.WatchOpts, sink chan<- *BeaconLightClientFinalizedExecutionPayloadHeaderImported) (event.Subscription, error) {

	logs, sub, err := _BeaconLightClient.contract.WatchLogs(opts, "FinalizedExecutionPayloadHeaderImported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconLightClientFinalizedExecutionPayloadHeaderImported)
				if err := _BeaconLightClient.contract.UnpackLog(event, "FinalizedExecutionPayloadHeaderImported", log); err != nil {
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

// ParseFinalizedExecutionPayloadHeaderImported is a log parse operation binding the contract event 0x098143735ba648c9551e1ad4e3b286472943aa05f93510276fad78b342b29398.
//
// Solidity: event FinalizedExecutionPayloadHeaderImported(uint256 block_number, bytes32 state_root)
func (_BeaconLightClient *BeaconLightClientFilterer) ParseFinalizedExecutionPayloadHeaderImported(log types.Log) (*BeaconLightClientFinalizedExecutionPayloadHeaderImported, error) {
	event := new(BeaconLightClientFinalizedExecutionPayloadHeaderImported)
	if err := _BeaconLightClient.contract.UnpackLog(event, "FinalizedExecutionPayloadHeaderImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeaconLightClientFinalizedHeaderImportedIterator is returned from FilterFinalizedHeaderImported and is used to iterate over the raw logs and unpacked data for FinalizedHeaderImported events raised by the BeaconLightClient contract.
type BeaconLightClientFinalizedHeaderImportedIterator struct {
	Event *BeaconLightClientFinalizedHeaderImported // Event containing the contract specifics and raw log

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
func (it *BeaconLightClientFinalizedHeaderImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconLightClientFinalizedHeaderImported)
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
		it.Event = new(BeaconLightClientFinalizedHeaderImported)
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
func (it *BeaconLightClientFinalizedHeaderImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconLightClientFinalizedHeaderImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconLightClientFinalizedHeaderImported represents a FinalizedHeaderImported event raised by the BeaconLightClient contract.
type BeaconLightClientFinalizedHeaderImported struct {
	FinalizedHeader BeaconChainBeaconBlockHeader
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFinalizedHeaderImported is a free log retrieval operation binding the contract event 0x4c6030db06afe5c2251138fd7b0c3aef3876f9f60cecfae80a2e3b9cdd3b6d5d.
//
// Solidity: event FinalizedHeaderImported((uint64,uint64,bytes32,bytes32,bytes32) finalized_header)
func (_BeaconLightClient *BeaconLightClientFilterer) FilterFinalizedHeaderImported(opts *bind.FilterOpts) (*BeaconLightClientFinalizedHeaderImportedIterator, error) {

	logs, sub, err := _BeaconLightClient.contract.FilterLogs(opts, "FinalizedHeaderImported")
	if err != nil {
		return nil, err
	}
	return &BeaconLightClientFinalizedHeaderImportedIterator{contract: _BeaconLightClient.contract, event: "FinalizedHeaderImported", logs: logs, sub: sub}, nil
}

// WatchFinalizedHeaderImported is a free log subscription operation binding the contract event 0x4c6030db06afe5c2251138fd7b0c3aef3876f9f60cecfae80a2e3b9cdd3b6d5d.
//
// Solidity: event FinalizedHeaderImported((uint64,uint64,bytes32,bytes32,bytes32) finalized_header)
func (_BeaconLightClient *BeaconLightClientFilterer) WatchFinalizedHeaderImported(opts *bind.WatchOpts, sink chan<- *BeaconLightClientFinalizedHeaderImported) (event.Subscription, error) {

	logs, sub, err := _BeaconLightClient.contract.WatchLogs(opts, "FinalizedHeaderImported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconLightClientFinalizedHeaderImported)
				if err := _BeaconLightClient.contract.UnpackLog(event, "FinalizedHeaderImported", log); err != nil {
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

// ParseFinalizedHeaderImported is a log parse operation binding the contract event 0x4c6030db06afe5c2251138fd7b0c3aef3876f9f60cecfae80a2e3b9cdd3b6d5d.
//
// Solidity: event FinalizedHeaderImported((uint64,uint64,bytes32,bytes32,bytes32) finalized_header)
func (_BeaconLightClient *BeaconLightClientFilterer) ParseFinalizedHeaderImported(log types.Log) (*BeaconLightClientFinalizedHeaderImported, error) {
	event := new(BeaconLightClientFinalizedHeaderImported)
	if err := _BeaconLightClient.contract.UnpackLog(event, "FinalizedHeaderImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeaconLightClientNextSyncCommitteeImportedIterator is returned from FilterNextSyncCommitteeImported and is used to iterate over the raw logs and unpacked data for NextSyncCommitteeImported events raised by the BeaconLightClient contract.
type BeaconLightClientNextSyncCommitteeImportedIterator struct {
	Event *BeaconLightClientNextSyncCommitteeImported // Event containing the contract specifics and raw log

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
func (it *BeaconLightClientNextSyncCommitteeImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconLightClientNextSyncCommitteeImported)
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
		it.Event = new(BeaconLightClientNextSyncCommitteeImported)
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
func (it *BeaconLightClientNextSyncCommitteeImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconLightClientNextSyncCommitteeImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconLightClientNextSyncCommitteeImported represents a NextSyncCommitteeImported event raised by the BeaconLightClient contract.
type BeaconLightClientNextSyncCommitteeImported struct {
	Period            uint64
	SyncCommitteeRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNextSyncCommitteeImported is a free log retrieval operation binding the contract event 0xeeb4e4e6976a9b191ffa4e48fccfc030141e94ef80f0fd28963b7bb4e3e31617.
//
// Solidity: event NextSyncCommitteeImported(uint64 indexed period, bytes32 indexed sync_committee_root)
func (_BeaconLightClient *BeaconLightClientFilterer) FilterNextSyncCommitteeImported(opts *bind.FilterOpts, period []uint64, sync_committee_root [][32]byte) (*BeaconLightClientNextSyncCommitteeImportedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var sync_committee_rootRule []interface{}
	for _, sync_committee_rootItem := range sync_committee_root {
		sync_committee_rootRule = append(sync_committee_rootRule, sync_committee_rootItem)
	}

	logs, sub, err := _BeaconLightClient.contract.FilterLogs(opts, "NextSyncCommitteeImported", periodRule, sync_committee_rootRule)
	if err != nil {
		return nil, err
	}
	return &BeaconLightClientNextSyncCommitteeImportedIterator{contract: _BeaconLightClient.contract, event: "NextSyncCommitteeImported", logs: logs, sub: sub}, nil
}

// WatchNextSyncCommitteeImported is a free log subscription operation binding the contract event 0xeeb4e4e6976a9b191ffa4e48fccfc030141e94ef80f0fd28963b7bb4e3e31617.
//
// Solidity: event NextSyncCommitteeImported(uint64 indexed period, bytes32 indexed sync_committee_root)
func (_BeaconLightClient *BeaconLightClientFilterer) WatchNextSyncCommitteeImported(opts *bind.WatchOpts, sink chan<- *BeaconLightClientNextSyncCommitteeImported, period []uint64, sync_committee_root [][32]byte) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var sync_committee_rootRule []interface{}
	for _, sync_committee_rootItem := range sync_committee_root {
		sync_committee_rootRule = append(sync_committee_rootRule, sync_committee_rootItem)
	}

	logs, sub, err := _BeaconLightClient.contract.WatchLogs(opts, "NextSyncCommitteeImported", periodRule, sync_committee_rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconLightClientNextSyncCommitteeImported)
				if err := _BeaconLightClient.contract.UnpackLog(event, "NextSyncCommitteeImported", log); err != nil {
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

// ParseNextSyncCommitteeImported is a log parse operation binding the contract event 0xeeb4e4e6976a9b191ffa4e48fccfc030141e94ef80f0fd28963b7bb4e3e31617.
//
// Solidity: event NextSyncCommitteeImported(uint64 indexed period, bytes32 indexed sync_committee_root)
func (_BeaconLightClient *BeaconLightClientFilterer) ParseNextSyncCommitteeImported(log types.Log) (*BeaconLightClientNextSyncCommitteeImported, error) {
	event := new(BeaconLightClientNextSyncCommitteeImported)
	if err := _BeaconLightClient.contract.UnpackLog(event, "NextSyncCommitteeImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

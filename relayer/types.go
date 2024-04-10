package relayer

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/ethereum/go-ethereum/common"
)

type BeaconBlock struct {
	Version             string `json:"version"`
	ExecutionOptimistic bool   `json:"execution_optimistic"`
	Finalized           bool   `json:"finalized"`
	Data                struct {
		Message struct {
			Slot          string                  `json:"slot"`
			ProposerIndex string                  `json:"proposer_index"`
			ParentRoot    string                  `json:"parent_root"`
			StateRoot     string                  `json:"state_root"`
			Body          capella.BeaconBlockBody `json:"body"`
		} `json:"message"`
		Signature string `json:"signature"`
	} `json:"data"`
}

type ForkVersion struct {
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized           bool `json:"finalized"`
	Data                struct {
		PreviousVersion string `json:"previous_version"`
		CurrentVersion  string `json:"current_version"`
		Epoch           string `json:"epoch"`
	} `json:"data"`
}

type SyncCommitteeUpdate struct {
	Version string   `json:"version"`
	Data    SyncData `json:"data"`
}

type AttestedHeader struct {
	Beacon Beacon `json:"beacon"`
}
type NextSyncCommittee struct {
	Pubkeys         []string `json:"pubkeys"`
	AggregatePubkey string   `json:"aggregate_pubkey"`
}
type FinalizedHeader struct {
	Beacon Beacon `json:"beacon"`
}

type SyncData struct {
	AttestedHeader          AttestedHeader    `json:"attested_header"`
	NextSyncCommittee       NextSyncCommittee `json:"next_sync_committee"`
	NextSyncCommitteeBranch []string          `json:"next_sync_committee_branch"`
	FinalizedHeader         FinalizedHeader   `json:"finalized_header"`
	FinalityBranch          []string          `json:"finality_branch"`
	SyncAggregate           SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot           string            `json:"signature_slot"`
}

type LightClientFinalityUpdate struct {
	Version string `json:"version"`
	Data    Data   `json:"data"`
}

type Data struct {
	AttestedHeader  BeaconBlockHeader `json:"attested_header"`
	FinalizedHeader BeaconBlockHeader `json:"finalized_header"`
	FinalityBranch  [][32]byte        `json:"finality_branch"`
	SyncAggregate   SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot   uint64            `json:"signature_slot"`
}

func (d *Data) UnmarshalJSON(data []byte) error {
	// Anonymous struct to mirror Data without causing recursion
	var raw struct {
		AttestedHeader  json.RawMessage `json:"attested_header"`
		FinalizedHeader json.RawMessage `json:"finalized_header"`
		FinalityBranch  []string        `json:"finality_branch"`
		SyncAggregate   json.RawMessage `json:"sync_aggregate"`
		SignatureSlot   string          `json:"signature_slot"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Unmarshal AttestedHeader
	if err := json.Unmarshal(raw.AttestedHeader, &d.AttestedHeader); err != nil {
		return fmt.Errorf("unmarshaling AttestedHeader: %v", err)
	}

	// Unmarshal FinalizedHeader
	if err := json.Unmarshal(raw.FinalizedHeader, &d.FinalizedHeader); err != nil {
		return fmt.Errorf("unmarshaling FinalizedHeader: %v", err)
	}

	// Decode FinalityBranch
	d.FinalityBranch = make([][32]byte, len(raw.FinalityBranch))
	for i, hexStr := range raw.FinalityBranch {
		cleanHexStr := strings.TrimPrefix(hexStr, "0x")
		bytes, err := hex.DecodeString(cleanHexStr)
		if err != nil {
			return fmt.Errorf("failed to decode 'FinalityBranch[%d]': %v", i, err)
		}
		if len(bytes) != 32 {
			return fmt.Errorf("decoded byte slice for 'FinalityBranch[%d]' is not 32 bytes long", i)
		}
		copy(d.FinalityBranch[i][:], bytes)
	}

	// Unmarshal SyncAggregate
	if err := json.Unmarshal(raw.SyncAggregate, &d.SyncAggregate); err != nil {
		return fmt.Errorf("unmarshaling SyncAggregate: %v", err)
	}

	// Assign SignatureSlot
	signatureSlot, err := strconv.ParseUint(raw.SignatureSlot, 10, 64)
	if err != nil {
		return fmt.Errorf("parsing SignatureSlot: %v", err)
	}
	d.SignatureSlot = signatureSlot

	return nil
}

type BeaconBlockHeader struct {
	Beacon          Beacon     `json:"beacon"`
	Execution       Execution  `json:"execution"`
	ExecutionBranch [][32]byte `json:"execution_branch"`
}

func (b *BeaconBlockHeader) UnmarshalJSON(data []byte) error {
	// Anonymous struct to mirror BeaconBlockHeader without causing recursion
	var raw struct {
		Beacon          json.RawMessage `json:"beacon"`
		Execution       json.RawMessage `json:"execution"`
		ExecutionBranch []string        `json:"execution_branch"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Unmarshal Beacon
	if err := json.Unmarshal(raw.Beacon, &b.Beacon); err != nil {
		return fmt.Errorf("unmarshaling Beacon: %v", err)
	}

	// Unmarshal Execution
	if err := json.Unmarshal(raw.Execution, &b.Execution); err != nil {
		return fmt.Errorf("unmarshaling Execution: %v", err)
	}

	// Decode ExecutionBranch
	b.ExecutionBranch = make([][32]byte, len(raw.ExecutionBranch))
	for i, hexStr := range raw.ExecutionBranch {
		cleanHexStr := strings.TrimPrefix(hexStr, "0x")
		bytes, err := hex.DecodeString(cleanHexStr)
		if err != nil {
			return fmt.Errorf("failed to decode 'ExecutionBranch[%d]': %v", i, err)
		}
		if len(bytes) != 32 {
			return fmt.Errorf("decoded byte slice for 'ExecutionBranch[%d]' is not 32 bytes long", i)
		}
		copy(b.ExecutionBranch[i][:], bytes)
	}

	return nil
}

type Beacon struct {
	Slot          uint64   `json:"slot"`
	ProposerIndex uint64   `json:"proposer_index"`
	ParentRoot    [32]byte `json:"parent_root"`
	StateRoot     [32]byte `json:"state_root"`
	BodyRoot      [32]byte `json:"body_root"`
}

// UnmarshalJSON customizes the unmarshaling of a Beacon.
func (b *Beacon) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Handle uint64 fields (Slot, ProposerIndex) directly.
	b.Slot, _ = raw["slot"].(uint64)
	b.ProposerIndex, _ = raw["proposer_index"].(uint64)

	// Convert hex string to [32]byte for ParentRoot, StateRoot, BodyRoot.
	for _, field := range []struct {
		name string
		dest *[32]byte
	}{
		{"parent_root", &b.ParentRoot},
		{"state_root", &b.StateRoot},
		{"body_root", &b.BodyRoot},
	} {
		if hexStr, ok := raw[field.name].(string); ok {
			cleanHexStr := strings.TrimPrefix(hexStr, "0x")
			bytes, err := hex.DecodeString(cleanHexStr)
			if err != nil {
				return fmt.Errorf("failed to decode '%s': %v", field.name, err)
			}
			copy(field.dest[:], bytes)
		}
	}

	return nil
}

type Execution struct {
	ParentHash       [32]byte       `json:"parent_hash"`
	FeeRecipient     common.Address `json:"fee_recipient"`
	StateRoot        [32]byte       `json:"state_root"`
	ReceiptsRoot     [32]byte       `json:"receipts_root"`
	LogsBloom        [256]byte      `json:"logs_bloom"`
	PrevRandao       [32]byte       `json:"prev_randao"`
	BlockNumber      uint64         `json:"block_number"`
	GasLimit         uint64         `json:"gas_limit"`
	GasUsed          uint64         `json:"gas_used"`
	Timestamp        uint64         `json:"timestamp"`
	ExtraData        [32]byte       `json:"extra_data"`
	BaseFeePerGas    *big.Int       `json:"base_fee_per_gas"`
	BlockHash        [32]byte       `json:"block_hash"`
	TransactionsRoot [32]byte       `json:"transactions_root"`
	WithdrawalsRoot  [32]byte       `json:"withdrawals_root"`
}

func (e *Execution) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Helper function to decode hex strings into byte arrays.
	decodeHex := func(key string, dest []byte) error {
		if hexStr, ok := raw[key].(string); ok {
			cleanHexStr := strings.TrimPrefix(hexStr, "0x")
			bytes, err := hex.DecodeString(cleanHexStr)
			if err != nil {
				return fmt.Errorf("failed to decode '%s': %v", key, err)
			}
			if len(bytes) > len(dest) {
				return fmt.Errorf("decoded byte slice for '%s' is not the expected length", key)
			}
			copy(dest, bytes)
		}
		return nil
	}

	// Decode [32]byte and [256]byte fields.
	for key, dest := range map[string][]byte{
		"parent_hash":       e.ParentHash[:],
		"state_root":        e.StateRoot[:],
		"receipts_root":     e.ReceiptsRoot[:],
		"logs_bloom":        e.LogsBloom[:],
		"prev_randao":       e.PrevRandao[:],
		"extra_data":        e.ExtraData[:],
		"block_hash":        e.BlockHash[:],
		"transactions_root": e.TransactionsRoot[:],
		"withdrawals_root":  e.WithdrawalsRoot[:],
	} {
		if err := decodeHex(key, dest); err != nil {
			return err
		}
	}

	// Convert numeric fields.
	if blockNumber, ok := raw["block_number"].(float64); ok {
		e.BlockNumber = uint64(blockNumber)
	}
	if gasLimit, ok := raw["gas_limit"].(float64); ok {
		e.GasLimit = uint64(gasLimit)
	}
	if gasUsed, ok := raw["gas_used"].(float64); ok {
		e.GasUsed = uint64(gasUsed)
	}
	if timestamp, ok := raw["timestamp"].(float64); ok {
		e.Timestamp = uint64(timestamp)
	}

	// Convert base_fee_per_gas to *big.Int
	if baseFeeStr, ok := raw["base_fee_per_gas"].(string); ok {
		e.BaseFeePerGas = big.NewInt(0)
		if _, success := e.BaseFeePerGas.SetString(strings.TrimPrefix(baseFeeStr, "0x"), 16); !success {
			return fmt.Errorf("failed to decode 'base_fee_per_gas'")
		}
	}

	// Assuming FeeRecipient needs special handling if it's not just a [20]byte.
	if feeRecipient, ok := raw["fee_recipient"].(string); ok {
		if len(feeRecipient) != 42 {
			return fmt.Errorf("invalid length for 'fee_recipient'")
		}
		if !strings.HasPrefix(feeRecipient, "0x") {
			return fmt.Errorf("missing '0x' prefix for 'fee_recipient'")
		}
		address := common.HexToAddress(feeRecipient)
		e.FeeRecipient = address
	}

	return nil
}

type SyncAggregate struct {
	SyncCommitteeBits      [64]byte `json:"sync_committee_bits"`
	SyncCommitteeSignature [96]byte `json:"sync_committee_signature"`
}

// UnmarshalJSON customizes the unmarshaling of a SyncAggregate.
func (s *SyncAggregate) UnmarshalJSON(data []byte) error {
	// Define a temporary struct where the byte arrays are represented as strings for easy unmarshaling.
	type tempSyncAggregate struct {
		SyncCommitteeBits      string `json:"sync_committee_bits"`
		SyncCommitteeSignature string `json:"sync_committee_signature"`
	}

	var temp tempSyncAggregate
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Convert the hexadecimal string for SyncCommitteeBits to a [64]byte array.
	bits, err := hexStringToByteArray(temp.SyncCommitteeBits, 64)
	if err != nil {
		return fmt.Errorf("failed to decode 'SyncCommitteeBits': %v", err)
	}
	copy(s.SyncCommitteeBits[:], bits)

	// Convert the hexadecimal string for SyncCommitteeSignature to a [96]byte array.
	signature, err := hexStringToByteArray(temp.SyncCommitteeSignature, 96)
	if err != nil {
		return fmt.Errorf("failed to decode 'SyncCommitteeSignature': %v", err)
	}
	copy(s.SyncCommitteeSignature[:], signature)

	return nil
}

// hexStringToByteArray converts a hexadecimal string to a byte array of the specified length.
func hexStringToByteArray(hexStr string, expectedLen int) ([]byte, error) {
	cleanHexStr := strings.TrimPrefix(hexStr, "0x")
	bytes, err := hex.DecodeString(cleanHexStr)
	if err != nil {
		return nil, err
	}
	if len(bytes) != expectedLen {
		return nil, fmt.Errorf("decoded byte slice is not %d bytes long", expectedLen)
	}
	return bytes, nil
}

package relayer

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
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

type NextSyncCommittee struct {
	Pubkeys         []string `json:"pubkeys"`
	AggregatePubkey string   `json:"aggregate_pubkey"`
}

type Header struct {
	Beacon phase0.BeaconBlockHeader `json:"beacon"`
}

type SyncData struct {
	AttestedHeader          Header               `json:"attested_header"`
	NextSyncCommittee       NextSyncCommittee    `json:"next_sync_committee"`
	NextSyncCommitteeBranch []string             `json:"next_sync_committee_branch"`
	FinalizedHeader         Header               `json:"finalized_header"`
	FinalityBranch          []string             `json:"finality_branch"`
	SyncAggregate           altair.SyncAggregate `json:"sync_aggregate"`
	SignatureSlot           string               `json:"signature_slot"`
}

type LightClientFinalityUpdate struct {
	Version string `json:"version"`
	Data    Data   `json:"data"`
}

type Data struct {
	AttestedHeader  BeaconBlockHeader    `json:"attested_header"`
	FinalizedHeader BeaconBlockHeader    `json:"finalized_header"`
	FinalityBranch  [][32]byte           `json:"finality_branch"`
	SyncAggregate   altair.SyncAggregate `json:"sync_aggregate"`
	SignatureSlot   uint64               `json:"signature_slot"`
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
	Beacon          phase0.BeaconBlockHeader     `json:"beacon"`
	Execution       deneb.ExecutionPayloadHeader `json:"execution"`
	ExecutionBranch [][32]byte                   `json:"execution_branch"`
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

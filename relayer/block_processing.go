package relayer

import (
	"encoding/hex"
	"log"
	"math/big"
	"strings"

	"relayer/BeaconLightClient"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBeaconBlockData(slot int64) (*BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate, error) {
	finalizedBlock, err := r.GetBlock("finalized")
	if err != nil {
		return nil, err
	}

	attestedBlock, err := r.GetBlock("head")
	if err != nil {
		return nil, err
	}

	// Fetch data from a Beacon Node API (you need to implement this based on your data source)
	// This is a placeholder for the actual implementation
	return &BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate{
		AttestedHeader: convertToBeaconChainLightClientHeader(*attestedBlock),
		// SignatureSyncCommittee: BeaconChainSyncCommittee,
		FinalizedHeader: convertToBeaconChainLightClientHeader(*finalizedBlock),
		// FinalityBranch:         [][32]byte,
		// SyncAggregate:          BeaconLightClientUpdateSyncAggregate,
		// ForkVersion:            [4]byte,
		// SignatureSlot:          uint64,
	}, nil
}

func (r *Relayer) updateNewHeader(slot int64) {
	log.Printf("Attempting to update new header for slot: %d", slot)

	// Fetch beacon block data for the given slot
	updateData, err := r.GetBeaconBlockData(slot)
	if err != nil {
		log.Printf("Failed to get beacon block data: %v", err)
		return
	}

	// Call the ImportFinalizedHeader method of the BeaconLightClient contract
	tx, err := r.beaconLightClient.ImportFinalizedHeader(r.auth, *updateData)
	if err != nil {
		log.Printf("Failed to import finalized header: %v", err)
		return
	}

	log.Printf("Submitted transaction %s for importing finalized header", tx.Hash().Hex())
}

// hexToBytes32 converts a hex string to a [32]byte array.
// It automatically strips the '0x' prefix if present.
func hexToBytes32(hexStr string) [32]byte {
	var b32 [32]byte
	// Check and remove the '0x' prefix if it's present
	cleanedHexStr := strings.TrimPrefix(hexStr, "0x")
	// Decode the hexadecimal string to bytes
	bytes, err := hex.DecodeString(cleanedHexStr)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}
	// Copy the bytes into the [32]byte array
	copy(b32[:], bytes[:32])
	return b32
}

// Convert string to uint64
func strToUint64(str string) uint64 {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Fatalf("Failed to convert string to uint64: %v", err)
	}
	return val
}

// Conversion function
func convertToBeaconChainLightClientHeader(block BeaconBlock) BeaconLightClient.BeaconChainLightClientHeader {
	beaconBlockHeader := BeaconLightClient.BeaconChainBeaconBlockHeader{
		Slot:          strToUint64(block.Data.Message.Slot),
		ProposerIndex: strToUint64(block.Data.Message.ProposerIndex),
		ParentRoot:    hexToBytes32(block.Data.Message.ParentRoot),
		StateRoot:     hexToBytes32(block.Data.Message.StateRoot),
		//BodyRoot: //TODO
	}

	// Assuming these values for demonstration; you'd extract or map these from your actual data
	exeData := block.Data.Message.Body.ExecutionPayload
	executionPayloadHeader := BeaconLightClient.BeaconChainExecutionPayloadHeader{
		ParentHash:   exeData.ParentHash,
		FeeRecipient: common.BytesToAddress(exeData.FeeRecipient[:]),
		StateRoot:    exeData.StateRoot,
		ReceiptsRoot: exeData.ReceiptsRoot,
		// LogsBloom:         [32]byte(exeData.LogsBloom), ???
		PrevRandao:    exeData.PrevRandao,
		BlockNumber:   exeData.BlockNumber,
		GasLimit:      exeData.GasLimit,
		GasUsed:       exeData.GasUsed,
		Timestamp:     exeData.Timestamp,
		ExtraData:     [32]byte(exeData.ExtraData),
		BaseFeePerGas: new(big.Int).SetBytes(exeData.BaseFeePerGas[:]),
		BlockHash:     exeData.BlockHash,
		// TransactionsRoot:  TODO
		// WithdrawalsRoot : TODO
	}

	return BeaconLightClient.BeaconChainLightClientHeader{
		Beacon:    beaconBlockHeader,
		Execution: executionPayloadHeader,
		// ExecutionBranch [][32]byte
	}
}

package relayer

import (
	"log"
	"math/big"

	"relayer/BeaconLightClient"

	"github.com/ethereum/go-ethereum/common"
)

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBeaconBlockData(slot int64) (*BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate, error) {

	finalityUpdate, err := r.GetLightClientFinalityUpdate()
	if err != nil {
		return nil, err
	}

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
		AttestedHeader: convertToBeaconChainLightClientHeader(*finalityUpdate.AttestedHeader, *attestedBlock),
		// SignatureSyncCommittee: BeaconChainSyncCommittee,
		FinalizedHeader: convertToBeaconChainLightClientHeader(*finalityUpdate.FinalizedHeader, *finalizedBlock),
		FinalityBranch:  finalityUpdate.FinalityBranch,
		// SyncAggregate:          BeaconLightClientUpdateSyncAggregate,
		// ForkVersion:            [4]byte,
		SignatureSlot: finalityUpdate.SignatureSlot,
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

// Conversion function
func convertToBeaconChainLightClientHeader(blockHeader BeaconBlockHeader, block BeaconBlock) BeaconLightClient.BeaconChainLightClientHeader {
	beaconBlockHeader := BeaconLightClient.BeaconChainBeaconBlockHeader{
		Slot:          blockHeader.Slot,
		ProposerIndex: blockHeader.ProposerIndex,
		ParentRoot:    blockHeader.ParentRoot,
		StateRoot:     blockHeader.Root,
		BodyRoot:      blockHeader.BodyRoot,
	}

	// Assuming these values for demonstration; you'd extract or map these from your actual data
	exeData := block.Data.Message.Body.ExecutionPayload
	executionPayloadHeader := BeaconLightClient.BeaconChainExecutionPayloadHeader{
		ParentHash:    exeData.ParentHash,
		FeeRecipient:  common.BytesToAddress(exeData.FeeRecipient[:]),
		StateRoot:     exeData.StateRoot,
		ReceiptsRoot:  exeData.ReceiptsRoot,
		PrevRandao:    exeData.PrevRandao,
		BlockNumber:   exeData.BlockNumber,
		GasLimit:      exeData.GasLimit,
		GasUsed:       exeData.GasUsed,
		Timestamp:     exeData.Timestamp,
		BaseFeePerGas: new(big.Int).SetBytes(exeData.BaseFeePerGas[:]),
		BlockHash:     exeData.BlockHash,
		// TransactionsRoot:  TODO
		// WithdrawalsRoot : TODO
	}

	copy(executionPayloadHeader.ExtraData[:], exeData.ExtraData)
	copy(executionPayloadHeader.LogsBloom[:], exeData.LogsBloom[:32]) //??? correct `ssz-size:"256"`

	return BeaconLightClient.BeaconChainLightClientHeader{
		Beacon:    beaconBlockHeader,
		Execution: executionPayloadHeader,
		// ExecutionBranch [][32]byte
	}
}

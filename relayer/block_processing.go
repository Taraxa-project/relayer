package relayer

import (
	"log"

	"relayer/BeaconLightClient"
)

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBeaconBlockData(slot int64) (*BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate, error) {
	finalityUpdate, err := r.GetLightClientFinalityUpdate()
	if err != nil {
		return nil, err
	}
	// Fetch data from a Beacon Node API (you need to implement this based on your data source)
	// This is a placeholder for the actual implementation
	return &BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate{
		AttestedHeader: convertToBeaconChainLightClientHeader(finalityUpdate.Data.AttestedHeader),
		// SignatureSyncCommittee: BeaconChainSyncCommittee,
		FinalizedHeader: convertToBeaconChainLightClientHeader(finalityUpdate.Data.FinalizedHeader),
		FinalityBranch:  finalityUpdate.Data.FinalityBranch,
		SyncAggregate:   ConvertSyncAggregateToBeaconLightClientUpdate(finalityUpdate.Data.SyncAggregate),
		// ForkVersion:            [4]byte,
		SignatureSlot: finalityUpdate.Data.SignatureSlot,
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
func convertToBeaconChainLightClientHeader(blockHeader BeaconBlockHeader) BeaconLightClient.BeaconChainLightClientHeader {
	beaconBlockHeader := BeaconLightClient.BeaconChainBeaconBlockHeader{
		Slot:          blockHeader.Beacon.Slot,
		ProposerIndex: blockHeader.Beacon.ProposerIndex,
		ParentRoot:    blockHeader.Beacon.ParentRoot,
		StateRoot:     blockHeader.Beacon.StateRoot,
		BodyRoot:      blockHeader.Beacon.BodyRoot,
	}

	// Assuming these values for demonstration; you'd extract or map these from your actual data
	executionPayloadHeader := BeaconLightClient.BeaconChainExecutionPayloadHeader{
		ParentHash:       blockHeader.Execution.ParentHash,
		FeeRecipient:     blockHeader.Execution.FeeRecipient,
		StateRoot:        blockHeader.Execution.StateRoot,
		ReceiptsRoot:     blockHeader.Execution.ReceiptsRoot,
		PrevRandao:       blockHeader.Execution.PrevRandao,
		BlockNumber:      blockHeader.Execution.BlockNumber,
		GasLimit:         blockHeader.Execution.GasLimit,
		GasUsed:          blockHeader.Execution.GasUsed,
		Timestamp:        blockHeader.Execution.Timestamp,
		BaseFeePerGas:    blockHeader.Execution.BaseFeePerGas,
		BlockHash:        blockHeader.Execution.BlockHash,
		TransactionsRoot: blockHeader.Execution.TransactionsRoot,
		WithdrawalsRoot:  blockHeader.Execution.WithdrawalsRoot,
		ExtraData:        blockHeader.Execution.ExtraData,
	}

	copy(executionPayloadHeader.LogsBloom[:], blockHeader.Execution.LogsBloom[:32]) //??? correct `ssz-size:"256"`

	return BeaconLightClient.BeaconChainLightClientHeader{
		Beacon:          beaconBlockHeader,
		Execution:       executionPayloadHeader,
		ExecutionBranch: blockHeader.ExecutionBranch,
	}
}

func ConvertSyncAggregateToBeaconLightClientUpdate(syncAggregate SyncAggregate) BeaconLightClient.BeaconLightClientUpdateSyncAggregate {
	var newSyncCommitteeBits [2][32]byte
	for i := 0; i < 64; i++ {
		newSyncCommitteeBits[i/32][i%32] = syncAggregate.SyncCommitteeBits[i]
	}

	return BeaconLightClient.BeaconLightClientUpdateSyncAggregate{
		SyncCommitteeBits:      newSyncCommitteeBits,
		SyncCommitteeSignature: syncAggregate.SyncCommitteeSignature[:],
	}
}

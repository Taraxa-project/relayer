package relayer

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"relayer/BeaconLightClient"

	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/ethereum/go-ethereum/common"
)

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBeaconBlockData(epoch int64) (*BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate, error) {
	finalityUpdate, err := r.GetLightClientFinalityUpdate()
	if err != nil {
		return nil, err
	}
	syncUpdate, err := r.GetSyncCommitteeUpdate(GetPeriodFromEpoch(epoch), 1)
	if err != nil {
		return nil, err
	}
	forkVersion, err := r.GetForkVersion("head")
	if err != nil {
		return nil, err
	}
	// Convert forkVersion.Data.CurrentVersion string to [4]byte
	var forkVersionBytes [4]byte

	forkBytes, err := hexStringToByteArray(forkVersion.Data.CurrentVersion, len(forkVersionBytes))
	if err != nil {
		return nil, err
	}

	copy(forkVersionBytes[:], forkBytes)

	// Fetch data from a Beacon Node API (you need to implement this based on your data source)
	// This is a placeholder for the actual implementation
	return &BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate{
		AttestedHeader:         convertToBeaconChainLightClientHeader(finalityUpdate.Data.AttestedHeader),
		SignatureSyncCommittee: ConvertToSyncCommittee(syncUpdate.Data.NextSyncCommittee),
		FinalizedHeader:        convertToBeaconChainLightClientHeader(finalityUpdate.Data.FinalizedHeader),
		FinalityBranch:         finalityUpdate.Data.FinalityBranch,
		SyncAggregate:          ConvertSyncAggregateToBeaconLightClientUpdate(finalityUpdate.Data.SyncAggregate),
		ForkVersion:            forkVersionBytes,
		SignatureSlot:          finalityUpdate.Data.SignatureSlot,
	}, nil
}

func (r *Relayer) UpdateLightClient(epoch int64, updateSyncCommittee bool) {
	log.Printf("Attempting to update new header for epoch: %d", epoch)

	// Fetch beacon block data for the given slot
	updateData, err := r.GetBeaconBlockData(epoch)
	if err != nil {
		log.Printf("Failed to get beacon block data: %v", err)
		return
	}
	if updateSyncCommittee {
		syncCommitteeData, err := r.GetSyncCommitteeData(epoch)
		if err != nil {
			log.Printf("Failed to get sync committee data: %v", err)
			return
		}
		// Call the ImportFinalizedHeader method of the BeaconLightClient contract
		tx, err := r.beaconLightClient.ImportNextSyncCommittee(r.auth, *updateData, *syncCommitteeData)
		if err != nil {
			log.Printf("Failed to import next sync committee: %v", err)
			return
		}

		log.Printf("Submitted transaction %s for importing next sync committee", tx.Hash().Hex())
	} else {
		// Call the ImportFinalizedHeader method of the BeaconLightClient contract
		tx, err := r.beaconLightClient.ImportFinalizedHeader(r.auth, *updateData)
		if err != nil {
			log.Printf("Failed to import finalized header: %v", err)
			return
		}

		log.Printf("Submitted transaction %s for importing finalized header", tx.Hash().Hex())
	}

}

func (r *Relayer) GetSyncCommitteeData(epoch int64) (*BeaconLightClient.BeaconLightClientUpdateSyncCommitteePeriodUpdate, error) {
	syncUpdate, err := r.GetSyncCommitteeUpdate(GetPeriodFromEpoch(epoch), 1)
	if err != nil {
		return nil, err
	}
	return &BeaconLightClient.BeaconLightClientUpdateSyncCommitteePeriodUpdate{
		NextSyncCommittee:       ConvertToSyncCommittee(syncUpdate.Data.NextSyncCommittee),
		NextSyncCommitteeBranch: ConvertNextSyncCommitteeBranch(syncUpdate.Data.NextSyncCommitteeBranch),
	}, nil
}

// Conversion function
func convertToBeaconChainLightClientHeader(blockHeader BeaconBlockHeader) BeaconLightClient.BeaconChainLightClientHeader {
	beaconBlockHeader := BeaconLightClient.BeaconChainBeaconBlockHeader{
		Slot:          uint64(blockHeader.Beacon.Slot),
		ProposerIndex: uint64(blockHeader.Beacon.ProposerIndex),
		ParentRoot:    blockHeader.Beacon.ParentRoot,
		StateRoot:     blockHeader.Beacon.StateRoot,
		BodyRoot:      blockHeader.Beacon.BodyRoot,
	}

	// Assuming these values for demonstration; you'd extract or map these from your actual data
	executionPayloadHeader := BeaconLightClient.BeaconChainExecutionPayloadHeader{
		ParentHash:       blockHeader.Execution.ParentHash,
		FeeRecipient:     common.Address(blockHeader.Execution.FeeRecipient),
		StateRoot:        blockHeader.Execution.StateRoot,
		ReceiptsRoot:     blockHeader.Execution.ReceiptsRoot,
		PrevRandao:       blockHeader.Execution.PrevRandao,
		BlockNumber:      blockHeader.Execution.BlockNumber,
		GasLimit:         blockHeader.Execution.GasLimit,
		GasUsed:          blockHeader.Execution.GasUsed,
		Timestamp:        blockHeader.Execution.Timestamp,
		BaseFeePerGas:    blockHeader.Execution.BaseFeePerGas.ToBig(),
		BlockHash:        blockHeader.Execution.BlockHash,
		TransactionsRoot: blockHeader.Execution.TransactionsRoot,
		WithdrawalsRoot:  blockHeader.Execution.WithdrawalsRoot,
		ExtraData:        sha256.Sum256(blockHeader.Execution.ExtraData),
		BlobGasUsed:      blockHeader.Execution.BlobGasUsed,
		ExcessBlobGas:    blockHeader.Execution.ExcessBlobGas,
		LogsBloom:        sha256.Sum256(blockHeader.Execution.LogsBloom[:]),
	}

	return BeaconLightClient.BeaconChainLightClientHeader{
		Beacon:          beaconBlockHeader,
		Execution:       executionPayloadHeader,
		ExecutionBranch: blockHeader.ExecutionBranch,
	}
}

func ConvertSyncAggregateToBeaconLightClientUpdate(syncAggregate altair.SyncAggregate) BeaconLightClient.BeaconLightClientUpdateSyncAggregate {
	var newSyncCommitteeBits [2][32]byte
	for i := 0; i < 64; i++ {
		newSyncCommitteeBits[i/32][i%32] = syncAggregate.SyncCommitteeBits[i]
	}

	return BeaconLightClient.BeaconLightClientUpdateSyncAggregate{
		SyncCommitteeBits:      newSyncCommitteeBits,
		SyncCommitteeSignature: syncAggregate.SyncCommitteeSignature[:],
	}
}

func ConvertToSyncCommittee(sc NextSyncCommittee) BeaconLightClient.BeaconChainSyncCommittee {
	var pubkeys [512][]byte

	for i, pubkey := range sc.Pubkeys {
		// Assuming the pubkey strings are prefixed with "0x" for hex encoding.
		decoded, _ := hex.DecodeString(pubkey[2:])
		pubkeys[i] = decoded
	}

	aggregatePubkey, _ := hex.DecodeString(sc.AggregatePubkey[2:])

	return BeaconLightClient.BeaconChainSyncCommittee{
		Pubkeys:         pubkeys,
		AggregatePubkey: aggregatePubkey,
	}
}

func ConvertNextSyncCommitteeBranch(input []string) [][32]byte {
	var result [][32]byte

	for _, hexStr := range input {
		// Check if the string is prefixed with "0x" and remove it
		if len(hexStr) >= 2 && hexStr[:2] == "0x" {
			hexStr = hexStr[2:]
		}

		// Decode the hex string to bytes
		bytes, _ := hex.DecodeString(hexStr)

		// Convert the byte slice to a [32]byte array
		var byteArray [32]byte
		copy(byteArray[:], bytes[:32])

		result = append(result, byteArray)
	}

	return result
}

package to_tara

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"relayer/BeaconLightClient"
	"relayer/common"

	"github.com/attestantio/go-eth2-client/spec/altair"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/herumi/bls-eth-go-binary/bls"
)

func init() {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)
}

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBeaconBlockData(epoch int64) (*BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate, error) {
	finalityUpdate, err := r.GetLightClientFinalityUpdate()
	if err != nil {
		return nil, err
	}
	syncUpdate, err := r.GetSyncCommitteeUpdate(common.GetPeriodFromEpoch(epoch)-1, 1)
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
	// print(*updateData)
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
		tx, err := r.beaconLightClient.ImportNextSyncCommittee(r.taraAuth, *updateData, *syncCommitteeData)
		if err != nil {
			log.Printf("Failed to import next sync committee: %v", err)
			return
		}

		log.Printf("Submitted transaction %s for importing next sync committee", tx.Hash().Hex())
	} else {
		// Call the ImportFinalizedHeader method of the BeaconLightClient contract
		tx, err := r.beaconLightClient.ImportFinalizedHeader(r.taraAuth, *updateData)
		if err != nil {
			log.Printf("Failed to import finalized header: %v", err)
			return
		}

		log.Printf("Submitted transaction %s for importing finalized header", tx.Hash().Hex())
	}

}

func (r *Relayer) GetSyncCommitteeData(epoch int64) (*BeaconLightClient.BeaconLightClientUpdateSyncCommitteePeriodUpdate, error) {
	syncUpdate, err := r.GetSyncCommitteeUpdate(common.GetPeriodFromEpoch(epoch), 1)
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
		FeeRecipient:     eth_common.Address(blockHeader.Execution.FeeRecipient),
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
	{
		data := ExtraData{ExtraData: blockHeader.Execution.ExtraData}
		extraData, err := data.HashTreeRoot()
		if err != nil {
			log.Fatalf("Failed to hash extra data: %v", err)
		}
		executionPayloadHeader.ExtraData = extraData
	}

	{
		data := LogsBloom{LogsBloom: blockHeader.Execution.LogsBloom}
		logBloom, err := data.HashTreeRoot()
		if err != nil {
			log.Fatalf("Failed to hash logs bloom: %v", err)
		}
		executionPayloadHeader.LogsBloom = logBloom
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

	var signature bls.Sign
	bytes := make([]byte, len(syncAggregate.SyncCommitteeSignature))
	copy(bytes, syncAggregate.SyncCommitteeSignature[:])

	if err := signature.Deserialize(bytes); err != nil {
		log.Fatalf("Failed to deserialize signature: %v", err)
	}

	return BeaconLightClient.BeaconLightClientUpdateSyncAggregate{
		SyncCommitteeBits:      newSyncCommitteeBits,
		SyncCommitteeSignature: signature.SerializeUncompressed(),
	}
}

func ConvertToSyncCommittee(sc NextSyncCommittee) BeaconLightClient.BeaconChainSyncCommittee {
	var pubkeys [512][]byte

	for i, pubkey := range sc.Pubkeys {
		// Assuming the pubkey strings are prefixed with "0x" for hex encoding.
		var key bls.PublicKey
		if err := key.DeserializeHexStr(pubkey[2:]); err != nil {
			log.Fatalf("Failed to deserialize pubkey: %v", err)
		}
		var p *bls.G1 = bls.CastFromPublicKey(&key)
		pubkeys[i] = p.SerializeUncompressed()
	}

	aggregatePubkey, _ := hex.DecodeString(sc.AggregatePubkey[2:])

	var key bls.PublicKey
	if err := key.DeserializeHexStr(sc.AggregatePubkey[2:]); err != nil {
		log.Fatalf("Failed to deserialize pubkey: %v", err)
	}

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

// Helper function to convert byte slices to hex string
func bytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// Print function for BeaconLightClientUpdateFinalizedHeaderUpdate
func print(update BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate) {
	// Print AttestedHeader fields
	fmt.Println("AttestedHeader:")
	printLightClientHeader(update.AttestedHeader)

	// Print SignatureSyncCommittee fields
	fmt.Println("SignatureSyncCommittee:")
	printSyncCommittee(update.SignatureSyncCommittee)

	// Print FinalizedHeader fields
	fmt.Println("FinalizedHeader:")
	printLightClientHeader(update.FinalizedHeader)

	// Print FinalityBranch
	fmt.Println("FinalityBranch:")
	for i, branch := range update.FinalityBranch {
		fmt.Printf("\tBranch[%d]: 0x%s\n", i, bytesToHex(branch[:]))
	}

	// Print SyncAggregate fields
	fmt.Println("SyncAggregate:")
	printSyncAggregate(update.SyncAggregate)

	// Print ForkVersion
	fmt.Printf("ForkVersion: 0x%s\n", bytesToHex(update.ForkVersion[:]))

	// Print SignatureSlot
	fmt.Printf("SignatureSlot: %d\n", update.SignatureSlot)
}

// Helper function to print LightClientHeader
// Helper function to print LightClientHeader including ExecutionPayloadHeader
func printLightClientHeader(header BeaconLightClient.BeaconChainLightClientHeader) {
	fmt.Println("\tBeacon (BeaconChainBeaconBlockHeader):")
	fmt.Printf("\t\tSlot: %d\n", header.Beacon.Slot)
	fmt.Printf("\t\tProposerIndex: %d\n", header.Beacon.ProposerIndex)
	fmt.Printf("\t\tParentRoot: 0x%s\n", bytesToHex(header.Beacon.ParentRoot[:]))
	fmt.Printf("\t\tStateRoot: 0x%s\n", bytesToHex(header.Beacon.StateRoot[:]))
	fmt.Printf("\t\tBodyRoot: 0x%s\n", bytesToHex(header.Beacon.BodyRoot[:]))

	fmt.Println("\tExecution (BeaconChainExecutionPayloadHeader):")
	fmt.Printf("\t\tParentHash: 0x%s\n", bytesToHex(header.Execution.ParentHash[:]))
	fmt.Printf("\t\tFeeRecipient: %s\n", header.Execution.FeeRecipient.Hex()) // Assuming common.Address has Hex() method
	fmt.Printf("\t\tStateRoot: 0x%s\n", bytesToHex(header.Execution.StateRoot[:]))
	fmt.Printf("\t\tReceiptsRoot: 0x%s\n", bytesToHex(header.Execution.ReceiptsRoot[:]))
	fmt.Printf("\t\tLogsBloom: 0x%s\n", bytesToHex(header.Execution.LogsBloom[:]))
	fmt.Printf("\t\tPrevRandao: 0x%s\n", bytesToHex(header.Execution.PrevRandao[:]))
	fmt.Printf("\t\tBlockNumber: %d\n", header.Execution.BlockNumber)
	fmt.Printf("\t\tGasLimit: %d\n", header.Execution.GasLimit)
	fmt.Printf("\t\tGasUsed: %d\n", header.Execution.GasUsed)
	fmt.Printf("\t\tTimestamp: %d\n", header.Execution.Timestamp)
	fmt.Printf("\t\tExtraData: 0x%s\n", bytesToHex(header.Execution.ExtraData[:]))
	if header.Execution.BaseFeePerGas != nil {
		fmt.Printf("\t\tBaseFeePerGas: %s\n", header.Execution.BaseFeePerGas.Text(10)) // Print as hexadecimal
	} else {
		fmt.Printf("\t\tBaseFeePerGas: nil\n")
	}
	fmt.Printf("\t\tBlockHash: 0x%s\n", bytesToHex(header.Execution.BlockHash[:]))
	fmt.Printf("\t\tTransactionsRoot: 0x%s\n", bytesToHex(header.Execution.TransactionsRoot[:]))
	fmt.Printf("\t\tWithdrawalsRoot: 0x%s\n", bytesToHex(header.Execution.WithdrawalsRoot[:]))
	fmt.Printf("\t\tBlobGasUsed: %d\n", header.Execution.BlobGasUsed)
	fmt.Printf("\t\tExcessBlobGas: %d\n", header.Execution.ExcessBlobGas)

	fmt.Println("\tExecutionBranch:")
	for i, branch := range header.ExecutionBranch {
		fmt.Printf("\t\tBranch[%d]: 0x%s\n", i, bytesToHex(branch[:]))
	}
}

// Helper function to print SyncCommittee
func printSyncCommittee(committee BeaconLightClient.BeaconChainSyncCommittee) {
	fmt.Println("\tPubkeys:")
	for i, pubkey := range committee.Pubkeys {
		fmt.Printf("\t\tpubkeys[%d]= hex\"%s\";\n", i, bytesToHex(pubkey))
	}
	fmt.Printf("\tAggregatePubkey: %s\n", bytesToHex(committee.AggregatePubkey))
}

// Helper function to print SyncAggregate
func printSyncAggregate(aggregate BeaconLightClient.BeaconLightClientUpdateSyncAggregate) {
	fmt.Printf("\tSyncCommitteeBits: [First Array Slice]: %s\n", bytesToHex(aggregate.SyncCommitteeBits[0][:]))
	fmt.Printf("\tSyncCommitteeBits: [Second Array Slice]: %s\n", bytesToHex(aggregate.SyncCommitteeBits[1][:]))
	// Add similar for the second array slice if needed
	fmt.Printf("\tSyncCommitteeSignature: %s\n", bytesToHex(aggregate.SyncCommitteeSignature))
}

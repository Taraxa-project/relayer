package to_tara

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"relayer/bindings/BeaconLightClient"
	"relayer/internal/common"
	"relayer/internal/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	// Fetch data from a Beacon Node API (you need to implement this based on your data source)
	// This is a placeholder for the actual implementation
	return &BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate{
		AttestedHeader:         finalityUpdate.Data.AttestedHeader.ConvertToBeaconChainLightClientHeader(r.log),
		SignatureSyncCommittee: syncUpdate.Data.NextSyncCommittee.ConvertToSyncCommittee(r.log),
		FinalizedHeader:        finalityUpdate.Data.FinalizedHeader.ConvertToBeaconChainLightClientHeader(r.log),
		FinalityBranch:         finalityUpdate.Data.FinalityBranch,
		SyncAggregate:          types.ConvertSyncAggregateToBeaconLightClientUpdate(finalityUpdate.Data.SyncAggregate),
		ForkVersion:            forkVersion,
		SignatureSlot:          finalityUpdate.Data.SignatureSlot,
	}, nil
}

func (r *Relayer) updateLightClient(epoch int64, blockNumber uint64) (*big.Int, error) {
	r.log.WithField("epoch", epoch).Info("Attempting to update new header for epoch")

	// Fetch beacon block data for the given slot
	updateData, err := r.GetBeaconBlockData(epoch)
	if blockNumber > updateData.FinalizedHeader.Execution.BlockNumber {
		return nil, fmt.Errorf("block number %d is greater than the block number in the finalized header %d", blockNumber, updateData.FinalizedHeader.Execution.BlockNumber)
	}
	// print(*updateData)
	if err != nil {
		return nil, fmt.Errorf("failed to get beacon block data: %v", err)
	}

	r.log.WithField("blockNumber", blockNumber).Info("Updating light client")

	// Call the ImportFinalizedHeader method of the BeaconLightClient contract
	tx, err := r.beaconLightClient.ImportFinalizedHeader(r.taraAuth, *updateData)
	if err != nil {
		return nil, fmt.Errorf("failed to import finalized header: %v", err)
	}

	r.log.WithField("trx", tx.Hash().Hex()).Info("Submitted next finalized header")

	receipt, err := bind.WaitMined(context.Background(), r.taraxaClient, tx)

	if err != nil {
		return nil, fmt.Errorf("failed to UpdateLightClient: %v", err)
	}

	r.log.WithField("blockNumber", receipt.BlockNumber.Uint64()).Info("Beacon chain light client updated")

	return big.NewInt(int64(updateData.FinalizedHeader.Execution.BlockNumber)), nil
}

func (r *Relayer) updateSyncCommittee(period int64) {
	go func() {
		r.log.WithField("period", period+1).Debug("Updating next sync committee")

		syncUpdate, err := r.GetSyncCommitteeUpdate(period, 1)
		if err != nil {
			r.log.WithError(err).Error("Failed to get sync committee update")
			return
		}

		oldsyncUpdate, err := r.GetSyncCommitteeUpdate(period-1, 1)
		if err != nil {
			r.log.WithError(err).Error("Failed to get previous sync committee update")
			return
		}

		forkVersion, err := r.GetForkVersion("head")
		if err != nil {
			r.log.WithError(err).Error("Failed to get fork version")
			return
		}

		finalityBranch, err := types.StringToByteArr(syncUpdate.Data.FinalityBranch)
		if err != nil {
			r.log.WithError(err).Error("Failed to convert fork version to bytes")
			return
		}

		signatureSlot, err := strconv.ParseUint(syncUpdate.Data.SignatureSlot, 10, 64)
		if err != nil {
			r.log.WithError(err).Error("Failed to parse signature slot")
			return
		}

		// Fetch beacon block data for the given slot
		updateData := BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate{
			AttestedHeader:         syncUpdate.Data.AttestedHeader.ConvertToBeaconChainLightClientHeader(r.log),
			SignatureSyncCommittee: oldsyncUpdate.Data.NextSyncCommittee.ConvertToSyncCommittee(r.log),
			FinalizedHeader:        syncUpdate.Data.FinalizedHeader.ConvertToBeaconChainLightClientHeader(r.log),
			FinalityBranch:         finalityBranch,
			SyncAggregate:          types.ConvertSyncAggregateToBeaconLightClientUpdate(syncUpdate.Data.SyncAggregate),
			ForkVersion:            forkVersion,
			SignatureSlot:          signatureSlot,
		}

		// print(updateData)

		syncCommitteeData := BeaconLightClient.BeaconLightClientUpdateSyncCommitteePeriodUpdate{
			NextSyncCommittee:       syncUpdate.Data.NextSyncCommittee.ConvertToSyncCommittee(r.log),
			NextSyncCommitteeBranch: types.ConvertNextSyncCommitteeBranch(r.log, syncUpdate.Data.NextSyncCommitteeBranch),
		}

		tx, err := r.beaconLightClient.ImportNextSyncCommittee(r.taraAuth, updateData, syncCommitteeData)
		if err != nil {
			r.log.WithError(err).Warn("Failed to import next sync committee")
			return
		}

		r.log.WithField("trx", tx.Hash().Hex()).Info("Submitted next sync committee")

		_, err = bind.WaitMined(context.Background(), r.taraxaClient, tx)

		if err != nil {
			r.log.WithError(err).Warn("Failed to update next sync committee")
		}
		r.onSyncCommitteeUpdate <- period + 1
	}()
}

// func bytesToHex(bytes []byte) string {
// 	return hex.EncodeToString(bytes)
// }

// func print(update BeaconLightClient.BeaconLightClientUpdateFinalizedHeaderUpdate) {
// 	// Print AttestedHeader fields
// 	fmt.Println("AttestedHeader:")
// 	printLightClientHeader(update.AttestedHeader)

// 	// Print SignatureSyncCommittee fields
// 	fmt.Println("SignatureSyncCommittee:")
// 	printSyncCommittee(update.SignatureSyncCommittee)

// 	// Print FinalizedHeader fields
// 	fmt.Println("FinalizedHeader:")
// 	printLightClientHeader(update.FinalizedHeader)

// 	// Print FinalityBranch
// 	fmt.Println("FinalityBranch:")
// 	for i, branch := range update.FinalityBranch {
// 		fmt.Printf("\tBranch[%d]: 0x%s\n", i, bytesToHex(branch[:]))
// 	}

// 	// Print SyncAggregate fields
// 	fmt.Println("SyncAggregate:")
// 	printSyncAggregate(update.SyncAggregate)

// 	// Print ForkVersion
// 	fmt.Printf("ForkVersion: 0x%s\n", bytesToHex(update.ForkVersion[:]))

// 	// Print SignatureSlot
// 	fmt.Printf("SignatureSlot: %d\n", update.SignatureSlot)
// }

// func printLightClientHeader(header BeaconLightClient.BeaconChainLightClientHeader) {
// 	fmt.Println("\tBeacon (BeaconChainBeaconBlockHeader):")
// 	fmt.Printf("\t\tSlot: %d\n", header.Beacon.Slot)
// 	fmt.Printf("\t\tProposerIndex: %d\n", header.Beacon.ProposerIndex)
// 	fmt.Printf("\t\tParentRoot: 0x%s\n", bytesToHex(header.Beacon.ParentRoot[:]))
// 	fmt.Printf("\t\tStateRoot: 0x%s\n", bytesToHex(header.Beacon.StateRoot[:]))
// 	fmt.Printf("\t\tBodyRoot: 0x%s\n", bytesToHex(header.Beacon.BodyRoot[:]))

// 	fmt.Println("\tExecution (BeaconChainExecutionPayloadHeader):")
// 	fmt.Printf("\t\tParentHash: 0x%s\n", bytesToHex(header.Execution.ParentHash[:]))
// 	fmt.Printf("\t\tFeeRecipient: %s\n", header.Execution.FeeRecipient.Hex()) // Assuming common.Address has Hex() method
// 	fmt.Printf("\t\tStateRoot: 0x%s\n", bytesToHex(header.Execution.StateRoot[:]))
// 	fmt.Printf("\t\tReceiptsRoot: 0x%s\n", bytesToHex(header.Execution.ReceiptsRoot[:]))
// 	fmt.Printf("\t\tLogsBloom: 0x%s\n", bytesToHex(header.Execution.LogsBloom[:]))
// 	fmt.Printf("\t\tPrevRandao: 0x%s\n", bytesToHex(header.Execution.PrevRandao[:]))
// 	fmt.Printf("\t\tBlockNumber: %d\n", header.Execution.BlockNumber)
// 	fmt.Printf("\t\tGasLimit: %d\n", header.Execution.GasLimit)
// 	fmt.Printf("\t\tGasUsed: %d\n", header.Execution.GasUsed)
// 	fmt.Printf("\t\tTimestamp: %d\n", header.Execution.Timestamp)
// 	fmt.Printf("\t\tExtraData: 0x%s\n", bytesToHex(header.Execution.ExtraData[:]))
// 	if header.Execution.BaseFeePerGas != nil {
// 		fmt.Printf("\t\tBaseFeePerGas: %s\n", header.Execution.BaseFeePerGas.Text(10)) // Print as hexadecimal
// 	} else {
// 		fmt.Printf("\t\tBaseFeePerGas: nil\n")
// 	}
// 	fmt.Printf("\t\tBlockHash: 0x%s\n", bytesToHex(header.Execution.BlockHash[:]))
// 	fmt.Printf("\t\tTransactionsRoot: 0x%s\n", bytesToHex(header.Execution.TransactionsRoot[:]))
// 	fmt.Printf("\t\tWithdrawalsRoot: 0x%s\n", bytesToHex(header.Execution.WithdrawalsRoot[:]))
// 	fmt.Printf("\t\tBlobGasUsed: %d\n", header.Execution.BlobGasUsed)
// 	fmt.Printf("\t\tExcessBlobGas: %d\n", header.Execution.ExcessBlobGas)

// 	fmt.Println("\tExecutionBranch:")
// 	for i, branch := range header.ExecutionBranch {
// 		fmt.Printf("\t\tBranch[%d]: 0x%s\n", i, bytesToHex(branch[:]))
// 	}
// }

// func printSyncCommittee(committee BeaconLightClient.BeaconChainSyncCommittee) {
// 	fmt.Println("\tPubkeys:")
// 	for i, pubkey := range committee.Pubkeys {
// 		fmt.Printf("\t\tpubkeys[%d]= hex\"%s\";\n", i, bytesToHex(pubkey))
// 	}
// 	fmt.Printf("\tAggregatePubkey: %s\n", bytesToHex(committee.AggregatePubkey))
// }

// func printSyncAggregate(aggregate BeaconLightClient.BeaconLightClientUpdateSyncAggregate) {
// 	fmt.Printf("\tSyncCommitteeBits: [First Array Slice]: %s\n", bytesToHex(aggregate.SyncCommitteeBits[0][:]))
// 	fmt.Printf("\tSyncCommitteeBits: [Second Array Slice]: %s\n", bytesToHex(aggregate.SyncCommitteeBits[1][:]))
// 	// Add similar for the second array slice if needed
// 	fmt.Printf("\tSyncCommitteeSignature: %s\n", bytesToHex(aggregate.SyncCommitteeSignature))
// }

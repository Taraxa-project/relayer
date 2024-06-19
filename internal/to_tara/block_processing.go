package to_tara

import (
	"context"
	"strconv"

	"relayer/bindings/EthClient"
	"relayer/internal/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/herumi/bls-eth-go-binary/bls"
)

func init() {
	_ = bls.Init(bls.BLS12_381)
	_ = bls.SetETHmode(bls.EthModeDraft07)
}

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBeaconBlockData(epoch int64) (*EthClient.BeaconLightClientUpdateFinalizedHeaderUpdate, error) {
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
	return &EthClient.BeaconLightClientUpdateFinalizedHeaderUpdate{
		AttestedHeader:         convertToBeaconChainLightClientHeader(r.log, finalityUpdate.Data.AttestedHeader),
		SignatureSyncCommittee: ConvertToSyncCommittee(r.log, syncUpdate.Data.NextSyncCommittee),
		FinalizedHeader:        convertToBeaconChainLightClientHeader(r.log, finalityUpdate.Data.FinalizedHeader),
		FinalityBranch:         finalityUpdate.Data.FinalityBranch,
		SyncAggregate:          ConvertSyncAggregateToBeaconLightClientUpdate(finalityUpdate.Data.SyncAggregate),
		ForkVersion:            forkVersion,
		SignatureSlot:          finalityUpdate.Data.SignatureSlot,
	}, nil
}

func (r *Relayer) updateSyncCommittee(period int64) {
	go func() {
		r.log.WithField("period", period+1).Info("Updating next sync committee")

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

		finalityBranch, err := stringToByteArr(syncUpdate.Data.FinalityBranch)
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
		updateData := EthClient.BeaconLightClientUpdateFinalizedHeaderUpdate{
			AttestedHeader:         convertToBeaconChainLightClientHeader(r.log, syncUpdate.Data.AttestedHeader),
			SignatureSyncCommittee: ConvertToSyncCommittee(r.log, oldsyncUpdate.Data.NextSyncCommittee),
			FinalizedHeader:        convertToBeaconChainLightClientHeader(r.log, syncUpdate.Data.FinalizedHeader),
			FinalityBranch:         finalityBranch,
			SyncAggregate:          ConvertSyncAggregateToBeaconLightClientUpdate(syncUpdate.Data.SyncAggregate),
			ForkVersion:            forkVersion,
			SignatureSlot:          signatureSlot,
		}

		// print(updateData)

		syncCommitteeData := EthClient.BeaconLightClientUpdateSyncCommitteePeriodUpdate{
			NextSyncCommittee:       ConvertToSyncCommittee(r.log, syncUpdate.Data.NextSyncCommittee),
			NextSyncCommitteeBranch: ConvertNextSyncCommitteeBranch(r.log, syncUpdate.Data.NextSyncCommitteeBranch),
		}

		tx, err := r.ethClientContract.ImportNextSyncCommittee(r.taraAuth, updateData, syncCommitteeData)
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

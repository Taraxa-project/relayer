package relayer

const SLOTS_PER_EPOCH = 32
const EPOCHS_PER_SYNC_COMMITTEE_PERIOD = 256

// GetSlotFromEpoch calculates the first slot number of a given epoch.
func GetSlotFromEpoch(epoch int64) int64 {
	return epoch * SLOTS_PER_EPOCH
}

func GetPeriodFromEpoch(epoch int64) int64 {
	return epoch / EPOCHS_PER_SYNC_COMMITTEE_PERIOD
}

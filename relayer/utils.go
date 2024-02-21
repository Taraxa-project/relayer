package relayer

const SLOTS_PER_EPOCH = 32

// GetSlotFromEpoch calculates the first slot number of a given epoch.
func GetSlotFromEpoch(epoch int64) int64 {
	return epoch * SLOTS_PER_EPOCH
}

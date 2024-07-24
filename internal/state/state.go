package state

import (
	"relayer/internal/types"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type State struct {
	stakes      map[common.Address]int32
	latestEpoch uint64
	totalStake  int32
	stakeGetter func(common.Address) int32
}

func NewState(getter func(common.Address) int32) *State {
	return &State{
		stakes:      make(map[common.Address]int32),
		stakeGetter: getter,
	}
}

func (s *State) UpdateState(block *types.PillarBlock) {
	if s.latestEpoch > uint64(block.Epoch) {
		return
	}
	s.latestEpoch = uint64(block.Epoch)
	for _, change := range block.VoteCountsChanges {
		if _, ok := s.stakes[change.Address]; !ok {
			stake := s.stakeGetter(change.Address)
			s.stakes[change.Address] = stake
			s.totalStake += stake
		}
		s.stakes[change.Address] += change.Value
		s.totalStake += change.Value
	}
}

type AccountWithSignature struct {
	Address   *common.Address
	Signature *types.CompactSignature
}

func (s *State) ReduceSignatures(block *types.PillarBlockData) ([]types.CompactSignature, error) {
	var totalStake int32
	accounts := make([]AccountWithSignature, 0, len(block.Signatures))
	for _, signature := range block.Signatures {
		pubKey, err := crypto.Ecrecover(block.PillarBlock.Hash[:], signature.ToCanonical())
		if err != nil {
			return nil, err
		}

		addr := common.BytesToAddress(common.BytesToHash(pubKey[1:]).Bytes()[12:])
		accounts = append(accounts, AccountWithSignature{Address: &addr, Signature: &signature})
	}

	for _, acc := range accounts {
		totalStake += s.stakes[*acc.Address]
	}

	threshold := s.totalStake*2/3 + 1

	if totalStake <= threshold {
		return block.Signatures, nil
	}

	sort.Slice(accounts, func(i, j int) bool { return s.stakes[*accounts[i].Address] > s.stakes[*accounts[j].Address] })

	var reducedSignatures []types.CompactSignature
	for _, acc := range accounts {
		threshold -= s.stakes[*acc.Address]
		reducedSignatures = append(reducedSignatures, *acc.Signature)
		if threshold <= 0 {
			break
		}
	}

	return reducedSignatures, nil
}

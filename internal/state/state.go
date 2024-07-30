package state

import (
	"relayer/bindings/TaraClient"
	relayer_common "relayer/internal/common"
	"relayer/internal/types"

	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type State struct {
	stakes          map[common.Address]int32
	latestPbftBlock uint64
	totalWeight     int32
	stakeGetter     func(common.Address) int32
}

func NewState(totalWeight int32, getter func(common.Address) int32) *State {
	return &State{
		stakes:      make(map[common.Address]int32),
		totalWeight: totalWeight,
		stakeGetter: getter,
	}
}

func (s *State) initAddressStake(address common.Address) {
	if _, ok := s.stakes[address]; !ok {
		s.stakes[address] = s.stakeGetter(address)
	}
}

func (s *State) GetStake(address common.Address) int32 {
	s.initAddressStake(address)
	return s.stakes[address]
}

func (s *State) UpdateStake(address common.Address, stake int32) {
	s.initAddressStake(address)
	s.stakes[address] += stake
	s.totalWeight += stake
}

func (s *State) UpdateState(block *types.PillarBlock) {
	if s.latestPbftBlock > uint64(block.PbftPeriod) {
		return
	}
	s.latestPbftBlock = uint64(block.PbftPeriod)
	for _, change := range block.VoteCountsChanges {
		s.UpdateStake(change.Address, change.Value)
	}
}

type AccountWithSignature struct {
	Address   *common.Address
	Signature types.CompactSignature
}

func ConvertToSortedSignatures(signatures []types.CompactSignature) []TaraClient.CompactSignature {
	sort.Slice(signatures, func(i, j int) bool { return signatures[i].R.Cmp(signatures[j].R) > 0 })

	tcSignatures := make([]TaraClient.CompactSignature, len(signatures))
	for i, signature := range signatures {
		tcSignatures[i] = TaraClient.CompactSignature{R: signature.R, Vs: signature.Vs}
	}
	return tcSignatures
}

func (s *State) ReduceSignatures(block *types.PillarBlockData) ([]TaraClient.CompactSignature, error) {
	var sigsStake int32
	accounts := make([]AccountWithSignature, 0, len(block.Signatures))
	pvh := block.GetVoteHash()
	for _, signature := range block.Signatures {
		pubKey, err := crypto.Ecrecover(pvh[:], signature.ToCanonical())
		if err != nil {
			return nil, err
		}
		addr := relayer_common.PubkeyToAddress(pubKey)

		accounts = append(accounts, AccountWithSignature{Address: &addr, Signature: signature})
		sigsStake += s.GetStake(addr)
	}

	threshold := s.totalWeight/2 + 1

	if sigsStake < threshold {
		panic("Not enough stake to reduce signatures")
	}

	if sigsStake == threshold {
		return ConvertToSortedSignatures(block.Signatures), nil
	}

	sort.Slice(accounts, func(i, j int) bool { return s.stakes[*accounts[i].Address] > s.stakes[*accounts[j].Address] })

	var reducedSignatures []types.CompactSignature
	for _, acc := range accounts {
		threshold -= s.stakes[*acc.Address]
		reducedSignatures = append(reducedSignatures, acc.Signature)
		if threshold <= 0 {
			break
		}
	}

	return ConvertToSortedSignatures(reducedSignatures), nil
}

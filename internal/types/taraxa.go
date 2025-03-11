package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

var (
	voteHash, _ = abi.NewType("tuple", "vote hash", []abi.ArgumentMarshaling{
		{Name: "period", Type: "uint256"},
		{Name: "hash", Type: "bytes32"},
	})

	args = abi.Arguments{
		abi.Argument{
			Name:    "",
			Type:    voteHash,
			Indexed: false,
		},
	}
)

type VoteCountChange struct {
	Address common.Address `json:"address"   `
	Value   int32          `json:"value"     `
}

type CompactSignature struct {
	R  common.Hash `json:"r"`
	Vs common.Hash `json:"vs"`
}

// PillarBlock represents a pillar block in the taraxa blockchain.
type PillarBlock struct {
	PbftPeriod        hexutil.Uint64    `json:"pbft_period"`
	StateRoot         common.Hash       `json:"state_root"`
	PreviousBlockHash common.Hash       `json:"previous_pillar_block_hash"`
	BridgeRoot        common.Hash       `json:"bridge_root"`
	Epoch             hexutil.Uint64    `json:"epoch"`
	VoteCountsChanges []VoteCountChange `json:"validators_vote_counts_changes"`
	Hash              common.Hash       `json:"hash"`
}

type PillarBlockData struct {
	PillarBlock PillarBlock        `json:"pillar_block"`
	Signatures  []CompactSignature `json:"signatures"`
}

type VoteHashArgs struct {
	Period *big.Int
	Hash   common.Hash
}

func (p *PillarBlockData) GetVoteHash() []byte {
	vote := VoteHashArgs{
		Period: big.NewInt(int64(p.PillarBlock.PbftPeriod + 1)),
		Hash:   p.PillarBlock.Hash,
	}
	packed, err := args.Pack(&vote)
	if err != nil {
		log.WithError(err).Panic("failed to pack vote hash")
	}

	vote_hash := crypto.Keccak256(packed)
	return vote_hash
}

// Config - parsed only partially, new fields can be added anytime
type TaraConfig struct {
	ChainId   uint64          `json:"chain_id"  rlp:"required"`
	Hardforks HardforksConfig `json:"hardforks" rlp:"required"`
}

type HardforksConfig struct {
	FicusHf FicusHfConfig `json:"ficus_hf"  rlp:"required"`
}

type FicusHfConfig struct {
	BlockNum                hexutil.Uint64 `json:"block_num"                   rlp:"required"`
	PbftInclusionDelay      hexutil.Uint64 `json:"pbft_inclusion_delay"        rlp:"required"`
	PillarBlocksInterval    hexutil.Uint64 `json:"pillar_blocks_interval"      rlp:"required"`
	PillarChainSyncInterval hexutil.Uint64 `json:"pillar_chain_sync_interval"  rlp:"required"`
}

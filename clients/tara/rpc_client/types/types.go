package tara_rpc_types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type VoteCountChange struct {
	Address common.Address `json:"address"   rlp:"required"`
	Value   int32          `json:"value"     rlp:"required"`
}

type CompactSignature struct {
	R  common.Hash `json:"r"  rlp:"required"`
	Vs common.Hash `json:"vs" rlp:"required"`
}

// PillarBlock represents a pillar block in the taraxa blockchain.
type PillarBlock struct {
	Hash              common.Hash       `json:"hash"                            rlp:"required"`
	PreviousBlockHash common.Hash       `json:"previous_pillar_block_hash"      rlp:"required"`
	BridgeRoot        common.Hash       `json:"bridge_root"                     rlp:"required"`
	StateRoot         common.Hash       `json:"state_root"                      rlp:"required"`
	PbftPeriod        uint64            `json:"pbft_period"                     rlp:"required"`
	VoteCountsChanges []VoteCountChange `json:"validators_vote_counts_changes"  rlp:"optional"`
}

type PillarBlockData struct {
	PillarBlock PillarBlock        `json:"pillar_block"  rlp:"required"`
	Signatures  []CompactSignature `json:"signatures"    rlp:"optional"`
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

package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"relayer/bindings/BeaconLightClient"
	"relayer/bindings/TaraClient"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/attestantio/go-eth2-client/spec/altair"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/herumi/bls-eth-go-binary/bls"
)

func (pillarBlockData *PillarBlockData) TransformPillarBlockData() (block TaraClient.PillarBlockWithChanges, signatures []TaraClient.CompactSignature) {
	block.Block.Period = big.NewInt(int64(pillarBlockData.PillarBlock.PbftPeriod))
	block.Block.BridgeRoot = pillarBlockData.PillarBlock.BridgeRoot
	block.Block.StateRoot = pillarBlockData.PillarBlock.StateRoot
	block.Block.PrevHash = pillarBlockData.PillarBlock.PreviousBlockHash
	block.Block.Epoch = big.NewInt(int64(pillarBlockData.PillarBlock.Epoch))
	for _, votesCountChange := range pillarBlockData.PillarBlock.VoteCountsChanges {
		block.ValidatorChanges = append(block.ValidatorChanges, TaraClient.PillarBlockVoteCountChange{Validator: votesCountChange.Address, Change: votesCountChange.Value})
	}

	// sort signatures by R value in descending order
	sort.Slice(pillarBlockData.Signatures, func(i, j int) bool {
		return pillarBlockData.Signatures[i].R.Cmp(pillarBlockData.Signatures[j].R) > 0
	})

	for _, signature := range pillarBlockData.Signatures {
		signatures = append(signatures, TaraClient.CompactSignature{R: signature.R, Vs: signature.Vs})
	}

	return
}

// / hexStringToByteArray converts a hexadecimal string to a byte array of the specified length.
func HexStringToByteArray(hexStr string, expectedLen int) ([]byte, error) {
	cleanHexStr := strings.TrimPrefix(hexStr, "0x")
	bytes, err := hex.DecodeString(cleanHexStr)
	if err != nil {
		return nil, err
	}
	if len(bytes) != expectedLen {
		return nil, fmt.Errorf("decoded byte slice is not %d bytes long", expectedLen)
	}
	return bytes, nil
}

// stringToByteArr converts a slice of hexadecimal strings to a slice of [32]byte arrays.
func StringToByteArr(hexStrings []string) ([][32]byte, error) {
	byteArr := make([][32]byte, len(hexStrings))
	for i, hexStr := range hexStrings {
		cleanHexStr := strings.TrimPrefix(hexStr, "0x")
		bytes, err := hex.DecodeString(cleanHexStr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode 'FinalityBranch[%d]': %v", i, err)
		}
		if len(bytes) != 32 {
			return nil, fmt.Errorf("decoded byte slice for 'FinalityBranch[%d]' is not 32 bytes long", i)
		}
		copy(byteArr[i][:], bytes)
	}
	return byteArr, nil
}

// convertToBeaconChainLightClientHeader converts a BeaconBlockHeader to a BeaconChainLightClientHeader.
func (blockHeader *BeaconBlockHeader) ConvertToBeaconChainLightClientHeader(log *log.Logger) BeaconLightClient.BeaconChainLightClientHeader {
	beaconBlockHeader := BeaconLightClient.BeaconChainBeaconBlockHeader{
		Slot:          uint64(blockHeader.Beacon.Slot),
		ProposerIndex: uint64(blockHeader.Beacon.ProposerIndex),
		ParentRoot:    blockHeader.Beacon.ParentRoot,
		StateRoot:     blockHeader.Beacon.StateRoot,
		BodyRoot:      blockHeader.Beacon.BodyRoot,
	}

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

// ConvertSyncAggregateToBeaconLightClientUpdate converts a SyncAggregate to BeaconLightClientUpdateSyncAggregate.
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

// ConvertToSyncCommittee converts a NextSyncCommittee to BeaconChainSyncCommittee.
func (sc *NextSyncCommittee) ConvertToSyncCommittee(log *log.Logger) BeaconLightClient.BeaconChainSyncCommittee {
	var pubkeys [512][]byte

	for i, pubkey := range sc.Pubkeys {
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
		log.Fatalf("Failed to deserialize aggregate pubkey: %v", err)
	}

	return BeaconLightClient.BeaconChainSyncCommittee{
		Pubkeys:         pubkeys,
		AggregatePubkey: aggregatePubkey,
	}
}

// ConvertNextSyncCommitteeBranch converts a slice of hexadecimal strings to a slice of [32]byte arrays.
func ConvertNextSyncCommitteeBranch(log *log.Logger, input []string) [][32]byte {
	var result [][32]byte

	for _, hexStr := range input {
		if len(hexStr) >= 2 && hexStr[:2] == "0x" {
			hexStr = hexStr[2:]
		}

		bytes, err := hex.DecodeString(hexStr)
		if err != nil {
			log.Fatalf("Failed to decode hex string: %v", err)
		}

		var byteArray [32]byte
		copy(byteArray[:], bytes[:32])

		result = append(result, byteArray)
	}

	return result
}

func (cs *CompactSignature) ToCanonical() []byte {
	rInt := new(big.Int).SetBytes(cs.R[:])

	// Mask to obtain the lowest 255 bits for s
	sMask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 255), big.NewInt(1))
	s := new(big.Int).And(cs.Vs.Big(), sMask)

	// Shift right to obtain yParity (the 256th bit)
	yParity := new(big.Int).Rsh(cs.Vs.Big(), 255).Bytes()

	if len(yParity) == 0 {
		yParity = []byte{0}
	}

	yParity[0] += 27
	return append(rInt.Bytes(), append(s.Bytes(), yParity[0])...)
}

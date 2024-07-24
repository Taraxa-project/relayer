package types_test

import (
	"reflect"
	"relayer/internal/types"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// Test case from https://github.com/ethereum/ercs/blob/master/ERCS/erc-2098.md#test-cases
func TestERC2098One(t *testing.T) {
	cs := types.CompactSignature{
		R:  common.HexToHash("68a020a209d3d56c46f38cc50a33f704f4a9a10a59377f8dd762ac66910e9b90"),
		Vs: common.HexToHash("7e865ad05c4035ab5792787d4a0297a43617ae897930a6fe4d822b8faea52064"),
	}

	expected := common.Hex2Bytes("68a020a209d3d56c46f38cc50a33f704f4a9a10a59377f8dd762ac66910e9b907e865ad05c4035ab5792787d4a0297a43617ae897930a6fe4d822b8faea520641b")
	got := cs.ToCanonical()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ToCanonical() = %v, want %v", got, expected)
	}

}

func TestERC2098Two(t *testing.T) {
	cs := types.CompactSignature{
		R:  common.HexToHash("9328da16089fcba9bececa81663203989f2df5fe1faa6291a45381c81bd17f76"),
		Vs: common.HexToHash("939c6d6b623b42da56557e5e734a43dc83345ddfadec52cbe24d0cc64f550793"),
	}

	expected := common.Hex2Bytes("9328da16089fcba9bececa81663203989f2df5fe1faa6291a45381c81bd17f76139c6d6b623b42da56557e5e734a43dc83345ddfadec52cbe24d0cc64f5507931c")
	got := cs.ToCanonical()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ToCanonical() = %v, want %v", got, expected)
	}
}

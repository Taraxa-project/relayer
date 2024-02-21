package relayer

import "github.com/attestantio/go-eth2-client/spec/capella"

type BeaconBlock struct {
	Version             string `json:"version"`
	ExecutionOptimistic bool   `json:"execution_optimistic"`
	Finalized           bool   `json:"finalized"`
	Data                struct {
		Message struct {
			Slot          string                  `json:"slot"`
			ProposerIndex string                  `json:"proposer_index"`
			ParentRoot    string                  `json:"parent_root"`
			StateRoot     string                  `json:"state_root"`
			Body          capella.BeaconBlockBody `json:"body"`
		} `json:"message"`
		Signature string `json:"signature"`
	} `json:"data"`
}

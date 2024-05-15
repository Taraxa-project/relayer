package to_tara

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBlockHeader(blockType string) (*BeaconBlock, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/headers/%s", r.beaconNodeEndpoint, blockType)
	var beaconBlockHeader *BeaconBlock
	beaconBlockHeader, err := FetchAndParseData[BeaconBlock](url)
	if err != nil {
		log.Fatalf("Error fetching and parsing beacon block header: %v", err)
		return nil, err
	}

	return beaconBlockHeader, nil
}

// Assume GetBeaconBlockData returns data needed to construct BeaconLightClientUpdateFinalizedHeaderUpdate
func (r *Relayer) GetBlock(slot string) (*BeaconBlock, error) {
	url := fmt.Sprintf("%s/eth/v2/beacon/blocks/%s", r.beaconNodeEndpoint, slot)
	var beaconBlockHeader *BeaconBlock
	beaconBlockHeader, err := FetchAndParseData[BeaconBlock](url)
	if err != nil {
		log.Fatalf("Error fetching and parsing beacon block header: %v", err)
		return nil, err
	}

	return beaconBlockHeader, nil
}

func (r *Relayer) GetForkVersion(state string) (*ForkVersion, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/states/%s/fork", r.lightNodeEndpoint, state)
	var forkVersion *ForkVersion
	forkVersion, err := FetchAndParseData[ForkVersion](url)
	if err != nil {
		log.Fatalf("Error fetching and parsing fork version: %v", err)
		return nil, err
	}

	return forkVersion, nil
}

func (r *Relayer) GetLightClientFinalityUpdate() (*LightClientFinalityUpdate, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/finality_update", r.lightNodeEndpoint)
	var finalityUpdate *LightClientFinalityUpdate
	finalityUpdate, err := FetchAndParseData[LightClientFinalityUpdate](url)
	if err != nil {
		log.Fatalf("Error fetching and parsing finality header: %v", err)
		return nil, err
	}

	return finalityUpdate, nil
}

func (r *Relayer) GetSyncCommitteeUpdate(startPeriod, count int64) (*SyncCommitteeUpdate, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/light_client/updates?start_period=%d&count=%d", r.lightNodeEndpoint, startPeriod, count)
	var syncUpdates *[]SyncCommitteeUpdate
	syncUpdates, err := FetchAndParseData[[]SyncCommitteeUpdate](url)
	if err != nil {
		log.Fatalf("Error fetching and parsing sync committee updates: %v", err)
		return nil, err
	}

	return &(*syncUpdates)[0], nil
}

// FetchAndParseData fetches data from a given URL and parses the JSON response into the provided type.
// T is a placeholder type that will be inferred from the caller, allowing for any struct type to be used.
func FetchAndParseData[T any](url string) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var responseData T
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response JSON: %v", err)
	}

	return &responseData, nil
}

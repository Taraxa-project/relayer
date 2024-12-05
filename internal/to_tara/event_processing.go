package to_tara

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (r *Relayer) startEventProcessing(ctx context.Context) {
	client := &http.Client{}

	// Construct the request to the Ethereum 2.0 node's event stream
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/eth/v1/events?topics=finalized_checkpoint", r.beaconNodeEndpoint), nil)
	if err != nil {
		r.log.WithError(err).Panic("Failed to create request")
	}
	req.Header.Add("accept", "text/event-stream")

	// Make the request and receive the response
	resp, err := client.Do(req)
	if err != nil {
		r.log.WithError(err).Panic("Failed to connect to event stream")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		r.log.WithField("statusCode", resp.StatusCode).Panic("Failed to subscribe to events")
	}

	// Assuming the use of a generic SSE client to parse the stream.
	// You need to replace this with actual code to listen and process the SSE stream.
	// Process the Server-Sent Events stream
	r.processSSEStream(resp.Body)
}

func (r *Relayer) processSSEStream(stream io.ReadCloser) {
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			dataLine := line[5:] // Skip "data:" prefix

			var subscriptionData map[string]interface{}
			if err := json.Unmarshal([]byte(dataLine), &subscriptionData); err != nil {
				r.log.WithError(err).Error("Error parsing JSON")
				continue
			}

			switch epoch := subscriptionData["epoch"].(type) {
			case float64:
				// JSON numbers are decoded into float64 by default
				r.onFinalizedEpoch <- int64(epoch)
			case int:
				// Handle int if by any chance it's parsed as such
				r.onFinalizedEpoch <- int64(epoch)
			case string:
				// If "epoch" is provided as a string, parse it to an integer
				if epochVal, err := strconv.ParseUint(epoch, 10, 64); err == nil {
					r.onFinalizedEpoch <- int64(epochVal)
					r.log.WithField("epoch", epochVal).Debug("Epoch value")
				} else {
					r.log.WithError(err).Error("Error converting epoch from string to uint64")
				}
			default:
				r.log.WithField("data", subscriptionData).Warn("Epoch value is of an unrecognized type")
			}
		}
		if err := scanner.Err(); err != nil {
			r.log.WithError(err).Error("Error reading stream")
		}
	}
}

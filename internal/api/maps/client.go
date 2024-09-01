package maps

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type PlaceId string
type Mode string

const (
	DRIVE   Mode = "driving"
	TRANSIT Mode = "transit"
)

type MapClient struct {
	apiKey string
}

func NewMapClient(apiKey string) *MapClient {
	return &MapClient{apiKey: apiKey}
}

func (c *MapClient) Fetch(commuteReq *CommuteTimeRequest) (*DistanceMatrixResponse, error) {
	if commuteReq.mode != DRIVE && commuteReq.mode != TRANSIT {
		return nil, errors.New("invalid mode")
	}

	commuteReq.apiKey = c.apiKey

	url, err := commuteReq.GetHttpRequestUrl()
	if err != nil {
		return nil, fmt.Errorf("something went wrong when building request url: %w", err)
	}

	resp, err := http.Get(*url)
	if err != nil {
		return nil, fmt.Errorf("something went wrong when fetching data: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var distanceMatrix DistanceMatrixResponse
	err = json.Unmarshal(body, &distanceMatrix)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	if distanceMatrix.Status != "OK" {
		return nil, fmt.Errorf("API returned non-OK status: %s", distanceMatrix.Status)
	}

	return &distanceMatrix, nil
}

package maps

import (
	"errors"
	"strings"
)

const baseUrl string = "https://maps.googleapis.com/maps/api/distancematrix/json"

type CommuteTimeRequest struct {
	apiKey string
	from   []PlaceId
	to     []PlaceId
	mode   Mode
}

func (r *CommuteTimeRequest) GetHttpRequestUrl() (*string, error) {
	if len(r.from) == 0 || len(r.to) == 0 {
		return nil, errors.New("missing origin and destination inputs")
	}

	var urlBuilder strings.Builder
	urlBuilder.WriteString(baseUrl)

	urlBuilder.WriteString("?origins=")
	for idx, from := range r.from {
		urlBuilder.WriteString("place_id:")
		urlBuilder.WriteString(string(from))
		if idx != len(r.from)-1 {
			urlBuilder.WriteString("|")
		}
	}

	urlBuilder.WriteString("&destinations=")
	for idx, to := range r.to {
		urlBuilder.WriteString("place_id:")
		urlBuilder.WriteString(string(to))
		if idx != len(r.to)-1 {
			urlBuilder.WriteString("|")
		}
	}

	urlBuilder.WriteString("&mode=")
	urlBuilder.WriteString(string(r.mode))

	urlBuilder.WriteString("&departure_time=now")
	urlBuilder.WriteString("&units=imperial")

	urlBuilder.WriteString("&key=")
	urlBuilder.WriteString(r.apiKey)

	url := urlBuilder.String()
	return &url, nil
}

type DistanceMatrixResponse struct {
	DestinationAddresses []PlaceId `json:"destination_addresses"`
	OriginAddresses      []PlaceId `json:"origin_addresses"`
	Rows                 []Row     `json:"rows"`
	Status               string    `json:"status"`
}

type Row struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Distance Distance `json:"distance"`
	Duration Duration `json:"duration"`
	Status   string   `json:"status"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type Duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

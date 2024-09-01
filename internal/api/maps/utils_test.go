package maps

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommuteTimeRequest_GetHttpRequestUrl(t *testing.T) {
	commuteTimeRequest := &CommuteTimeRequest{
		apiKey: "abc123",
		from:   []PlaceId{"ChIJWZMPlaQ92jERDPUVvwyFdms"},
		to:     []PlaceId{"ChIJ59pmmlMa2jER7vHtsApwXTc"},
		mode:   DRIVE,
	}

	expectedUrl := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=place_id:ChIJWZMPlaQ92jERDPUVvwyFdms&destinations=place_id:ChIJ59pmmlMa2jER7vHtsApwXTc&mode=driving&departure_time=now&units=imperial&key=abc123"
	url, err := commuteTimeRequest.GetHttpRequestUrl()
	require.NoError(t, err)

	assert.Equal(t, expectedUrl, *url)
}

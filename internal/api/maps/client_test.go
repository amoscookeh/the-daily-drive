package maps

import (
	"github.com/amoscookeh/the-daily-drive/internal/utils"
	"testing"
)

var tests = map[string]struct {
	input *CommuteTimeRequest

	expResp *DistanceMatrixResponse
	expErr  error
}{
	"happy path": {
		input: &CommuteTimeRequest{
			from: []PlaceId{"ChIJWZMPlaQ92jERDPUVvwyFdms"},
			to:   []PlaceId{"ChIJ59pmmlMa2jER7vHtsApwXTc"},
			mode: DRIVE,
		},
		expErr: nil,
	},
}

func TestCommuteTimeRequest(t *testing.T) {
	envpath := "../../../.env"
	utils.SetupEnv(&envpath)

	mapClient := MapClient{
		apiKey: utils.GetMapClientApiKey(),
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			resp, err := mapClient.Fetch(tc.input)
			t.Logf("INFO: %v", resp)
			if err != tc.expErr {
				t.Fatalf("expected: %v, got: %v", tc.expErr, err)
			}
		})
	}
}

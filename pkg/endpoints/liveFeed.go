package endpoints

import (
	"fmt"
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/pkg/common"
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/protos/live_feed"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log"
)

func GetLiveFlights() (string, error) {
	liveFlightsRequest := &live_feed.LiveFeedRequest{
		Bounds: &live_feed.LocationBoundaries{
			North: 53.0,
			South: 45.0,
			West:  -43.0,
			East:  -20.0,
		},
		FieldMask: &fieldmaskpb.FieldMask{
			Paths: []string{
				"flights_list",
				"stats",
			},
		},
	}

	data, err := common.EncodeGRPCMessage(liveFlightsRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	return "", nil
}

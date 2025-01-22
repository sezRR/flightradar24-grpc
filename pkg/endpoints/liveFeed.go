package endpoints

import (
	"bytes"
	"fmt"
	pkgCommon "github.com/sezrr/go-playground/examples/protobuf-flightradar24/pkg/common"
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/protos/common"
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/protos/live_feed"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"io"
	"log"
	"net/http"
)

type BoundingBox struct {
	North float32
	South float32
	West  float32
	East  float32
}

type LiveFeedParams struct {
	BoundingBox BoundingBox
	Stats       bool
	Limit       int32
	MaxAge      int32
	Fields      map[string]struct{} // Using a map for a set equivalent in Go
}

func NewLiveFeedParams() *LiveFeedParams {
	return &LiveFeedParams{
		BoundingBox: BoundingBox{
			North: 49.0, // Example default (BBOX_FRANCE_UIR)
			South: 42.0,
			West:  -5.0,
			East:  8.0,
		},
		Stats:  true,
		Limit:  1500,
		MaxAge: 14400,
		Fields: map[string]struct{}{
			"flight": {}, "reg": {}, "route": {}, "type": {},
		},
	}
}

func (params *LiveFeedParams) ToProto() *live_feed.LiveFeedRequest {
	// Convert the map (set equivalent) to a slice of strings
	fields := make([]string, 0, len(params.Fields))
	for field := range params.Fields {
		fields = append(fields, field)
	}

	// Create the LiveFeedRequest protobuf object
	return &live_feed.LiveFeedRequest{
		Bounds: &live_feed.LocationBoundaries{
			North: params.BoundingBox.North,
			South: params.BoundingBox.South,
			West:  params.BoundingBox.West,
			East:  params.BoundingBox.East,
		},
		Settings: &live_feed.VisibilitySettings{
			SourcesList:    []common.DataSource{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			ServicesList:   []common.Service{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			TrafficType:    common.TrafficType_ALL,
			OnlyRestricted: proto.Bool(false), // Direct pointer for optional bool
		},
		FieldMask: &fieldmaskpb.FieldMask{
			Paths: fields,
		},
		HighlightMode:   false,
		Stats:           &params.Stats,                                   // Use address for optional bool
		Limit:           &params.Limit,                                   // Use address for optional int32
		Maxage:          &params.MaxAge,                                  // Use address for optional int32
		RestrictionMode: common.RestrictionVisibility_NOT_VISIBLE.Enum(), // Use Enum() method for enums
	}
}
func GetLiveFeed() (*live_feed.LiveFeedResponse, error) {
	liveFlightsRequest := NewLiveFeedParams().ToProto()

	protoRequest, err := pkgCommon.EncodeGRPCMessage(liveFlightsRequest)
	if err != nil {
		log.Fatal(err)
	}

	url := "https://data-feed.flightradar24.com/fr24.feed.api.v1.Feed/LiveFeed"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(protoRequest))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	pkgCommon.SetHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	liveFeed, err := pkgCommon.DecodeGRPCMessage(body, &live_feed.LiveFeedResponse{})
	if err != nil {
		log.Fatalf("Failed to decode gRPC frame: %v", err)
	}

	return liveFeed, nil
}

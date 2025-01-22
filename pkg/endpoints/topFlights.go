package endpoints

import (
	"bytes"
	"fmt"
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/pkg/common"
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/protos/top_flights"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"log"
	"net/http"
)

func GetTopFlights() (string, error) {
	topFlightsRequest := &top_flights.TopFlightsRequest{
		Top: 10,
	}

	protoData, err := common.EncodeGRPCMessage(topFlightsRequest)
	if err != nil {
		return "", fmt.Errorf("failed to encode top_flights.Request object: %w", err)
	}

	url := "https://data-feed.flightradar24.com/fr24.feed.api.v1.Feed/TopFlights"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(protoData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	common.SetHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	topFlights, err := common.DecodeGRPCMessage(body, &top_flights.TopFlightsResponse{})
	if err != nil {
		log.Fatalf("Failed to decode gRPC frame: %v", err)
	}

	jsonData, err := protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}.Marshal(topFlights)

	if err != nil {
		log.Fatalf("Failed to marshal to JSON: %v", err)
	}

	return string(jsonData), nil
}

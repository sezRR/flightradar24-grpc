package main

import (
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/pkg/endpoints"
	"github.com/sezrr/go-playground/examples/protobuf-flightradar24/pkg/http"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	http.Init()

	topFlights, err := endpoints.GetTopFlights()
	if err != nil {
		log.Fatalf("Failed to get top flights: %v", err)
	}

	log.Println("Top flights:", topFlights)

	_, err = endpoints.GetLiveFlights()
	if err != nil {
		return
	}

	wg.Wait()
}

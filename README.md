# FlightRadar24 RPC Request Scraper
It is a basic repository to be able to use FlightRadar24 gRPC request responses within API with using Go!

## Todos
- [X] Implement `LiveFeed` method
- [ ] Implement `NearestFlights` method
- [ ] Implement `LiveFlightsStatus` method
- [ ] Implement `FollowFlight` method
- [ ] Add configuration file for endpoints

## Methods
- [X] `LiveFeed` https://data-feed.flightradar24.com/fr24.feed.api.v1.Feed/LiveFeed
- `Playback`
- `NearestFlights`
- `LiveFlightsStatus`
- `FollowFlight`
- [X] `TopFlights` https://data-feed.flightradar24.com/fr24.feed.api.v1.Feed/TopFlights
- `LiveTrail`

## Top Flights
1. We need payload for post request to the address https://data-feed.flightradar24.com/fr24.feed.api.v1.Feed/TopFlights
   - That payload contains numerical constants, specified in [top_flights_input.proto](https://github.com/sezRR/go-playground/tree/main/examples/protobuf-flightradar24/playground/top_flights_input.proto)
   - You can specify 10 to the field `top`, however, if you try to assign a value which is more than 10, API will return empty data. On the other hand, you can assign a positive number which is lower than 10.
2. After you have sent to request to the specified url in the above, you have to clear frame header and trailer, first 5 bytes and grpc-status field need to be removed.
3. You can marshal your protobuf!

## Usage
To generate .pb.go files, you can use the following command:
```bash
make -f .\scripts\protobuf.mk
```

## Special Thanks
https://github.com/cathaypacific8747/fr24
https://protobuf-decoder.netlify.app/
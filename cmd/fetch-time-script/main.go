package main

import (
	"flag"
	"fmt"
	"github.com/amoscookeh/the-daily-drive/internal/api/maps"
	"github.com/amoscookeh/the-daily-drive/internal/utils"
	"os"
	"strings"
	"time"
)

type Interval int32 // Minutes

type InputParams struct {
	userId       string
	fromPlace    maps.PlaceId
	toPlaces     []maps.PlaceId
	commuteStart time.Time
	commuteEnd   time.Time
	interval     Interval // minutes
	routeId      string
	envPath      *string
}

func main() {
	fmt.Println("Running commute time data fetcher")

	inputParams, err := parseFlags()
	if err != nil {
		panic(fmt.Sprintf("Error parsing input params: %v", err))
	}
	fmt.Println("Input params:", inputParams)

	utils.SetupEnv(inputParams.envPath)
}

func parseFlags() (*InputParams, error) {
	userID := flag.String("user", "", "User ID")
	fromPlace := flag.String("from", "", "From place ID")
	toPlaces := flag.String("to", "", "Comma-separated list of to place IDs")
	commuteStart := flag.String("start", "", "Commute start time (format: 2006-01-02T15:04:05)")
	commuteEnd := flag.String("end", "", "Commute end time (format: 2006-01-02T15:04:05)")
	interval := flag.Int("interval", 15, "Interval in minutes")
	routeId := flag.String("route", "", "Route ID")
	envPath := flag.String("env", ".env", "Path to env file")

	flag.Parse()

	numFlags := flag.NArg()
	if numFlags < 7 {
		return nil, fmt.Errorf("Expected at least 7 flags (-user -from -to -start -end - interval -route), received: %d", numFlags)
	}

	// Parse toPlaces
	toPlacesList := strings.Split(*toPlaces, ",")
	toPlacesIDs := make([]maps.PlaceId, len(toPlacesList))
	for i, place := range toPlacesList {
		toPlacesIDs[i] = maps.PlaceId(strings.TrimSpace(place))
	}

	// Parse times
	start, err := time.Parse("2006-01-02T15:04:05", *commuteStart)
	if err != nil {
		fmt.Println("Error parsing commute start time:", err)
		os.Exit(1)
	}
	end, err := time.Parse("2006-01-02T15:04:05", *commuteEnd)
	if err != nil {
		fmt.Println("Error parsing commute end time:", err)
		os.Exit(1)
	}

	inputParams := InputParams{
		userId:       *userID,
		fromPlace:    maps.PlaceId(*fromPlace),
		toPlaces:     toPlacesIDs,
		commuteStart: start,
		commuteEnd:   end,
		interval:     Interval(*interval),
		routeId:      *routeId,
		envPath:      envPath,
	}

	return &inputParams, nil
}

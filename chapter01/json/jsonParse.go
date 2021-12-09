package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`

	Data struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

type station struct {
	ID                     string `json:"station_id"`
	NumBikesAvailable      int    `json:"num_bikes_available"`
	NumEbikesAvailable     int    `json:"num_ebikes_available"`
	LastReported           int    `json:"last_reported"`
	NumDocksAvailable      int    `json:"num_docks_available"`
	EightdHasAvailableKeys bool   `json:"eightd_has_available_keys"`
	StationStatus          string `json:"station_status"`
	IsRenting              int    `json:"is_renting"`
	LegacyID               string `json:"legacy_id"`
	IsInstalled            int    `json:"is_installed"`
	NumDocksDisabled       int    `json:"num_docks_disabled"`
	NumBikesDisabled       int    `json:"num_bikes_disabled"`
	IsReturning            int    `json:"is_returning"`
}

func main() {
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd stationData
	err = json.Unmarshal(body, &sd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", sd.Data.Stations[0])

	outputData, err := json.Marshal(sd)
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile("citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}

}

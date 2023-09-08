package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Pokeapi struct {
	client http.Client
}

func (pokeapi *Pokeapi) GetLocationAreas(nextUrl *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	url := baseUrl + endpoint

	if nextUrl != nil {
		url = *nextUrl
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	response, err := pokeapi.client.Do(request)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResponse, nil
}

func createFetcher() Pokeapi {
	return Pokeapi{
		client: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

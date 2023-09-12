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

func (pokeapi *Pokeapi) GetLocationAreas(nextUrl *string, cache *Cache) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	url := baseUrl + endpoint

	if nextUrl != nil {
		url = *nextUrl
	}

	data, ok := cache.Get(url)
	if ok {
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResponse, nil
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

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	cache.Add(url, data)

	return locationAreasResponse, nil
}

func (pokeapi *Pokeapi) GetLocationArea(name string, cache *Cache) (LocationAreaResponse, error) {
	endpoint := fmt.Sprintf("/location-area/%s", name)
	url := baseUrl + endpoint

	data, ok := cache.Get(url)
	if ok {
		locationAreaResponse := LocationAreaResponse{}
		err := json.Unmarshal(data, &locationAreaResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationAreaResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	response, err := pokeapi.client.Do(request)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	cache.Add(url, data)

	return locationAreaResponse, nil
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
type LocationAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

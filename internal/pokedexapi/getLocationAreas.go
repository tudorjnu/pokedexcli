package pokedexapi

import (
	"encoding/json"
	"net/http"
)

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (locationAreasResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationAreasResponse{}, err
	}
	defer res.Body.Close()

	var resParsed locationAreasResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resParsed)
	if err != nil {
		return locationAreasResponse{}, err
	}

	return resParsed, nil
}

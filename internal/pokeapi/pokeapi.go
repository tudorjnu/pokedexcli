package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/tudorjnu/pokedexcli/internal/pokecache"
)

type PokeAPI struct {
	cache pokecache.Cache
}

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (papi *PokeAPI) GetLocationAreas(url string) (locationAreasResponse, error) {
	v, ok := papi.cache.Get(url)
	var body []byte

	if ok {
		body = v
	}

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locationAreasResponse{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return locationAreasResponse{}, err
		}

		papi.cache.Add(url, body)
	}

	var resParsed locationAreasResponse
	err := json.Unmarshal(body, &resParsed)
	if err != nil {
		return locationAreasResponse{}, err
	}

	return resParsed, nil

}

func NewPokeApi(cache pokecache.Cache) PokeAPI {
	return PokeAPI{
		cache: cache,
	}
}

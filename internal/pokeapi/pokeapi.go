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

func NewPokeApi(cache pokecache.Cache) PokeAPI {
	return PokeAPI{
		cache: cache,
	}
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
	body, ok := papi.cache.Get(url)

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

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type locationAreasPokemonsResponse struct {
	PokemonEncounters []struct {
		Pokemon        `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (papi *PokeAPI) GetLocationAreaPokemons(url string) ([]Pokemon, error) {
	body, ok := papi.cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []Pokemon{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []Pokemon{}, nil
		}

		papi.cache.Add(url, body)
	}

	var resJson locationAreasPokemonsResponse
	err := json.Unmarshal(body, &resJson)

	if err != nil {
		return []Pokemon{}, err
	}

	pokeList := make([]Pokemon, 1)

	for _, pokeEncounter := range resJson.PokemonEncounters {
		pokeList = append(pokeList, pokeEncounter.Pokemon)
	}

	return pokeList, nil
}

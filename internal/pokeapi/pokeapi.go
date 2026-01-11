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
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
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

type pokemonResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}

func (papi *PokeAPI) GetPokemon(pokemonString string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonString
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var pokemon pokemonResponse
	err = decoder.Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return Pokemon{
		Name:           pokemon.Name,
		URL:            url,
		BaseExperience: pokemon.BaseExperience,
	}, nil
}

func (papi *PokeAPI) InspectPokemon(pokemonString string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/inspect/" + pokemonString
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var pokemon pokemonResponse
	err = decoder.Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return Pokemon{
		Name:           pokemon.Name,
		URL:            url,
		BaseExperience: pokemon.BaseExperience,
		Height:         pokemon.Height,
		Weight:         pokemon.Weight,
		Stats:          pokemon.Stats,
	}, nil
}

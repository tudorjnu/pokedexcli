package config

import "github.com/tudorjnu/pokedexcli/internal/pokeapi"

type Config struct {
	Previous string
	Next     string
	PokeApi  pokeapi.PokeAPI
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/tudorjnu/pokedexcli/internal/config"
	"github.com/tudorjnu/pokedexcli/internal/pokeapi"
	"github.com/tudorjnu/pokedexcli/internal/pokecache"
	"github.com/tudorjnu/pokedexcli/internal/repl"
)

func main() {
	duration, err := time.ParseDuration("10s")
	if err != nil {
		fmt.Printf("Invalid duration!")
	}

	cache := pokecache.NewCache(duration)
	pokeapi := pokeapi.NewPokeApi(cache)

	config := config.Config{
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PokeApi:  pokeapi,
	}

	scanner := bufio.NewScanner(os.Stdin)

	commandMap := repl.InitMap()

	for {
		fmt.Fprintf(os.Stdout, "Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
		}

		inputText := scanner.Text()
		parsedInput := repl.CleanInput(inputText)

		if parsedInput == nil {
			fmt.Printf("Unknown command")
			continue
		}

		command := parsedInput[0]
		args := parsedInput[1:]
		_, ok := commandMap[command]
		if !ok {
			fmt.Fprintln(os.Stdout, "Unknown command")
			continue
		}

		fmt.Println()
		err := commandMap[command].Callback(&config, args)
		if err != nil {
			fmt.Println(err)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tudorjnu/pokedexcli/internal/config"
	"github.com/tudorjnu/pokedexcli/internal/repl"
)

func main() {
	fmt.Printf("Your Program was initiated successfully!\n")
	config := config.Config{
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	}
	scanner := bufio.NewScanner(os.Stdin)

	commandMap := repl.InitMap()

	for {
		fmt.Fprintf(os.Stdout, "Pokedex > ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
		}

		command := scanner.Text()

		_, ok := commandMap[command]
		if !ok {
			fmt.Fprintln(os.Stdout, "Unknown command")
			continue
		}

		commandMap[command].Callback(&config)

		// fmt.Printf("Your command was: %v\n", cleanText[0])
	}

}

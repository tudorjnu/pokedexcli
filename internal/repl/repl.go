package repl

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/tudorjnu/pokedexcli/internal/config"
)

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit(c *config.Config, args []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *config.Config, args []string) error
}

func InitMap() map[string]CliCommand {
	m := make(map[string]CliCommand)

	m["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the pokedex",
		Callback:    commandExit,
	}

	m["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback: func(c *config.Config, args []string) error {
			fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
			for _, v := range m {
				fmt.Printf("%s: %s\n", v.Name, v.Description)
			}
			return nil
		},
	}

	m["map"] = CliCommand{
		Name:        "map",
		Description: "Displays the names of 20 location areas in the Pokemon world.",
		Callback:    commandMap,
	}

	m["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "Displays the names of 20 location areas in the Pokemon world (goes back).",
		Callback:    commandMapB,
	}

	m["explore"] = CliCommand{
		Name:        "explore",
		Description: "Find pokemons inside an area",
		Callback:    commandExplore,
	}
	return m
}

func commandMap(c *config.Config, args []string) error {
	if c.Next == "" {
		fmt.Printf("you're on the last page\n")
	}

	r, err := c.PokeApi.GetLocationAreas(c.Next)
	if err != nil {
		return err
	}

	for _, v := range r.Results {
		fmt.Println(v.Name)
	}

	if r.Next == nil {
		c.Next = ""
	} else {
		c.Next = *r.Next
	}

	if r.Previous == nil {
		c.Previous = ""
	} else {
		c.Previous = *r.Previous
	}

	return nil
}

func commandMapB(c *config.Config, args []string) error {
	if c.Previous == "" {
		fmt.Printf("you're on the first page\n")
	}

	r, err := c.PokeApi.GetLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	for _, location := range r.Results {
		fmt.Println(location.Name)
	}

	if r.Next == nil {
		c.Next = ""
	} else {
		c.Next = *r.Next
	}

	if r.Previous == nil {
		c.Previous = ""
	} else {
		c.Previous = *r.Previous
	}

	return nil
}

func commandExplore(c *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Missing argument for command explore")
	}
	location := args[0]

	url := "https://pokeapi.co/api/v2/location-area/" + location
	pokemons, err := c.PokeApi.GetLocationAreaPokemons(url)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemons {
		fmt.Println(pokemon.Name)
	}
	return nil
}

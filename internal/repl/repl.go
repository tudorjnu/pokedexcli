package repl

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

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

	m["catch"] = CliCommand{
		Name:        "catch",
		Description: "Catch a pokemon. Syntax 'catch <pokemon>'",
		Callback:    commandCatch,
	}

	m["inspect"] = CliCommand{
		Name:        "inspect",
		Description: "Inspect a caught pokemon. Syntax 'inspect <pokemon>'",
		Callback:    commandInspect,
	}

	m["pokedex"] = CliCommand{
		Name:        "pokedex",
		Description: "List all caught pokemons",
		Callback:    commandPokedex,
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

func commandCatch(c *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Missing argument for command catch (pokemon)")
	}

	_, ok := c.PokeDex[args[0]]
	if ok {
		fmt.Println("You already have the pokemon")
		return nil
	}

	pokemon, err := c.PokeApi.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	r := rand.New(rand.NewSource(time.Now().Unix()))

	maxExperience := 635.0
	baseExp := float64(pokemon.BaseExperience)

	normProb := maxExperience / (maxExperience + baseExp)
	probability := math.Min(0.50, normProb)

	// fmt.Printf("Probability: %v vs %v\n", normProb, probability)

	if r.Float64() >= probability {
		fmt.Println("Sorry, he escaped")
		return nil
	}

	fmt.Println("You caught him")

	c.PokeDex[args[0]] = pokemon
	return nil
}

func commandInspect(c *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("Missing Argument")
	}

	pokemon, ok := c.PokeDex[args[0]]
	if !ok {
		return errors.New("You have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}

func commandPokedex(c *config.Config, args []string) error {
	fmt.Println("Your Pokedex:")
	for k, _ := range c.PokeDex {
		fmt.Printf("- %s\n", k)
	}
	fmt.Printf("\n")

	return nil
}

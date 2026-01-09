package repl

import (
	"fmt"
	"github.com/tudorjnu/pokedexcli/internal/config"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Split(strings.TrimSpace(lower), " ")

	cleanSplit := []string{}
	for _, word := range split {
		if word != "" {
			cleanSplit = append(cleanSplit, word)
		}
	}

	return cleanSplit
}

func commandExit(*config.Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config.Config) error
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
		Callback: func(*config.Config) error {
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
		Callback:    CommandMap,
	}

	m["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "Displays the names of 20 location areas in the Pokemon world (goes back).",
		Callback:    commandMapB,
	}
	return m
}

func CommandMap(c *config.Config) error {
	fmt.Println("You called the map command, wow!")

	if c.Next == "" {
		fmt.Printf("you're on the last page\n")
	}

	r, err := c.PokeApi.GetLocationAreas(c.Next)
	if err != nil {
		return err
	}

	for _, v := range r.Results {
		fmt.Printf("%s\n", v.Name)
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

func commandMapB(c *config.Config) error {
	if c.Previous == "" {
		fmt.Printf("you're on the first page\n")
	}

	r, err := c.PokeApi.GetLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	for _, location := range r.Results {
		fmt.Printf("%s\n", location.Name)
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

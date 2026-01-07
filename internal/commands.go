package main

import (
	"fmt"
	"github.com/tudorjnu/pokedexcli/internal/config"
	"github.com/tudorjnu/pokedexcli/internal/pokedexapi"
	"os"
)

func commandExit(*config.Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type CliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
}

func InitMap() map[string]CliCommand {
	m := make(map[string]CliCommand)

	m["exit"] = CliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	}

	m["help"] = CliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func(*config.Config) error {
			fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
			for _, v := range m {
				fmt.Printf("%s: %s\n", v.name, v.description)
			}
			return nil
		},
	}

	m["map"] = CliCommand{
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world.",
		callback:    commandMap,
	}

	m["mapb"] = CliCommand{
		name:        "mapb",
		description: "Displays the names of 20 location areas in the Pokemon world (goes back).",
		callback:    commandMapB,
	}
	return m
}

func commandMap(c *config.Config) error {
	fmt.Println("You called the map command, wow!")

	if c.Next == "" {
		fmt.Printf("you're on the last page\n")
	}

	r, err := pokedexapi.GetLocationAreas(c.Next)
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

	r, err := pokedexapi.GetLocationAreas(c.Previous)
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

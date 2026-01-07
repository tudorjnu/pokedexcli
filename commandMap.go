package main

import (
	"fmt"

	"github.com/tudorjnu/pokedexcli/internal/config"
	"github.com/tudorjnu/pokedexcli/internal/pokedexapi"
)

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

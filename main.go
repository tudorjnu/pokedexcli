package main

import (
	"bufio"
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Fprintln(os.Stdout, "Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func initMap(m map[string]cliCommand) {
	m["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	}

	m["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
			for _, v := range m {
				fmt.Printf("%s: %s\n", v.name, v.description)
			}
			return nil
		},
	}

}

func main() {
	fmt.Printf("Your Program was initiated successfully!\n")
	scanner := bufio.NewScanner(os.Stdin)

	commandMap := make(map[string]cliCommand)
	initMap(commandMap)

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

		commandMap[command].callback()

		// fmt.Printf("Your command was: %v\n", cleanText[0])
	}

}

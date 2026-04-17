package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          string
	Previous      string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 location-areas from the pokeapi",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 location-areas from the pokeapi",
			callback:    commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	slice := strings.Fields(text)
	for i, v := range slice {
		slice[i] = strings.ToLower(v)

	}

	return slice
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		token := cleanInput(scanner.Text())[0]

		if _, ok := getCommands()[token]; ok {
			err := getCommands()[token].callback(cfg)
			if err != nil {
				fmt.Printf("Error found: %v\n", err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}

	}
}

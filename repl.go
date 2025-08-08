package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commandRegistry := map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays next 20 locations.",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations.",
			callback:    commandMapBackward,
		},
		"help": {
			name:        "help",
			description: "Prints the list of available commands.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex.",
			callback:    commandExit,
		},
	}
	cfg := config{
		nextLocURL: "https://pokeapi.co/api/v2/location-area/",
		prevLocURL: "",
	}

	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		command := reader.Text()
		words := cleanInput(command)
		if len(words) == 0 {
			continue
		}

		cmd, exists := commandRegistry[words[0]]
		if !exists {
			fmt.Printf("Command '%s' not found!\n", words[0])
			continue
		}

		err := cmd.callback(&cfg, commandRegistry)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

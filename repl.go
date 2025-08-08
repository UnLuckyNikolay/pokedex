package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/UnLuckyNikolay/pokedex/internal/pokeapi"
	"github.com/UnLuckyNikolay/pokedex/internal/pokecache"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := config{
		httpClient: pokeapi.NewClient(5 * time.Second),
		cache:      pokecache.NewCache(1 * time.Hour),
		nextLocURL: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		prevLocURL: "",
	}
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

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Write 'help' to see available commands.")
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

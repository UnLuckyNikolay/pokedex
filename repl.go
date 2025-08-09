package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chzyer/readline"

	"github.com/UnLuckyNikolay/pokedex/internal/pokeapi"
	"github.com/UnLuckyNikolay/pokedex/internal/pokecache"
)

func startRepl() {
	readerTemp, err := readline.New("\033[31mPokedex > \033[0m")
	if err != nil {
		log.Fatal(err)
	}

	var command string

	cfg := config{
		httpClient: pokeapi.NewClient(5 * time.Second),
		cache:      pokecache.NewCache(1 * time.Hour),
		baseURL:    "https://pokeapi.co/api/v2/",
		reader:     readerTemp,

		locPage: 0,
		locMax:  0,

		pokedex: map[string]pokeapi.Pokemon{},
	}

	commandRegistry := map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays 20 next locations.",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 previous locations.",
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
		"explore": {
			name:        "explore <id/location>",
			description: "Moves you into the specified location. Leave empty to reexplore current location.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Tries to catch the specified pokemon. You need to be in the same location as them.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect the pokemon that was previously caught.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the list of the caught pokemon.",
			callback:    commandPokedex,
		},
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Write 'help' to see available commands.")
	fmt.Println()

	//Main loop
	for {
		/*if cfg.locCurrent != nil {
			fmt.Print(getLocationName(*cfg.locCurrent) + " > ")
		}*/

		//Command input
		command, err = cfg.reader.Readline()
		words := cleanInput(command)
		if len(words) == 0 {
			continue
		}

		cmd, exists := commandRegistry[words[0]]
		if !exists {
			fmt.Printf("Command '%s' not found!\n", words[0])
			fmt.Println()
			continue
		}

		//Callback
		var err error
		if len(words) >= 1 {
			err = cmd.callback(&cfg, commandRegistry, words[1:])
		} else {
			err = cmd.callback(&cfg, commandRegistry, []string{})
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println()
	}
}

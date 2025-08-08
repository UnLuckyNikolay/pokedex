package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, commandRegistry map[string]cliCommand) error
}

func commandHelp(cfg *config, commandRegistry map[string]cliCommand) error {
	fmt.Println("List of available commands:")
	for _, cmd := range commandRegistry {
		fmt.Printf(" > %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config, commandRegistry map[string]cliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapForward(cfg *config, commandRegistry map[string]cliCommand) error {
	if cfg.nextLocURL == "" {
		return fmt.Errorf("At the end of the list, use 'mapb' to go backward.")
	}

	data, err := cfg.httpClient.GetLocations(cfg.nextLocURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.nextLocURL = data.Next
	cfg.prevLocURL = data.Previous
	for _, loc := range data.Results {
		fmt.Println(" > " + loc.Name)
	}

	return nil

}

func commandMapBackward(cfg *config, commandRegistry map[string]cliCommand) error {
	if cfg.prevLocURL == "" {
		return fmt.Errorf("At the start of the list, use 'map' to go forward.")
	}

	data, err := cfg.httpClient.GetLocations(cfg.prevLocURL, cfg.cache)
	if err != nil {
		return err
	}

	cfg.nextLocURL = data.Next
	cfg.prevLocURL = data.Previous
	for _, loc := range data.Results {
		fmt.Println(" > " + loc.Name)
	}

	return nil
}

package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfgPtr *config, commandRegistry map[string]cliCommand) error
}

func commandHelp(cfgPtr *config, commandRegistry map[string]cliCommand) error {
	fmt.Println("List of available commands:")
	for _, cmd := range commandRegistry {
		fmt.Printf(" > %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfgPtr *config, commandRegistry map[string]cliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapForward(cfgPtr *config, commandRegistry map[string]cliCommand) error {
	if cfgPtr.nextLocURL == "" {
		return fmt.Errorf("At the end of the list, use 'mapb' to go backward.")
	}

	data, err := cfgPtr.httpClient.GetLocations(cfgPtr.nextLocURL)
	if err != nil {
		return err
	}

	cfgPtr.nextLocURL = data.Next
	cfgPtr.prevLocURL = data.Previous
	for _, loc := range data.Results {
		fmt.Println(" > " + loc.Name)
	}

	return nil

}

func commandMapBackward(cfgPtr *config, commandRegistry map[string]cliCommand) error {
	if cfgPtr.prevLocURL == "" {
		return fmt.Errorf("At the start of the list, use 'map' to go forward.")
	}

	data, err := cfgPtr.httpClient.GetLocations(cfgPtr.prevLocURL)
	if err != nil {
		return err
	}

	cfgPtr.nextLocURL = data.Next
	cfgPtr.prevLocURL = data.Previous
	for _, loc := range data.Results {
		fmt.Println("  " + loc.Name)
	}

	return nil
}

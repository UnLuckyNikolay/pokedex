package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, commandRegistry map[string]cliCommand, args []string) error
}

func commandHelp(cfg *config, commandRegistry map[string]cliCommand, args []string) error {
	fmt.Println("List of available commands:")
	for _, cmd := range commandRegistry {
		fmt.Printf(" > %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(cfg *config, commandRegistry map[string]cliCommand, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapForward(cfg *config, commandRegistry map[string]cliCommand, args []string) error {
	if cfg.locMax != 0 && cfg.locOffset >= cfg.locMax {
		return fmt.Errorf("At the end of the list, use 'mapb' to go backward.")
	}

	//Building url
	url := fmt.Sprintf("%slocation-area/?limit=20&offset=%d", cfg.baseURL, cfg.locOffset)

	//Getting data
	data, err := cfg.httpClient.GetLocations(url, cfg.cache)
	if err != nil {
		return err
	}

	//Printing the list
	lastLocNum := cfg.locOffset + 20
	if data.Count < lastLocNum {
		lastLocNum = data.Count
	}

	fmt.Printf("Showing locations %v-%v out of %v\n", cfg.locOffset+1, lastLocNum, data.Count)
	for _, loc := range data.Results {
		locIndex := strings.TrimPrefix(loc.URL, "https://pokeapi.co/api/v2/location-area/")
		locIndex = strings.TrimSuffix(locIndex, "/")

		fmt.Printf(" > %v - %v\n", locIndex, loc.Name)
	}

	//Updating config
	cfg.locMax = data.Count
	cfg.locOffset += 20
	if data.Count < cfg.locOffset {
		cfg.locOffset = data.Count
	}

	return nil
}

func commandMapBackward(cfg *config, commandRegistry map[string]cliCommand, args []string) error {
	if cfg.locOffset == 0 {
		return fmt.Errorf("At the start of the list, use 'map' to go forward.")
	}

	//Building url
	limit := 20
	offset := cfg.locOffset - 20
	if offset < 0 {
		limit += offset
		offset = 0
	}
	url := fmt.Sprintf("%slocation-area/?limit=%d&offset=%d", cfg.baseURL, limit, offset)

	//Getting data
	data, err := cfg.httpClient.GetLocations(url, cfg.cache)
	if err != nil {
		return err
	}

	//Printing the list
	lastLocNum := cfg.locOffset + 1 + limit
	if data.Count < lastLocNum {
		lastLocNum = data.Count
	}

	fmt.Printf("Showing locations %v-%v out of %v\n", offset+1, offset+limit, data.Count)
	for _, loc := range data.Results {
		locIndex := strings.TrimPrefix(loc.URL, "https://pokeapi.co/api/v2/location-area/")
		locIndex = strings.TrimSuffix(locIndex, "/")

		fmt.Printf(" > %v - %v\n", locIndex, loc.Name)
	}

	//Updating config
	cfg.locMax = data.Count
	cfg.locOffset -= limit

	return nil
}

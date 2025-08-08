package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfgPtr *config, commandRegistry map[string]cliCommand) error
}

func init() {
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

type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMapForward(cfgPtr *config, commandRegistry map[string]cliCommand) error {
	if cfgPtr.next == "" {
		return fmt.Errorf("At the end of the list, use 'mapb' to go backward.")
	}

	err := commandMap(cfgPtr, cfgPtr.next)

	return err

}

func commandMapBackward(cfgPtr *config, commandRegistry map[string]cliCommand) error {
	if cfgPtr.previous == "" {
		return fmt.Errorf("At the start of the list, use 'map' to go forward.")
	}

	err := commandMap(cfgPtr, cfgPtr.previous)

	return err
}

func commandMap(cfgPtr *config, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failure making a GET request: %v", err)
	}
	defer res.Body.Close()

	data := location{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("Failure decoding JSON: %v", err)
	}

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	cfgPtr.next = data.Next
	cfgPtr.previous = data.Previous

	return nil
}

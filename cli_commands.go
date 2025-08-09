package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/UnLuckyNikolay/pokedex/internal/pokeapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	//Building url
	limit := 20
	offset := cfg.locPage * limit
	url := fmt.Sprintf("%slocation-area/?limit=%d&offset=%d", cfg.baseURL, limit, offset)

	//Getting data
	data, err := cfg.httpClient.GetLocationAreaList(url, cfg.cache)
	if err != nil {
		return err
	}

	//Printing the list
	lastLocNum := offset + limit
	if data.Count < lastLocNum {
		lastLocNum = data.Count
	}

	fmt.Printf("Showing locations %v-%v out of %v\n", offset+1, lastLocNum, data.Count)
	for _, loc := range data.Results {
		locIndex := strings.TrimPrefix(loc.URL, "https://pokeapi.co/api/v2/location-area/")
		locIndex = strings.TrimSuffix(locIndex, "/")

		fmt.Printf(" > %v - %v\n", locIndex, loc.Name)
	}

	//Updating config
	cfg.locMax = data.Count
	if data.Count > (cfg.locPage+1)*limit {
		cfg.locPage++
	}

	return nil
}

func commandMapBackward(cfg *config, commandRegistry map[string]cliCommand, args []string) error {
	//Building url
	limit := 20
	offset := (cfg.locPage - 2) * limit
	if offset < 0 {
		offset = 0
	}
	url := fmt.Sprintf("%slocation-area/?limit=%d&offset=%d", cfg.baseURL, limit, offset)

	//Getting data
	data, err := cfg.httpClient.GetLocationAreaList(url, cfg.cache)
	if err != nil {
		return err
	}

	//Printing the list
	lastLocNum := offset + limit
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
	cfg.locPage--
	if cfg.locPage < 1 {
		cfg.locPage = 1
	}

	return nil
}

func commandExplore(cfg *config, commandRegistry map[string]cliCommand, args []string) error {
	//Getting data
	var data pokeapi.LocationArea
	if len(args) == 0 && cfg.locCurrent != nil {
		data = *cfg.locCurrent
	} else if len(args) == 0 {
		return fmt.Errorf("You are not currently in a location. Write the name or the id of the destination.")
	} else {
		url := fmt.Sprintf("%slocation-area/%s", cfg.baseURL, args[0])

		var err error
		data, err = cfg.httpClient.GetLocationArea(url, cfg.cache)
		if err != nil {
			return fmt.Errorf("%v: %v", args[0], err)
		}
	}

	//Printing the list of pokemons
	fmt.Printf("Exploring the %s...\n", getLocationName(data))
	fmt.Printf("Encountered pokemon:\n")
	for _, pokemon := range data.PokemonEncounters {
		fmt.Printf(" > %s\n", pokemon.Pokemon.Name)
	}

	//Updating config
	cfg.locCurrent = &data

	return nil
}

func commandCatch(cfg *config, commandRegistry map[string]cliCommand, args []string) error {
	caser := cases.Title(language.English)
	nameTitle := caser.String(args[0])
	var pokemon pokeapi.Pokemon

	//Error checks
	_, caught := cfg.pokedex[args[0]]
	if caught {
		return fmt.Errorf("%s: pokemon already caught!", nameTitle)
	} else if cfg.locCurrent == nil {
		return fmt.Errorf("You are not currently in a location. Use command 'explore <id/location>' or 'map'.")
	} else if !checkIfPokemonIsPresent(*cfg.locCurrent, args[0]) {
		return fmt.Errorf("%s: pokemon not found in the current location!", nameTitle)
		//Getting data
	} else if cfg.lastWildPoke != nil && cfg.lastWildPoke.Name == args[0] {
		pokemon = *cfg.lastWildPoke
	} else {
		//Building url
		url := fmt.Sprintf("%spokemon/%s", cfg.baseURL, args[0])

		var err error
		pokemon, err = cfg.httpClient.GetPokemon(url, cfg.cache)
		if err != nil {
			return err
		}
	}

	//Rolling for catch
	fmt.Printf("Throwing a Pokeball at %s...\n", nameTitle)
	roll := rand.IntN(10)
	roll = roll * roll * roll //Base experience range - 36 (Sunkern) to 635 (Blissey)

	//Updating config
	if roll >= pokemon.BaseExperience {
		fmt.Printf("Successfully caught %s!\n", nameTitle)

		cfg.pokedex[args[0]] = pokemon
		cfg.lastWildPoke = nil
	} else {
		fmt.Printf("%s escaped!\n", nameTitle)

		cfg.lastWildPoke = &pokemon //Saved for recatching
	}

	return nil
}

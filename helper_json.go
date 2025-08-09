package main

import "github.com/UnLuckyNikolay/pokedex/internal/pokeapi"

func getLocationName(loc pokeapi.LocationArea) string {
	for _, lang := range loc.Names {
		if lang.Language.Name == "en" {
			return lang.Name
		}
	}

	return loc.Name
}

package main

import (
	"github.com/UnLuckyNikolay/pokedex/internal/pokeapi"
	"github.com/UnLuckyNikolay/pokedex/internal/pokecache"
	"github.com/chzyer/readline"
)

type config struct {
	httpClient pokeapi.Client
	cache      *pokecache.Cache
	baseURL    string
	reader     *readline.Instance

	locPage    int
	locMax     int
	locCurrent *pokeapi.LocationArea

	pokedex      map[string]pokeapi.Pokemon
	lastWildPoke *pokeapi.Pokemon
}

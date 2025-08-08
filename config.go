package main

import (
	"github.com/UnLuckyNikolay/pokedex/internal/pokeapi"
	"github.com/UnLuckyNikolay/pokedex/internal/pokecache"
)

type config struct {
	httpClient pokeapi.Client
	cache      *pokecache.Cache
	nextLocURL string
	prevLocURL string
}

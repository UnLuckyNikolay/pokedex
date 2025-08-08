package main

import (
	"github.com/UnLuckyNikolay/pokedex/pokeapi"
)

type config struct {
	httpClient pokeapi.Client
	nextLocURL string
	prevLocURL string
}

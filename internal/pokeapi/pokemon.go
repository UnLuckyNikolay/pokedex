package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/UnLuckyNikolay/pokedex/internal/pokecache"
)

func (c *Client) GetPokemon(url string, cache *pokecache.Cache) (Pokemon, error) {
	var data []byte

	data, exists := cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		cache.Add(url, data)
	}

	pokemon := Pokemon{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

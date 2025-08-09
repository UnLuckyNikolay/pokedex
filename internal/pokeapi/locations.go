package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/UnLuckyNikolay/pokedex/internal/pokecache"
)

func (c *Client) GetLocations(url string, cache *pokecache.Cache) (Locations, error) {
	var data []byte

	data, exists := cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Locations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Locations{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Locations{}, err
		}

		cache.Add(url, data)
	}

	locs := Locations{}
	err := json.Unmarshal(data, &locs)
	if err != nil {
		return Locations{}, err
	}

	return locs, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/UnLuckyNikolay/pokedex/internal/pokecache"
)

func (c *Client) GetLocationAreaList(url string, cache *pokecache.Cache) (LocationAreaList, error) {
	var data []byte

	data, exists := cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaList{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaList{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaList{}, err
		}

		cache.Add(url, data)
	}

	locs := LocationAreaList{}
	err := json.Unmarshal(data, &locs)
	if err != nil {
		return LocationAreaList{}, err
	}

	return locs, nil
}

func (c *Client) GetLocationArea(url string, cache *pokecache.Cache) (LocationArea, error) {
	var data []byte

	data, exists := cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		if res.StatusCode == 404 {
			return LocationArea{}, fmt.Errorf("Invalid location name or id")
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}

		cache.Add(url, data)
	}

	loc := LocationArea{}
	err := json.Unmarshal(data, &loc)
	if err != nil {
		return LocationArea{}, err
	}

	return loc, nil
}

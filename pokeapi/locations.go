package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocations(url string) (Locations, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	data := Locations{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return Locations{}, err
	}

	return data, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocations(url string) (Locations, error) {
	res, err := http.Get(url)
	if err != nil {
		return Locations{}, fmt.Errorf("Failure making a GET request: %v", err)
	}
	defer res.Body.Close()

	data := Locations{}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return Locations{}, fmt.Errorf("Failure decoding JSON: %v", err)
	}

	return data, nil
}

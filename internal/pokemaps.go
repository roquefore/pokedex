package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokedexConfig struct {
	NextUrl string
	PrevUrl string
}

type LocationArea struct {
	Name 	string	`json:"name"`
	Url 	string	`json:"url"`
}

type PokedexPage struct {
	Count 		int				`json:"count"`
	Next  		string			`json:"next"`
	Previous 	string			`json:"previous"`
	Results 	[]LocationArea	`json:"results"`
}

var MapsConfig = pokedexConfig{
	NextUrl: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	PrevUrl: "",
}

var pokemapsCache = InitCache(10)

func MapsRequest(url string) error {
	var locBytes []byte
	
	if ok, cached := pokemapsCache.Get(url); ok {
		locBytes = cached.value
	} else {		
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("unable to get: %w", err)
		}
		defer res.Body.Close()
		
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("unable to read bytes: %w", err)
		}

		locBytes = bytes
		pokemapsCache.Add(url, locBytes)
	}

	var page PokedexPage

	if err := json.Unmarshal(locBytes, &page); err != nil {
			return fmt.Errorf("unable to unmarshal bytes: %w", err)
	}

	MapsConfig.NextUrl = page.Next
	MapsConfig.PrevUrl = page.Previous

	for _, el := range page.Results {
		fmt.Println(el.Name)
	}

	return nil
}
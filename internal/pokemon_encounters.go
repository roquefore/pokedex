package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemonEncounter struct {
	ID int 							`json:"id"`
	Name string 					`json:"name"`
	PokemonEncounters []struct{
		Pokemon struct {
			Name string 				`json:"name"`
			URL string 					`json:"url"`
		}
	} 								`json:"pokemon_encounters"`
}

var pokemonEncountersCache = InitCache(10)

func PokemonEncounterRequest(loc string) error {
	var pocBytes []byte

	if exists, entry := pokemonEncountersCache.Get(loc); exists {
		pocBytes = entry.value
	} else {	
		res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + loc)
		if err != nil {
			return fmt.Errorf("unable to get: %w", err)
		}
		defer res.Body.Close()

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("unable to read bytes: %w", err)
		}

		pocBytes = b
		pokemonEncountersCache.Add(loc, pocBytes)
	}

	var pokemon pokemonEncounter

	if err := json.Unmarshal(pocBytes, &pokemon); err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	fmt.Println("Exploring", loc, "...")
	for _, pok := range pokemon.PokemonEncounters {
		fmt.Println("-", pok.Pokemon.Name)
	}

	return nil
}
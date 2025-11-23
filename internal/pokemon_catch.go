package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	ID int			`json:"id"`
	Name string			`json:"name"`
	Species struct{
		Name string			`json:"name"`
		URL string			`json:"url"`
	}					`json:"species"`
}

type pokemonSpecies struct {
	ID int			`json:"id"`
	Name string			`json:"name"`
	CaptureRate int 	`json:"capture_rate"`
}

var Pokedex = map[string]Pokemon{}

var pokemonCache = InitCache(30)

func checkPokemonSpecies(name string) (pokemonSpecies, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon-species/" + name)
	if err != nil {
		return pokemonSpecies{}, fmt.Errorf("unabled to get pokemon species: %w", err)
	}

	var ps pokemonSpecies

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&ps); err != nil {
		return pokemonSpecies{}, fmt.Errorf("unabled to decode pokemon species: %w", err)
	}

	return ps, nil
}

func CatchPokemonRequest(name string) error {
	var pokBytes []byte

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if exists, b := pokemonCache.Get(name); exists {
		pokBytes = b.value
	} else {	
		res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
		if err != nil {
			return fmt.Errorf("unable to get pokemon: %w", err)
		}

		if b, err := io.ReadAll(res.Body); err != nil {
			return fmt.Errorf("unable to ready json body bytes: %w", err)
		} else {
			pokBytes = b
		}
	}

	var pokemonData Pokemon
	if err := json.Unmarshal(pokBytes, &pokemonData); err != nil {
		return fmt.Errorf("unable to unmarshal pokBytes: %w", err)
	}

	species, err := checkPokemonSpecies(pokemonData.Name)
	if err != nil {
		return fmt.Errorf("unable to check pokemon species: %w", err)
	}

	catchValue := rand.Intn(255)

	if catchValue >= species.CaptureRate {
		fmt.Printf("%s was caught!\n", name)
		Pokedex[name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", name)
	} 

	return nil
}

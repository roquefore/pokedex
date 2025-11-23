package internal

import "fmt"

func InspectPokemon(name string) error {
	if val, exists := Pokedex[name]; exists {
		fmt.Printf("Name: %s\n", val.Name)
        fmt.Printf("ID: %d\n", val.ID)
        fmt.Printf("Species: %s\n", val.Species.Name)
		return nil
	} else {
		fmt.Printf("Pokemon %s was not yet caught", name)
		return fmt.Errorf("Pokemon missing in Pokedex")
	}
}

func PrintPokedex() error {
	for key := range Pokedex  {
		fmt.Println("-", key)
	}

	return nil
}
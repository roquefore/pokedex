package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/roquefore/pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
}

var commandsRegistry = map[string]cliCommand {
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays next 20 location areas",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays prev 20 location areas",
		callback: commandMapb,
	},
	"explore": {
		name: "explore",
		description: "Displays popkemons in <location-area>",
		callback: commandExplore,
	},
	"catch": {
		name: "catch",
		description: "Performs attempt to catch pokemon by <pokemon-name>",
		callback: commandCatch,
	},
	"inspect": {
		name: "inspect",
		description: "Inspects pokemon from your Pokedex",
		callback: commandInspect,
	},
	"pokedex": {
		name: "pokedex",
		description: "List all pokemons in pokedex",
		callback: commandPokedex,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		lineParts := cleanInput(scanner.Text())
		if len(lineParts) == 0 {
			continue
		}

		cmd := lineParts[0]
		args := lineParts[1:]


		if handler, ok := commandsRegistry[cmd];  ok {
			if err := handler.callback(args); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	return strings.Fields(text)
}

func commandExit(_ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(_ []string) error {
	if internal.MapsConfig.NextUrl == "" {
		fmt.Println("Error: no more location areas")
		return fmt.Errorf("no more location areas")
	}

	return internal.MapsRequest(internal.MapsConfig.NextUrl)
}

func commandMapb(_ []string) error {
	if internal.MapsConfig.PrevUrl == "" {
		fmt.Println("you're on the first page")
		return fmt.Errorf("no previous page")
	}

	return internal.MapsRequest(internal.MapsConfig.PrevUrl)
}

func commandExplore(args []string) error {
	if len(args) == 0 {
		fmt.Println("Location area name expected")
		return nil
	} else {
		return internal.PokemonEncounterRequest(args[0])
	}
}

func commandCatch(args []string) error {
	if len(args) == 0  {
		fmt.Println("Pokemon name expected")
		return nil
	} else {
		return internal.CatchPokemonRequest(args[0])
	}
}

func commandInspect(args []string) error {
	if len(args) == 0  {
		fmt.Println("Pokemon name expected")
		return nil
	} else {
		return internal.InspectPokemon(args[0])
	}
}

func commandPokedex(_ []string) error {
	return internal.PrintPokedex()
}
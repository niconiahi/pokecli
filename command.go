package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Command struct {
	name        string
	description string
	execute     func(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error
}

type Pagination struct {
	next     *string
	previous *string
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Shows the help menu",
			execute:     executeHelp,
		},
		"exit": {
			name:        "exit",
			description: "Quits the program",
			execute:     executeExit,
		},
		"map": {
			name:        "map",
			description: "Get next page of areas",
			execute:     executeMap,
		},
		"explore": {
			name:        "explore {area_name}",
			description: "Lists the pokemons of an area",
			execute:     executeExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attemps to catch the selected pokemon",
			execute:     executeCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Check if the pokemon is caught",
			execute:     executeInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all captured pokemons",
			execute:     executePokedex,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of areas",
			execute:     executeMapb,
		},
	}
}

func executePokedex(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	fmt.Println("captured pokemons:")
	for _, pokemon := range pokemons {
		fmt.Printf("%s\n", pokemon.Name)
	}

	return nil
}

func executeInspect(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	name := args[0]
	pokemon, ok := pokemons[name]
	if !ok {
		return errors.New("you haven't caught this pokemon")
	}

	fmt.Printf("name: %s\n", pokemon.Name)
	fmt.Printf("height: %v\n", pokemon.Height)
	fmt.Printf("weight: %v\n", pokemon.Weight)
	fmt.Println("abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("  - %s\n", ability.Ability.Name)
	}

	return nil
}

func executeCatch(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	name := args[0]
	pokemon, err := pokeapi.GetPokemon(name, cache)
	if err != nil {
		return err
	}

	const threshold = 50
	random := rand.Intn(pokemon.BaseExperience)
	if random < threshold {
		fmt.Printf("failed to catch %s\n", name)

		return nil
	}

	fmt.Printf("%s was caught\n", name)
	pokemons[name] = pokemon

	return nil
}

func executeExplore(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area name provided")
	}

	name := args[0]
	location, err := pokeapi.GetLocationArea(name, cache)
	if err != nil {
		return err
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

func executeHelp(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	commands := getCommands()
	fmt.Println("Available commands:")

	for _, command := range commands {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}

	return nil
}

func executeMap(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	response, err := pokeapi.GetLocationAreas(pagination.next, cache)
	if err != nil {
		return errors.New("failed get the location areas")
	}

	pagination.next = response.Next
	pagination.previous = response.Previous

	for _, area := range response.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return nil
}

func executeMapb(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	if pagination.previous == nil {
		return errors.New("you are on the first page")
	}

	response, err := pokeapi.GetLocationAreas(pagination.previous, cache)
	if err != nil {
		return errors.New("failed get the location areas")
	}

	pagination.next = response.Next
	pagination.previous = response.Previous

	for _, area := range response.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	return nil
}

func executeExit(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, pokemons map[string]PokemonResponse, args ...string) error {
	os.Exit(0)

	return nil
}

package main

import (
	"errors"
	"fmt"
	"os"
)

type Command struct {
	name        string
	description string
	execute     func(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, args ...string) error
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
			name:        "explore",
			description: "Lists the pokemons of an area",
			execute:     executeExplore,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of areas",
			execute:     executeMapb,
		},
	}
}

func executeExplore(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area name provided")
	}

	name := args[0]
	response, err := pokeapi.GetLocationArea(name, cache)
	if err != nil {
		return err
	}

	for _, pokemon := range response.PokemonEncounters {
		fmt.Printf("- %s", pokemon.Pokemon.Name)
		fmt.Println()
	}

	return nil
}

func executeHelp(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, args ...string) error {
	commands := getCommands()
	fmt.Println("Available commands:")

	for _, command := range commands {
		fmt.Printf("- %s: %s", command.name, command.description)
		fmt.Println("")
	}

	return nil
}

func executeMap(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, args ...string) error {
	response, err := pokeapi.GetLocationAreas(pagination.next, cache)
	if err != nil {
		return errors.New("failed get the location areas")
	}

	pagination.next = response.Next
	pagination.previous = response.Previous

	for _, area := range response.Results {
		fmt.Printf("- %s", area.Name)
		fmt.Println()
	}

	return nil
}

func executeMapb(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, args ...string) error {
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
		fmt.Printf("- %s", area.Name)
		fmt.Println()
	}

	return nil
}

func executeExit(pokeapi *Pokeapi, pagination *Pagination, cache *Cache, args ...string) error {
	os.Exit(0)

	return nil
}

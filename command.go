package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Command struct {
	name        string
	description string
	execute     func(pokeapi *Pokeapi, pagination *Pagination) error
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
		"mapb": {
			name:        "mapb",
			description: "Get previous page of areas",
			execute:     executeMapb,
		},
	}
}

func executeHelp(pokeapi *Pokeapi, pagination *Pagination) error {
	commands := getCommands()
	fmt.Println("Available commands:")

	for _, command := range commands {
		fmt.Printf("- %s: %s", command.name, command.description)
		fmt.Println("")
	}

	return nil
}

func executeMap(pokeapi *Pokeapi, pagination *Pagination) error {
	response, err := pokeapi.GetLocationAreas(pagination.next)
	if err != nil {
		log.Fatal("failed get the location areas")
	}

	pagination.next = response.Next
	pagination.previous = response.Previous

	for _, area := range response.Results {
		fmt.Printf("- %s", area.Name)
		fmt.Println()
	}

	return nil
}

func executeMapb(pokeapi *Pokeapi, pagination *Pagination) error {
	if pagination.previous == nil {
		return errors.New("you are on the first page")
	}

	response, err := pokeapi.GetLocationAreas(pagination.previous)
	if err != nil {
		log.Fatal("failed get the location areas")
	}

	pagination.next = response.Next
	pagination.previous = response.Previous

	for _, area := range response.Results {
		fmt.Printf("- %s", area.Name)
		fmt.Println()
	}

	return nil
}

func executeExit(pokeapi *Pokeapi, pagination *Pagination) error {
	os.Exit(0)

	return nil
}

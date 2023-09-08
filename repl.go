package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start() {
	fetcher := createFetcher()
	pagination := Pagination{next: nil, previous: nil}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		words := getWords(text)

		if len(words) == 0 {
			fmt.Println("missing command: run help to get a list of available commands")
			continue
		}

		commands := getCommands()
		command, ok := commands[words[0]]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.execute(&fetcher, &pagination)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getWords(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

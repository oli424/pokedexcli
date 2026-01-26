package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	text = strings.ToLower(text)
	split_text := strings.Split(text, " ")
	return split_text
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

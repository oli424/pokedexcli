package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cmd_reg := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show help",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		cmd := words[0]

		if c, ok := cmd_reg[cmd]; ok {
			if err := c.callback(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/oli424/pokedexcli/internal/pokeapi"
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
		"map": {
			name:        "map",
			description: "List 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 areas",
			callback:    commandMapb,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		pokeClient: pokeapi.NewClient(5*time.Second, 5*time.Second),
	}
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		cmd := words[0]

		if c, ok := cmd_reg[cmd]; ok {
			if err := c.callback(cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

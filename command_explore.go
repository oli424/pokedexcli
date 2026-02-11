package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location name>")
	}
	name := args[0]
	loc, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", loc.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

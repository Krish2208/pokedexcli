package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}

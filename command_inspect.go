package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("expected 1 argument")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught a %s", pokemonName)
	}

	fmt.Printf("Inspecting %s...\n", pokemon.Name)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Order: %d\n", pokemon.Order)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, ty := range pokemon.Types {
		fmt.Printf(" - %s\n", ty.Type.Name)
	}
	return nil
}

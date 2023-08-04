package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("expected 1 argument")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s\n", locationArea.Name)
	for _, poke := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", poke.Pokemon.Name)
	}
	return nil
}

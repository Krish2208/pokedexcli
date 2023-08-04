package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("expected 1 argument")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing pokeball at %s...\n", pokemon.Name)
	const threshold = 50
	randomNumber := rand.Intn(pokemon.BaseExperience)
	if randomNumber > threshold {
		return fmt.Errorf("you failed to catch %s", pokemon.Name)
	}
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}

package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params ...string) error {
	pokemonName := params[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonData, _ := cfg.pokeapiClient.GetPokemon(pokemonName)

	BaseExperience := pokemonData.BaseExperience
	const difficulty = 50.0
	propability := difficulty / (float64(BaseExperience) - 39 + difficulty)
	randomNumber := rand.Float32()

	if randomNumber < float32(propability) {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

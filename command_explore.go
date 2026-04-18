package main

import "fmt"

func commandExplore(cfg *config, params ...string) error {
	urlPokemon := "https://pokeapi.co/api/v2/location-area/" + params[1]

	pokemons, err := cfg.pokeapiClient.ListPokemons(urlPokemon)
	if err != nil {
		return err
	}

	pokemon_encounters := pokemons.PokemonEncounters
	for _, pokemon := range pokemon_encounters {
		fmt.Printf("%s\n", pokemon.Pokemon.Name)
	}

	return nil
}

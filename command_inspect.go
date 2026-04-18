package main

import (
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
	if len(params) == 1 {
		for pokemon := range cfg.pokedex {
			fmt.Println(pokemon)
		}
		fmt.Println("Type `Inspect <pokemon_name>` to inspect the pokemon")
		return nil
	}

	pokemonName := params[1]
	if _, ok := cfg.pokedex[pokemonName]; !ok {
		fmt.Println("You have not catch that pokemon")
		return nil
	}
	pokemonData, _ := cfg.pokeapiClient.GetPokemon(pokemonName)
	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Printf("Stats: \n")
	fmt.Printf("	- hp : %d\n", pokemonData.Stats[0].BaseStat)
	fmt.Printf("	- attack : %d\n", pokemonData.Stats[1].BaseStat)
	fmt.Printf("	- defense : %d\n", pokemonData.Stats[2].BaseStat)
	fmt.Printf("	- special attack : %d\n", pokemonData.Stats[3].BaseStat)
	fmt.Printf("	- special defense : %d\n", pokemonData.Stats[4].BaseStat)
	fmt.Printf("	- speed : %d\n", pokemonData.Stats[5].BaseStat)
	fmt.Printf("Types: \n")
	for _, pokeType := range pokemonData.Types {
		fmt.Printf("	- %s\n", pokeType.Type.Name)
	}
	return nil
}

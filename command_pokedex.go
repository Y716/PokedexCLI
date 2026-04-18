package main

import "fmt"

func commandPokedex(cfg *config, params ...string) error {
	fmt.Println("Your pokedex:")
	for pokemon := range cfg.pokedex {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}

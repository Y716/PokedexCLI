package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location-area/"
	}

	areas, err := cfg.pokeapiClient.ListLocations(&cfg.Next)
	if err != nil {
		return err
	}

	locations := areas.Results
	for _, loc := range locations {
		fmt.Printf("%s\n", loc.Name)
	}

	cfg.Next = areas.Next
	cfg.Previous = areas.Previous

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	areas, err := cfg.pokeapiClient.ListLocations(&cfg.Previous)
	if err != nil {
		return err
	}

	locations := areas.Results
	for _, loc := range locations {
		fmt.Printf("%s\n", loc.Name)
	}

	cfg.Next = areas.Next
	cfg.Previous = areas.Previous

	return nil
}

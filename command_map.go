package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location_areas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(url *config) error {
	if url.Next == "" {
		url.Next = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url.Next)
	if err != nil {
		return fmt.Errorf("Error accessing API: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error read data: %w", err)
	}
	var location_areas location_areas
	err = json.Unmarshal(body, &location_areas)
	if err != nil {
		return fmt.Errorf("Error unmarshaling data: %w", err)
	}
	locations := location_areas.Results
	for _, loc := range locations {
		fmt.Printf("%s\n", loc.Name)
	}
	url.Next = location_areas.Next
	url.Previous = location_areas.Previous

	return nil
}

func commandMapB(url *config) error {
	if url.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(url.Previous)
	if err != nil {
		return fmt.Errorf("Error accessing API: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error read data: %w", err)
	}
	var location_areas location_areas
	err = json.Unmarshal(body, &location_areas)
	if err != nil {
		return fmt.Errorf("Error unmarshaling data: %w", err)
	}
	locations := location_areas.Results
	for _, loc := range locations {
		fmt.Printf("%s\n", loc.Name)
	}
	url.Next = location_areas.Next
	url.Previous = location_areas.Previous

	return nil
}

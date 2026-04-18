package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListPokemons(pageURL string) (Pokemons_in_area, error) {
	data, gotCache := c.cache.Get(pageURL)

	if gotCache {
		var pokemons Pokemons_in_area
		err := json.Unmarshal(data, &pokemons)
		if err != nil {
			return Pokemons_in_area{}, fmt.Errorf("Error unmarshaling data: %w", err)
		}
		return pokemons, nil
	}

	res, err := c.httpClient.Get(pageURL)
	if err != nil {
		return Pokemons_in_area{}, fmt.Errorf("Error accessing API: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemons_in_area{}, fmt.Errorf("Error read data: %w", err)
	}
	c.cache.Add(pageURL, body)

	var pokemons Pokemons_in_area
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return Pokemons_in_area{}, fmt.Errorf("Error unmarshaling data: %w", err)
	}

	return pokemons, nil
}

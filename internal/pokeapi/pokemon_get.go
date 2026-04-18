package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon_Data, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	data, gotCache := c.cache.Get(url)

	if gotCache {
		var pokemon Pokemon_Data
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon_Data{}, fmt.Errorf("Error unmarshaling data: %w", err)
		}
		return pokemon, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon_Data{}, fmt.Errorf("Error accessing API: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon_Data{}, fmt.Errorf("Error read data: %w", err)
	}
	c.cache.Add(url, body)

	var pokemon Pokemon_Data
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon_Data{}, fmt.Errorf("Error unmarshaling data: %w", err)
	}
	return pokemon, nil

}

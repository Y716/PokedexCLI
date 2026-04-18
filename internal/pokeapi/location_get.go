package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListPokemons(pageURL string) (location, error) {
	data, gotCache := c.cache.Get(pageURL)

	if gotCache {
		var location_data location
		err := json.Unmarshal(data, &location_data)
		if err != nil {
			return location{}, fmt.Errorf("Error unmarshaling data: %w", err)
		}
		return location_data, nil
	}

	res, err := c.httpClient.Get(pageURL)
	if err != nil {
		return location{}, fmt.Errorf("Error accessing API: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return location{}, fmt.Errorf("Error read data: %w", err)
	}
	c.cache.Add(pageURL, body)

	var location_data location
	err = json.Unmarshal(body, &location_data)
	if err != nil {
		return location{}, fmt.Errorf("Error unmarshaling data: %w", err)
	}

	return location_data, nil
}

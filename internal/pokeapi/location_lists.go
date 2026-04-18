package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (location_areas, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, gotCache := c.cache.Get(url)

	if gotCache {
		var areas location_areas
		err := json.Unmarshal(data, &areas)
		if err != nil {
			return location_areas{}, fmt.Errorf("Error unmarshaling data: %w", err)
		}
		return areas, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return location_areas{}, fmt.Errorf("Error accessing API: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return location_areas{}, fmt.Errorf("Error read data: %w", err)
	}
	c.cache.Add(url, body)

	var areas location_areas
	err = json.Unmarshal(body, &areas)
	if err != nil {
		return location_areas{}, fmt.Errorf("Error unmarshaling data: %w", err)
	}

	return areas, nil
}

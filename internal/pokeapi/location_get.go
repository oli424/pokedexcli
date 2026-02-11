package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (Location, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + name
	if val, ok := c.cache.Get(url); ok {
		var loc Location
		if err := json.Unmarshal(val, &loc); err != nil {
			return Location{}, err
		}
		return loc, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var loc Location
	if err := json.Unmarshal(dat, &loc); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return loc, nil
}

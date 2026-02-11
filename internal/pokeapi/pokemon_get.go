package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	if val, ok := c.cache.Get(url); ok {
		var poke Pokemon
		if err := json.Unmarshal(val, &poke); err != nil {
			return Pokemon{}, err
		}
		return poke, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var poke Pokemon
	if err := json.Unmarshal(dat, &poke); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return poke, nil
}

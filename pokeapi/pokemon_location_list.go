package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Lists and caches pokemons on any location
func(c *Client) ListPokemonsLoc(pageURL *string, locationName string) (RespPokemonsLoc, error){
	url := baseURL + "/location-area/" + locationName
	if pageURL != nil {
		url = *pageURL
	}
	cached, ok := c.cache.Get(url)
	if ok {
		pokeResp := RespPokemonsLoc{}
		err := json.Unmarshal(cached, &pokeResp)
		if err != nil {
			return RespPokemonsLoc{}, err
		}
		return pokeResp, nil
	}
	req, err:= http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonsLoc{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return RespPokemonsLoc{}, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonsLoc{}, err
	}
	respPokemons := RespPokemonsLoc{}
	err = json.Unmarshal(dat, &respPokemons)
	if err != nil{
		return respPokemons, err
	}
	c.cache.Add(url, dat)
	return respPokemons,nil
}
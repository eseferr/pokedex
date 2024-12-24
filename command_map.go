package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *Config) error{
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
	url = *cfg.Next
	}
	res, err := http.Get(url)
	if err !=nil{
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	data, err := io.ReadAll((res.Body))
	if err != nil {
		return err
	}

	locationResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return err
	}
	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous
	for _, loc := range locationResp.Results{
		fmt.Println(loc.Name)
	}
	return nil
}
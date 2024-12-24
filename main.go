package main

import (
	"time"

	"github.com/eseferr/pokedexcli/pokeapi"
)


func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}

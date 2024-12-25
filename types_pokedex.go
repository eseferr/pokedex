package main

import "github.com/eseferr/pokedexcli/pokeapi"

type Pokedex struct{
	pokemons map[string]pokeapi.RespPokemon
}
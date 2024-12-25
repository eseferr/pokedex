package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/eseferr/pokedexcli/pokeapi"
)

func commandCatch(cfg *config, pokemonName string)error{
	fmt.Println("Throwing a Pokeball at "+ pokemonName+"...")
	pokemon, err:= cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return errors.New("Couldn't reached Pokemon: "+ pokemonName)
	}
	catchNumber := float64(rand.Intn(100))*(2/float64(pokemon.BaseExperience))
	if catchNumber > 1{
		//addPokedex(pokemon)
		fmt.Println(pokemonName+" was cought!")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}
func addPokedex(cfg *config, pokemon pokeapi.RespPokemon)error{
	//cfg.pokedex.pokemons[pokemon.Name] = pokemon
	return nil
}
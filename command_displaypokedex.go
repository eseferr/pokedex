package main

import "fmt"

func commandDisplayPokedex(cfg *config, _ string) error{
	if len(cfg.pokedex.pokemons) != 0 {
		fmt.Println(cfg.pokedex.pokemons["bulbasaur"])
	}else {
		fmt.Println("empty pokedex")
	}
	return nil
}
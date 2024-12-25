package main

import "fmt"

func commandPokedex(cfg *config, _ string) error{
	if len(cfg.pokedex.pokemons) != 0 {
		for _,pokemon := range cfg.pokedex.pokemons{
			fmt.Println("-"+pokemon.Name)
		}
	}else{
		fmt.Println("Your pokedex is empty!")
	}
	return nil
}
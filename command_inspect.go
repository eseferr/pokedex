package main

import (
	"fmt"
	"strconv"
)

func commandInspect(cfg *config, pokemonName string) error {
	 pokemon, ok :=  cfg.pokedex.pokemons[pokemonName]
		if !ok{
			fmt.Println("you have not caught that pokemon")
			return nil
		}
		fmt.Println("Name: "+pokemon.Name)
		fmt.Println("Name: "+strconv.Itoa(pokemon.Height))
		
		fmt.Println("Weight: "+strconv.Itoa(pokemon.Weight))
		
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats{
			fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
    	for _, t := range pokemon.Types {
        	fmt.Printf("  - %s\n", t.Type.Name)
    	}
	return nil
	}
	
	
	
	
	
	


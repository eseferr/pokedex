package main

import "fmt"

func commandExplore(cfg *config, locationName string)error{
	fmt.Println("Exploring " +locationName+"...")
	fmt.Println("Found Pokemon:")
	pokeResp, err := cfg.pokeapiClient.ListPokemonsLoc(cfg.pokemonURL, locationName)
	if err != nil {
		return err
	}
	for _,poke:= range pokeResp.PokemonEncounters{
		fmt.Println(poke.Pokemon.Name)
	}
	return nil


}
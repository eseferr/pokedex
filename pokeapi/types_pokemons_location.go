package pokeapi

// RespPokemonsLoc
type RespPokemonsLoc struct {
	PokemonEncounters []struct{
		Pokemon struct{
		Name string `json:"name"`
		URL string `json:"url"`
		}`json:"pokemon"`
	}	`json:"pokemon_encounters"`
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eseferr/pokedexcli/pokeapi"
)
func startRepl(cfg *config){
	scanner:= bufio.NewScanner(os.Stdin)
	args := ""
	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		commandName := userInput[0]
		if len(userInput)> 1 {
			args = userInput[1]
		}
		if cmd, exists := getCommands()[commandName]; exists{
			err := cmd.callback(cfg, args)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokemonURL *string
	pokedex Pokedex
}
type cliCommand struct{
	name string
	description string
	callback func(*config, string) error
} 
func getCommands() map[string]cliCommand{

return map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help":{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "displays the names of 20 location areas in the Pokemon world",
		callback: commandMap,
	},
	"mapb":{
		name: "mapb",
		description:"displays the names of previous 20 location areas in the Pokemon world",
		callback: commandMapb,
	},
	"explore":{
		name: "explore",
		description: "Explore Pokemons",
		callback: commandExplore,
	},
	"catch":{
		name: "catch",
		description: "Catch Pokemons",
		callback: commandCatch,
	},
	"inspect":{
		name: "inspect",
		description: "Inspect Pokemon",
		callback: commandInspect,

	},
	"pokedex":{
		name: "pokedex",
		description: "Display Pokedex",
		callback: commandDisplayPokedex,
	},
}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func startRepl(){
	scanner:= bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		commandName := userInput[0]
		if cmd, exists := getCommands()[commandName]; exists{
			err := cmd.callback()
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
type cliCommand struct{
	name string
	description string
	callback func() error
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
}
}
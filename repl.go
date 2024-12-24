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
		fmt.Printf("Your command was: %s", userInput[0])
	}
}



func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

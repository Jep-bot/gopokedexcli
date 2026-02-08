package main

import (
	"fmt"
	"bufio"
	"os"
) 
// Main function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanInputText := cleanInput(input) 
		fmt.Printf("Your command was: %v\n", cleanInputText[0])
	}
}

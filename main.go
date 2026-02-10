package main

import (
	"fmt"
	"bufio"
	"os"
) 
// Main function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandRegistry := getCommands()
	commandConfing := config{
		Next: "",
		Previous: "",
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanTextList := cleanInput(input) 
		switch cleanTextList[0] {
			case "mapb":
				commandRegistry["mapb"].callback(&commandConfing)
			case "map":
				commandRegistry["map"].callback(&commandConfing)
			case "help":
				commandRegistry["help"].callback(&commandConfing)
			case "exit":
				commandRegistry["exit"].callback(&commandConfing)
			default:
				fmt.Println("Command not supported")
			}	
	}
}

package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache" 
) 
// Main function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const interval = 7 * time.Second
	cache := pokecache.NewCache(interval)
	commandRegistry := getCommands()
	commandConfing := config{

		Next: "",
		Previous: "",
		Arg: "",
		Pokedex: make(map[string]Pokemon),
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanTextList := cleanInput(input) 
		switch cleanTextList[0] {
			case "mapb":
				commandRegistry["mapb"].callback(&commandConfing, &cache)
			case "map":
				commandRegistry["map"].callback(&commandConfing, &cache)
			case "explore":
				if len(cleanTextList) >= 2 {
					commandConfing.Arg = cleanTextList[1]
					commandRegistry["explore"].callback(&commandConfing, &cache)
				}else{
					fmt.Println("no location provided")
				}
			case "catch":
				if len(cleanTextList) >= 2 {
					commandConfing.Arg = cleanTextList[1]
					commandRegistry["catch"].callback(&commandConfing, &cache)
				}else{
					fmt.Println("no pokemon provided")
				}
				for key,_ := range commandConfing.Pokedex{
					fmt.Println(key)
				}
			case "help":
				commandRegistry["help"].callback(&commandConfing, &cache)
			case "exit":
				commandRegistry["exit"].callback(&commandConfing, &cache)
			default:
				fmt.Println("Command not supported")
			}	
	}
}

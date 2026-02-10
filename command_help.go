package main
import (
	"fmt"
)

func commandHelp(urls *config ) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commandRegistry := getCommands()
	for _, value := range commandRegistry{
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}



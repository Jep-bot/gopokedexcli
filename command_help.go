package main
import (
	"fmt"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache"
)

func commandHelp(urls *config, cache *pokecache.Cache) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commandRegistry := getCommands()
	for _, value := range commandRegistry{
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}



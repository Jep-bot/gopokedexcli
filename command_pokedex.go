package main
import (
	"fmt"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache"
)
func commandPokedex(urls *config, cache *pokecache.Cache) error{
	fmt.Println("Your Pokedex:")
	for key,_ := range urls.Pokedex{
		fmt.Printf("  - %v\n", key )
	}
	return nil
}

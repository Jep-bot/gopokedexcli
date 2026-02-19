package main
import (
	"fmt"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache"
)
func commandInspect(urls *config, cache *pokecache.Cache) error{
	pokemon := urls.Pokedex[urls.Arg]
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStatus:\n",pokemon.Name, pokemon.Height, pokemon.Weight)
	for _,value := range pokemon.Stats{
		fmt.Printf("    -%v: %v\n", value.Stat.Name, value.BaseStat )
	}	
	fmt.Printf("Types:\n")
	for _,value := range pokemon.Types{
		fmt.Printf("    -%v\n",value.Type.Name)	
	}
	return nil
}

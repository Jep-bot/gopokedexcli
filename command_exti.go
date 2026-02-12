package main
import (
	"fmt"
	"os"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache"
)

func commandExit(urls *config, cache *pokecache.Cache) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

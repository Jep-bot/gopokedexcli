package main
import (
	"strings"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

type config struct {
	Next	string
	Previous string
	Arg string
	Pokedex map[string]Pokemon
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
		"help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
		"map": {
        name:        "map",
        description: "Displays the names of 20 locations in the Pokemon world.",
        callback:    commandMap,
    },
		"mapb": {
        name:        "mapb",
        description: "Displays the previous results of the map command",
        callback:    commandMapB,
    },
		"explore": {
        name:        "explore",
        description: "Displays a list of all pokemon from the provided loaction name",
        callback:    commandExplore,
    },
		"catch": {
        name:        "catch",
        description: "Try to catch the provided Pokemon",
        callback:    commandCatch,
    },
	}
}

func cleanInput(text string) []string{
    // ...
		text = strings.TrimSpace(text) 
		tempList := strings.Split(text," ")
		clearWords := make([]string, 0)
		for _,word := range tempList{
			if len(strings.TrimSpace(word)) != 0 {
				clearWords = append(clearWords,strings.ToLower(word))
			}
		} 
		return clearWords
}

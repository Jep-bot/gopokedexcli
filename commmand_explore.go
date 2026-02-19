package main
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache"
)

type EncounterMethodAtLocation struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(urls *config, cache *pokecache.Cache) error{
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/",urls.Arg)
	fmt.Printf("Exploring %v...\n",urls.Arg)
	//fmt.Println(url)
	body, ok := cache.Get(url)	
	//fmt.Println(ok)
	if !ok {
		fmt.Println(body)
		
		res, err := http.Get(url)
		defer res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		ok = cache.Add(url, body)
		//fmt.Printf("%s", body)
	}
	encounters := EncounterMethodAtLocation{}
	err := json.Unmarshal(body, &encounters)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found Pokemom:")
	for _,encounter := range encounters.PokemonEncounters {
		fmt.Printf("  - %v\n", encounter.Pokemon.Name)
	}	
	return nil	
}

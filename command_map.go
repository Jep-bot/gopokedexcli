package main
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
	pokecache "github.com/Jep-bot/gopokedexcli/internal/pokecache"
)

type result struct {
	Name string `json:"name"`
	Url string	`json:"url"`
}

type locationArea struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []result `json:"results"` 
}

func getLocations(url string, cache *pokecache.Cache) (locationArea, error) {
	
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
	location := locationArea{}
	err := json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	//}

	}
	
	return location, nil
}

func commandMap(urls *config, cache *pokecache.Cache) error{
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", 0)
	if urls.Next != "" {
		url = urls.Next
	}
	location, _ := getLocations(url, cache)
	for _,loc := range location.Results {
		fmt.Println(loc.Name)
	}
	urls.Previous = location.Previous
	urls.Next = location.Next
	return nil
}

func commandMapB(urls *config, cache *pokecache.Cache ) error{
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", 0)
	if urls.Previous != "" {
		url = urls.Previous
	} 
	location, _ := getLocations(url, cache)
	for _,loc := range location.Results {
		fmt.Println(loc.Name)
	}

	urls.Previous = location.Previous
	urls.Next = location.Next
	return nil
}

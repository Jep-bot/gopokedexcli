package main
import (
	"strings"
)
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

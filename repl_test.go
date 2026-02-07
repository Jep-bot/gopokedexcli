package main

import (
				"testing"
				"reflect"
				"fmt"
)
func TestCleanInput(t *testing.T) {
  cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " Charmander       Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "234",
			expected: []string{"234"},
		},
		// add more cases here
	}
	failCount :=0
	passCount := len(cases)
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if !reflect.DeepEqual(len(c.expected),len(actual)){
			t.Fatalf(`
-----CleanInputLength-----------------------
Input: %v
Expected: %v
Actual: %v
Fail
					`, c.input, len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if !reflect.DeepEqual(expectedWord,word){
					t.Errorf(
`-----CleanInput-----------------------
Input: %v
Expected: %v
Actual: %v
Fail
					`, c.input, word, expectedWord)
					failCount++
			}
		 
		}
	}
	fmt.Println("CleanInput-----------------------")
	fmt.Printf("Number of cases %d:\n%d passed, %d failed\n", len(cases), passCount - failCount, failCount)
}

package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Split(lower, " ")

	// Removes empty strings
	wordsFiltered := []string{}
	for _, word := range words {
		if word != "" {
			wordsFiltered = append(wordsFiltered, word)
		}
	}

	return wordsFiltered
}

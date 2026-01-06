package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Split(strings.TrimSpace(lower), " ")

	cleanSplit := []string{}
	for _, word := range split {
		if word != "" {
			cleanSplit = append(cleanSplit, word)
		}
	}

	return cleanSplit
}

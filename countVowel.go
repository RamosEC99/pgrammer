package main

import "strings"

func CountVowel(text string) int {
	lowerText := strings.ToLower(text)
	vowels := "aeiou"
	count := 0
	for _, letter := range lowerText {
		if strings.ContainsRune(vowels, letter) {
			count++
		}
	}

	return count
}

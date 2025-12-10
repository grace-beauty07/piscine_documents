package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	var allRunes []rune
	for i, arg := range args {
		for _, r := range arg {
			allRunes = append(allRunes, r)
		}
		if i < len(args)-1 {
			allRunes = append(allRunes, ' ')
		}
	}

	vowels := []rune{}
	for _, r := range allRunes {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}

	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	vowelIndex := 0
	for i, r := range allRunes {
		if isVowel(r) {
			allRunes[i] = vowels[vowelIndex]
			vowelIndex++
		}
	}

	for _, r := range allRunes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

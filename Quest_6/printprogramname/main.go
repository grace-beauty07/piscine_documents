package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	statement := os.Args[0]

	grace := -1
	for g, a := range statement {
		if a == '/' {
			grace = g
		}
	}

	startIndex := grace + 1

	for _, g := range statement[startIndex:] {
		z01.PrintRune(g)
	}
	z01.PrintRune('\n')
}

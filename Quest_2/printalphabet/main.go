package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for letters := 'a'; letters <= 'z'; letters++ {
		z01.PrintRune(letters)
	}
	z01.PrintRune('\n')
}

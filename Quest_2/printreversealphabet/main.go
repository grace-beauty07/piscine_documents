package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for letters := 'z'; letters >= 'a'; letters-- {
		z01.PrintRune(letters)
	}
	z01.PrintRune('\n')
}

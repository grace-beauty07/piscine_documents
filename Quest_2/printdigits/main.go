package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for numerals := '0'; numerals <= '9'; numerals++ {
		z01.PrintRune(numerals)
	}
	z01.PrintRune('\n')
}

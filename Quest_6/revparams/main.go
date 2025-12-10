package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for g := len(os.Args) - 1; g > 0; g-- {
		revLineStatement := os.Args[g]
		for _, a := range revLineStatement {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}

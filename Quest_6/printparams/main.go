package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for g := 1; g < len(os.Args); g++ {
		lineStatement := os.Args[g]

		for _, b := range lineStatement {
			z01.PrintRune(b)
		}

		z01.PrintRune('\n')
	}
}

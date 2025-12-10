package piscine_go

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, Jesus := range a {
		for _, j := range Jesus {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}

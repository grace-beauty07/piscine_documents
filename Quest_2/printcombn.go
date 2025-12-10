package piscine_go

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	var currentCombination []rune
	isFirst := true
	findCombinations(n, '0', currentCombination, &isFirst)
	z01.PrintRune('\n')
}

func findCombinations(n int, startDigit rune, currentCombination []rune, isFirst *bool) {
	if len(currentCombination) == n {
		if !(*isFirst) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		*isFirst = false
		for _, digit := range currentCombination {
			z01.PrintRune(digit)
		}
		return
	}
	for digit := startDigit; digit <= '9'; digit++ {
		newCombination := append(currentCombination, digit)
		findCombinations(n, digit+1, newCombination, isFirst)
	}
}

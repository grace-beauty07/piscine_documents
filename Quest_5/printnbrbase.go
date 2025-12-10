package piscine_go

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		printNbrBaseRecursiveUint(uint(-nbr), base)
	} else {
		printNbrBaseRecursiveUint(uint(nbr), base)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func printNbrBaseRecursiveUint(nbr uint, base string) {
	baseLen := uint(len(base))
	if nbr >= baseLen {
		printNbrBaseRecursiveUint(nbr/baseLen, base)
	}
	z01.PrintRune(rune(base[nbr%baseLen]))
}

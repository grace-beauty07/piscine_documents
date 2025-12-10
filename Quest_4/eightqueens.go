package piscine_go

import "github.com/01-edu/z01"

func EightQueens() {
	var queens [8]int
	solve(queens, 0)
}

func solve(queens [8]int, col int) {
	if col == 8 {
		printSolution(queens)
		return
	}
	for row := 1; row <= 8; row++ {
		if canPlace(queens, col, row) {
			queens[col] = row
			solve(queens, col+1)
		}
	}
}

func canPlace(queens [8]int, col, row int) bool {
	for g := 0; g < col; g++ {
		b := queens[g]
		if b == row || b == row-(col-g) || b == row+(col-g) {
			return false
		}
	}
	return true
}

func printSolution(queens [8]int) {
	for _, row := range queens {
		z01.PrintRune(rune(row + '0'))
	}
	z01.PrintRune('\n')
}

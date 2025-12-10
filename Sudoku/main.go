package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		printError()
	}

	var board [9][9]byte

	for i := 0; i < 9; i++ {
		row := args[i]
		if len(row) != 9 {
			printError()
		}
		for j := 0; j < 9; j++ {
			ch := row[j]
			if (ch < '1' || ch > '9') && ch != '.' {
				printError()
			}
			board[i][j] = ch
		}
	}

	if !solve(&board) {
		printError()
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func printError() {
	fmt.Println("Error")
	os.Exit(0)
}

func isValid(board [9][9]byte, row, col int, ch byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == ch || board[i][col] == ch {
			return false
		}
	}
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] == ch {
				return false
			}
		}
	}
	return true
}

func solve(board *[9][9]byte) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				for ch := byte('1'); ch <= '9'; ch++ {
					if isValid(*board, r, c, ch) {
						board[r][c] = ch
						if solve(board) {
							return true
						}
						board[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

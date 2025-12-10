package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for g := 0; g < y; g++ {
		if g == 0 || g == y-1 {
			drawTopBottom(x)
		} else {
			drawSide(x)
		}
	}
	z01.PrintRune('\n')
}

func drawTopBottom(row int) {
	strBegin := 'A'
	strBtwn := 'B'
	strEnd := 'C'

	for g := 0; g < row; g++ {
		switch g {
		case 0:
			z01.PrintRune(strBegin)
		case row - 1:
			z01.PrintRune(strEnd)
		default:
			z01.PrintRune(strBtwn)
		}
	}
	z01.PrintRune('\n')
}

func drawSide(row int) {
	strBtwn := ' '
	strSide := 'B'

	for g := 0; g < row; g++ {
		if g == 0 || g == row-1 {
			z01.PrintRune(strSide)
		} else {
			z01.PrintRune(strBtwn)
		}
	}
	z01.PrintRune('\n')
}

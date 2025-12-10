package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for g := 0; g < y; g++ {
		switch g {
		case 0:
			drawTop(x)
		case y - 1:
			drawBottom(x)
		default:
			drawSides(x)
		}
	}
	z01.PrintRune('\n')
}

func drawTop(row int) {
	strStartEnd := 'A'
	strBetween := 'B'

	for g := 0; g < row; g++ {
		if g == 0 || g == row-1 {
			z01.PrintRune(strStartEnd)
		} else {
			z01.PrintRune(strBetween)
		}
	}
	z01.PrintRune('\n')
}

func drawBottom(row int) {
	strStartEnd := 'C'
	strBetween := 'B'

	for g := 0; g < row; g++ {
		if g == 0 || g == row-1 {
			z01.PrintRune(strStartEnd)
		} else {
			z01.PrintRune(strBetween)
		}
	}
	z01.PrintRune('\n')
}

func drawSides(row int) {
	strStartEnd := 'B'
	strBetween := ' '

	for g := 0; g < row; g++ {
		if g == 0 || g == row-1 {
			z01.PrintRune(strStartEnd)
		} else {
			z01.PrintRune(strBetween)
		}
	}
	z01.PrintRune('\n')
}

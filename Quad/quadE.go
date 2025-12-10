package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		switch i {
		case 0:
			drawTopFwd(x)
		case y - 1:
			drawBtmFwd(x)
		default:
			drawMidFwd(x)
		}
	}
	z01.PrintRune('\n')
}

func drawTopFwd(row int) {
	strStart := 'A'
	strBtwn := 'B'
	strEnd := 'C'

	for i := 0; i <= row-1; i++ {
		switch i {
		case 0:
			z01.PrintRune(strStart)
		case row - 1:
			z01.PrintRune(strEnd)
		default:
			z01.PrintRune(strBtwn)
		}
	}
	z01.PrintRune('\n')
}

func drawMidFwd(row int) {

	strB := 'B'
	strSpace := ' '

	for i := 0; i <= row-1; i++ {
		if i == 0 || i == row-1 {
			z01.PrintRune(strB)
		} else {
			z01.PrintRune(strSpace)
		}
	}
	z01.PrintRune('\n')
}

func drawBtmFwd(row int) {

	strStart := 'C'
	strBtwn := 'B'
	strEnd := 'A'

	for i := 0; i <= row-1; i++ {
		switch i {
		case 0:
			z01.PrintRune(strStart)
		case row - 1:
			z01.PrintRune(strEnd)
		default:
			z01.PrintRune(strBtwn)
		}
	}
	z01.PrintRune('\n')
}

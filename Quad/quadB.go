package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		switch i {
		case 0:
			drawTopForward(x)
		case y - 1:
			drawBottomForward(x)
		default:
			drwSide(x)
		}
	}
	z01.PrintRune('\n')
}

func drawTopForward(xAxis int) {
	strBgn := '/'
	strEnd := '\\'
	strBtw := '*'

	for i := 0; i < xAxis; i++ {
		switch i {
		case 0:
			z01.PrintRune(strBgn)
		case xAxis - 1:
			z01.PrintRune(strEnd)
		default:
			z01.PrintRune(strBtw)
		}
	}
	z01.PrintRune('\n')
}

func drwSide(xAxis int) {
	strBgnEnd := '*'
	strBtw := ' '

	for i := 0; i < xAxis; i++ {
		if i == 0 || i == xAxis-1 {
			z01.PrintRune(strBgnEnd)
		} else {
			z01.PrintRune(strBtw)
		}
	}
	z01.PrintRune('\n')
}

func drawBottomForward(xAxis int) {
	strBgn := '\\'
	strEnd := '/'
	strBtw := '*'

	for i := 0; i < xAxis; i++ {
		switch i {
		case 0:
			z01.PrintRune(strBgn)
		case xAxis - 1:
			z01.PrintRune(strEnd)
		default:
			z01.PrintRune(strBtw)
		}
	}
	z01.PrintRune('\n')
}

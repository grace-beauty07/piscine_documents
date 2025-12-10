package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		if i == 0 || i == y-1 {
			drawForward(x)
		} else {
			drwSides(x)
		}
	}
	z01.PrintRune('\n')
}

func drawForward(forward int) {
	stringBtw := '-'
	strBgnEnd := 'o'
	for i := 0; i < forward; i++ {
		if i == 0 || i == forward-1 {
			z01.PrintRune(strBgnEnd)
		} else {
			z01.PrintRune(stringBtw)
		}
	}
	z01.PrintRune('\n')
}

func drwSides(side int) {
	strDown := '|'

	for i := 0; i < side; i++ {
		if i == 0 || i == side-1 {
			z01.PrintRune(strDown)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

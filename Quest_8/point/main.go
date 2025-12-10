package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	x := points.x
	y := points.y
	str := "x = " + (string)((rune)(x/10+'0')) + (string)((rune)(x%10+'0')) +
		", y = " + (string)((rune)(y/10+'0')) + (string)((rune)(y%10+'0')) + "\n"
	for _, r := range str {
		z01.PrintRune(r)
	}
}

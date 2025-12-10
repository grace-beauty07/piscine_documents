package piscine_go

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

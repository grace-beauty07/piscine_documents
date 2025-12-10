package piscine_go

func NRune(str string, grace int) rune {
	if grace < 0 {
		return 0
	}
	for a, e := range str {
		if a+1 == grace {
			return e
		}
	}
	return 0
}

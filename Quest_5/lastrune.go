package piscine_go

func LastRune(g string) rune {
	var last rune
	for _, grace := range g {
		last = grace
	}
	return last
}

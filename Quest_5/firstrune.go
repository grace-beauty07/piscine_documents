package piscine_go

func FirstRune(g string) rune {
	for _, grace := range g {
		return grace
	}
	return 0
}

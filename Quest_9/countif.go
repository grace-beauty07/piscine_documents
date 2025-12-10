package piscine_go

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, h := range tab {
		if f(h) {
			count++
		}
	}
	return count
}

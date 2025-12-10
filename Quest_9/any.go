package piscine_go

func Any(f func(string) bool, a []string) bool {
	for _, s := range a {
		if f(s) {
			return true
		}
	}
	return false
}

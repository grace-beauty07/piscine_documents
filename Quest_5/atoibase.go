package piscine_go

func isValidBaseCheck(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' {
			return false
		}
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func indexInBase(ch rune, base string) int {
	for i, b := range base {
		if b == ch {
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	if !isValidBaseCheck(base) {
		return 0
	}
	baseLen := len(base)
	result := 0
	for _, ch := range s {
		val := indexInBase(ch, base)
		if val == -1 {
			return 0
		}
		result = result*baseLen + val
	}
	return result
}

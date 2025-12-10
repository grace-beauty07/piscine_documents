package piscine_go

func TrimAtoi(s string) int {
	sign := 1
	numStarted := false
	result := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if !numStarted && ch == '-' {
			sign = -1
			continue
		} else if !numStarted && ch == '+' {
			continue
		}

		if ch >= '0' && ch <= '9' {
			numStarted = true
			result = result*10 + int(ch-'0')
		} else if numStarted {
			continue
		}
	}

	if !numStarted {
		return 0
	}
	return sign * result
}

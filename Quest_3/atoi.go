package piscine_go

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	i := 0

	if s[0] == '+' {
		i = 1
	} else if s[0] == '-' {
		sign = -1
		i = 1
	}

	result := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}

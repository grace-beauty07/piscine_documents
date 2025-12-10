package piscine_go

func BasicAtoi(s string) int {
	result := 0
	for _, ch := range s {
		digit := int(ch - '0')
		result = result*10 + digit
	}
	return result
}

package piscine_go

func ToUpper(s string) string {
	result := []rune{}

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r -= 'a' - 'A'
		}
		result = append(result, r)
	}

	return string(result)
}

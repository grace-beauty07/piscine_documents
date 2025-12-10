package piscine_go

func ToLower(s string) string {
	result := []rune{}

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}
		result = append(result, r)
	}

	return string(result)
}

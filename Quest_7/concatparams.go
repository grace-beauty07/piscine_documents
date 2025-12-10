package piscine_go

func ConcatParams(parameter []string) string {
	var result string
	for p, g := range parameter {
		result += g
		if p < len(parameter)-1 {
			result += "\n"
		}
	}
	return result
}

package piscine_go

func SplitWhiteSpaces(s string) []string {
	var result []string
	Jesus := ""

	for _, g := range s {
		if g == ' ' || g == '\t' || g == '\n' {
			if Jesus != "" {
				result = append(result, Jesus)
				Jesus = ""
			}
		} else {
			Jesus += string(g)
		}
	}

	if Jesus != "" {
		result = append(result, Jesus)
	}

	return result
}

package piscine_go

func BasicJoin(elems []string) string {
	result := ""
	for _, str := range elems {
		result += str
	}
	return result
}

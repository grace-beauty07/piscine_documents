package piscine_go

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	for i := 2; i <= nb; i++ {
		result = result * i
		if result < 0 {
			return 0
		}
	}
	return result
}

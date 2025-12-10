package piscine_go

func IterativePower(numeral int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	result := 1
	for grace := 1; grace <= power; grace++ {
		result = result * numeral
	}
	return result
}

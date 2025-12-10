package piscine_go

func RecursiveFactorial(factorial int) int {
	if factorial < 0 {
		return 0
	}
	if factorial == 0 || factorial == 1 {
		return 1
	}

	if factorial > 20 {
		return 0
	}
	return factorial * RecursiveFactorial(factorial-1)
}

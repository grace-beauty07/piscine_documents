package piscine_go

func Sqrt(grace int) int {
	if grace < 0 {
		return 0
	}
	if grace == 0 || grace == 1 {
		return grace
	}

	for g := 1; g <= grace/2+1; g++ {
		if g*g == grace {
			return g
		}
		if g*g > grace {
			break
		}
	}
	return 0
}

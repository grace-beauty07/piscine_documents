package piscine_go

func IsPrime(grace int) bool {
	if grace <= 1 {
		return false
	}
	if grace == 2 {
		return true
	}
	if grace%2 == 0 {
		return false
	}

	for g := 3; g*g <= grace; g += 2 {
		if grace%g == 0 {
			return false
		}
	}
	return true
}

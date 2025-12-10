package piscine_go

func FindNextPrime(grace int) int {
	if grace <= 2 {
		return 2
	}
	for {
		if IsPrime(grace) {
			return grace
		}
		grace++
	}
}

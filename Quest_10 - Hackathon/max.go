package piscine_go

func Max(a []int) int {
	maximumvalue := 0
	for g := 0; g < len(a); g++ {
		if a[g] > maximumvalue {
			maximumvalue = a[g]
		}
	}
	return maximumvalue
}

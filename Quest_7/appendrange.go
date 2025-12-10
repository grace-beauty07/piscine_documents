package piscine_go

func AppendRange(small, great int) []int {
	if small >= great {
		return nil
	}

	var result []int
	for i := small; i < great; i++ {
		result = append(result, i)
	}
	return result
}

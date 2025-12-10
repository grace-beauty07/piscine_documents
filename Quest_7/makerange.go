package piscine_go

func MakeRange(little, big int) []int {
	if little >= big {
		return nil
	}

	size := big - little
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = little + i
	}
	return result
}

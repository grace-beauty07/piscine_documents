package piscine_go

func StringToIntSlice(str string) []int {
	var nums []int

	for _, char := range str {
		nums = append(nums, int(char))
	}

	return nums
}

package piscine_go

func ShoppingListSort(slice []string) []string {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

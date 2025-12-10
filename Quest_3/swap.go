package piscine_go

func Swap(a *int, b *int) {
	aValue := *a
	bValue := *b

	*a = bValue
	*b = aValue
}

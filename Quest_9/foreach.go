package piscine_go

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

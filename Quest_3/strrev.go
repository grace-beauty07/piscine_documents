package piscine_go

func StrRev(s string) string {
	revStr := ""

	for i := StrLen(s) - 1; i >= 0; i-- {
		revStr += string(s[i])
	}

	return revStr
}

package piscine_go

func ConvertBase(nbr, baseJesus, baseChrist string) string {
	decimal := 0
	baseLenJesus := len(baseJesus)

	for _, r := range nbr {
		decimal = decimal*baseLenJesus + findIndex(r, baseJesus)
	}

	if decimal == 0 {
		return string(baseChrist[0])
	}

	baseLenChrist := len(baseChrist)
	result := ""
	for decimal > 0 {
		remainder := decimal % baseLenChrist
		result = string(baseChrist[remainder]) + result
		decimal /= baseLenChrist
	}
	return result
}

func findIndex(r rune, base string) int {
	for j, g := range base {
		if g == r {
			return j
		}
	}
	return -1
}

package piscine_go

func RockAndRoll(digit int) string {
	if digit < 0 {
		return "error: number is negative\n"
	}
	if digit%2 == 0 && digit%3 == 0 {
		return "rock and roll\n"
	}
	if digit%2 == 0 {
		return "rock\n"
	}
	if digit%3 == 0 {
		return "roll\n"
	}
	return "error: non divisible\n"
}

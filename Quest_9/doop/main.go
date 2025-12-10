package main

import (
	"os"
)

const (
	MaxInt64 = 1<<63 - 1
	MinInt64 = -1 << 63
)

func printNbr(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte{'0'})
		return
	}
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}
	var digits []byte
	for n > 0 {
		d := byte(n % 10)
		digits = append([]byte{'0' + d}, digits...)
		n /= 10
	}
	os.Stdout.Write(digits)
}

func parseInt(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	i := 0
	sign := int64(1)
	if s[0] == '-' {
		sign = -1
		i++
		if i == len(s) {
			return 0, false
		}
	}
	var n int64
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := int64(c - '0')
		// overflow check before multiplication
		if n > (MaxInt64-d)/10 {
			return 0, false
		}
		n = n*10 + d
	}
	n *= sign
	if n > MaxInt64 || n < MinInt64 {
		return 0, false
	}
	return n, true
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, okA := parseInt(os.Args[1])
	op := os.Args[2]
	b, okB := parseInt(os.Args[3])

	if !okA || !okB {
		return
	}

	var result int64
	switch op {
	case "+":
		if (b > 0 && a > MaxInt64-b) || (b < 0 && a < MinInt64-b) {
			return
		}
		result = a + b
	case "-":
		if (b > 0 && a < MinInt64+b) || (b < 0 && a > MaxInt64+b) {
			return
		}
		result = a - b
	case "*":
		if a > 0 && b > 0 && a > MaxInt64/b ||
			a > 0 && b < 0 && b < MinInt64/a ||
			a < 0 && b > 0 && a < MinInt64/b ||
			a < 0 && b < 0 && a < MaxInt64/b {
			return
		}
		result = a * b
	case "/":
		if b == 0 {
			printStr("No division by 0\n")
			return
		}
		if a == MinInt64 && b == -1 {
			return // overflow
		}
		result = a / b
	case "%":
		if b == 0 {
			printStr("No modulo by 0\n")
			return
		}
		if a == MinInt64 && b == -1 {
			return // overflow
		}
		result = a % b
	default:
		return
	}

	printNbr(result)
	os.Stdout.Write([]byte{'\n'})
}

package main

import (
	"fmt"
	"os"
)

func hasPrefix(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func sortString(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		for j := 0; j < len(r)-1-i; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
	return string(r)
}

func main() {
	args := os.Args[1:]

	var insertVal string
	orderFlag := false
	helpFlag := false
	baseStr := ""

	for _, arg := range args {
		switch {
		case arg == "--help" || arg == "-h":
			helpFlag = true
		case arg == "--order" || arg == "-o":
			orderFlag = true
		case hasPrefix(arg, "--insert="):
			insertVal = arg[len("--insert="):]
		case hasPrefix(arg, "-i="):
			insertVal = arg[len("-i="):]
		default:
			if baseStr == "" {
				baseStr = arg
			}
		}
	}

	if helpFlag || len(args) == 0 {
		printHelp()
		return
	}

	finalStr := baseStr
	if insertVal != "" {
		finalStr += insertVal
	}

	if orderFlag {
		finalStr = sortString(finalStr)
	}

	fmt.Println(finalStr)
}

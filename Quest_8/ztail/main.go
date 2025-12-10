package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 || args[0] != "-c" {
		return
	}

	n := 0
	for _, ch := range args[1] {
		if ch < '0' || ch > '9' {
			return
		}
		n = n*10 + int(ch-'0')
	}
	if n <= 0 {
		return
	}

	files := args[2:]
	if len(files) == 0 {
		return
	}

	exitStatus := 0
	multiple := len(files) > 1

	for i, name := range files {

		data, err := os.ReadFile(name)
		if err != nil {

			fmt.Printf("open %s: %v\n", name, err.(*os.PathError).Err)
			exitStatus = 1
			continue
		}

		if multiple {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", name)
		}

		if len(data) <= n {
			fmt.Printf("%s", string(data))
		} else {
			fmt.Printf("%s", string(data[len(data)-n:]))
		}
	}

	if exitStatus != 0 {
		os.Exit(1)
	}
}

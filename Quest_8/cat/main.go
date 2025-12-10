package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		err := z01.PrintRune(r)
		if err != nil {
			return
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if line != "" {
				printString(line)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				printString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
		}
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if line != "" {
				printString(line)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				printString("ERROR: " + err.Error() + "\n")
				file.Close()
				os.Exit(1)
			}
		}
		file.Close()
	}
}

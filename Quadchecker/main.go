package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	inputStr := string(input)
	if len(inputStr) == 0 {
		return
	}

	inputStr = strings.TrimRight(inputStr, "\n")
	lines := strings.Split(inputStr, "\n")
	height := len(lines)
	width := 0
	if height > 0 {
		width = len(lines[0])
	}

	if width == 0 || height == 0 {
		fmt.Println("Not a quad function")
		return
	}

	matches := []string{}
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	for _, quad := range quads {
		cmd := exec.Command("./"+quad, fmt.Sprintf("%d", width), fmt.Sprintf("%d", height))
		output, err := cmd.Output()
		if err != nil {
			continue
		}

		outputStr := strings.TrimRight(string(output), "\n")
		if outputStr == inputStr {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad, width, height))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}

package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type InputOptions struct {
	Day        int
	FileNumber int
}

func ReadInput(opt InputOptions) []string {
	day := fmt.Sprintf("day%02d", opt.Day)

	var filename string
	if opt.FileNumber == 0 {
		filename = "input"
	} else {
		filename = fmt.Sprintf("input_%v", opt.FileNumber)
	}

	path := fmt.Sprintf("%v/%v", day, filename)

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Cannot find input file %v", path)
		os.Exit(1)
	}

	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

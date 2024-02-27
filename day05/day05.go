package day05

import (
	"fmt"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 5, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	fmt.Println(lines)
	return "", nil
}

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 5, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	fmt.Println(lines)
	return "", nil
}

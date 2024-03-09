package day09

import (
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

/*
 * PART 1
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 9, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	answer := 0
	for _, l := range lines {
		numbers := strings.Split(l, " ")
		s := []int{}
		for _, n := range numbers {
			number, _ := strconv.Atoi(n)
			s = append(s, number)
		}
		answer += extrapolateLast(s)
	}

	return strconv.Itoa(answer), nil
}

/*
 * PART 2
 */

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 9, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	answer := 0
	for _, l := range lines {
		numbers := strings.Split(l, " ")
		s := []int{}
		for _, n := range numbers {
			number, _ := strconv.Atoi(n)
			s = append(s, number)
		}
		answer += extrapolateFirst(s)
	}

	return strconv.Itoa(answer), nil
}

/*
 * HELPER
 */

func extrapolateLast(sequence []int) int {
	// Terminal case: If sequence is an array of 0, return 0
	found := false
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			found = true
		}
	}
	if !found {
		return 0
	}

	s := []int{}
	for i := 1; i < len(sequence); i++ {
		s = append(s, sequence[i]-sequence[i-1])
	}
	return extrapolateLast(s) + sequence[len(sequence)-1]
}

func extrapolateFirst(sequence []int) int {
	// Terminal case: If sequence is an array of 0, return 0
	found := false
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			found = true
		}
	}
	if !found {
		return 0
	}

	s := []int{}
	for i := 1; i < len(sequence); i++ {
		s = append(s, sequence[i]-sequence[i-1])
	}
	return sequence[0] - extrapolateFirst(s)
}

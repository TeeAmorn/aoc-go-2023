package day01

import (
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 1, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, line := range lines {
		numStr := ""
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= '0' && c <= '9' {
				numStr += string(c)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c >= '0' && c <= '9' {
				numStr += string(c)
				break
			}
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "", err
		}
		sum += num
	}

	return strconv.Itoa(sum), nil
}

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 1, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	sum := 0
	for _, line := range lines {
		numStr := ""
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= '0' && c <= '9' {
				numStr += string(c)
				goto next
			}
			for _, p := range prefixes {
				if strings.HasPrefix(line[i:], p) {
					numStr += prefixMap[p]
					goto next
				}
			}
		}
	next:
		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c >= '0' && c <= '9' {
				numStr += string(c)
				goto final
			}
			for _, p := range prefixes {
				if strings.HasPrefix(line[i:], p) {
					numStr += prefixMap[p]
					goto final
				}
			}
		}
	final:
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return "", err
		}
		sum += num
	}

	return strconv.Itoa(sum), nil
}

var prefixMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var prefixes = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

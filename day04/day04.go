package day04

import (
	"math"
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 4, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	points := 0
	for _, line := range lines {
		scoresString := strings.Split(line, ": ")[1]
		scores := strings.Split(scoresString, " | ")

		winningSet := make(map[string]bool)
		for i := 0; i < len(scores[0]); i += 3 {
			winningSet[scores[0][i:i+2]] = true
		}

		wonCount := 0
		for i := 0; i < len(scores[1]); i += 3 {
			num := scores[1][i : i+2]
			_, hasNum := winningSet[num]

			if hasNum {
				wonCount++
			}
		}

		if wonCount > 0 {
			points += int(math.Pow(float64(2), float64(wonCount-1)))
		}
	}

	return strconv.Itoa(points), nil
}

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 4, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	sum := 0
	dp := make([]int, len(lines))
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]

		scoresString := strings.Split(line, ": ")[1]
		scores := strings.Split(scoresString, " | ")

		winningSet := make(map[string]bool)
		for j := 0; j < len(scores[0]); j += 3 {
			winningSet[scores[0][j:j+2]] = true
		}

		wonCount := 0
		for j := 0; j < len(scores[1]); j += 3 {
			num := scores[1][j : j+2]
			_, hasNum := winningSet[num]

			if hasNum {
				wonCount++
			}
		}

		totalCardWons := 1
		for j := range wonCount {
			totalCardWons += dp[i+j+1]
		}
		dp[i] = totalCardWons
		sum += dp[i]
	}

	return strconv.Itoa(sum), nil
}

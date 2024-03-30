package day11

import (
	"fmt"
	"math"
	"strconv"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

/*
 * PART 1
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 11, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	galaxies := getGalaxyPositions(lines)
	ans := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			ans += getManhattanDistance(galaxies[i], galaxies[j])
		}
	}

	return strconv.Itoa(ans), nil
}

/*
 * PART 2
 */

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 11, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)
	fmt.Println(lines)
	return "", nil
}

/*
 * HELPER
 */

type pos struct {
	x int
	y int
}

func getGalaxyPositions(lines []string) []pos {
	galaxies := []pos{}
	nonEmptyRows, nonEmptyCols := getNonEmptyRowsAndCols(lines)

	rowOffset := 0
	for i, line := range lines {
		colOffset := 0

		if !nonEmptyRows[i] {
			rowOffset += 1
			continue
		}

		for j, c := range line {
			if !nonEmptyCols[j] {
				colOffset += 1
				continue
			}

			if c == '.' {
				continue
			}

			galaxies = append(galaxies, pos{i + rowOffset, j + colOffset})
		}
	}

	return galaxies
}

func getNonEmptyRowsAndCols(lines []string) (rows, cols map[int]bool) {
	rows = make(map[int]bool)
	cols = make(map[int]bool)

	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	return rows, cols
}

func getManhattanDistance(a, b pos) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

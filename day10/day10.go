package day10

import (
	"fmt"
	"strconv"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

/*
 * PART 1
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 10, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	start := findStartTile(lines)
	possibleNeighbors := []coordinate{}

	if isInside(start.i-1, start.j, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i - 1, start.j})
	}
	if isInside(start.i+1, start.j, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i + 1, start.j})
	}
	if isInside(start.i, start.j-1, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i, start.j - 1})
	}
	if isInside(start.i, start.j+1, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i, start.j + 1})
	}

	length := 0
	for _, n := range possibleNeighbors {
		seen := make(map[coordinate]bool)
		found, steps := lengthToStart(n, lines, seen)
		if found {
			length = steps + 1
		}
	}

	return strconv.Itoa(length / 2), nil
}

/*
 * PART 2
 */

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 10, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	fmt.Println(lines)
	return "", nil
}

/*
 * HELPER
 */

type coordinate struct {
	i int
	j int
}

func findStartTile(tiles []string) coordinate {
	for i, l := range tiles {
		for j, c := range l {
			if c == 'S' {
				return coordinate{i, j}
			}
		}
	}
	return coordinate{-1, -1}
}

func lengthToStart(pos coordinate, tiles []string, seen map[coordinate]bool) (bool, int) {
	if tiles[pos.i][pos.j] == 'S' {
		return true, 0
	}

	seen[pos] = true
	neighbors := getNeighbors(pos, tiles)
	for _, n := range neighbors {
		if !seen[n] {
			if len(seen) == 1 && tiles[n.i][n.j] == 'S' {
				continue
			}
			found, steps := lengthToStart(n, tiles, seen)
			return found, steps + 1
		}
	}

	return false, 0
}

func getNeighbors(pos coordinate, tiles []string) []coordinate {
	neighbors := []coordinate{}
	curr := tiles[pos.i][pos.j]

	if curr == '|' {
		ni, nj := pos.i-1, pos.j
		if canGetToFromBelow(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}

		ni, nj = pos.i+1, pos.j
		if canGetToFromAbove(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}
	} else if curr == '-' {
		ni, nj := pos.i, pos.j-1
		if canGetToFromRight(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}

		ni, nj = pos.i, pos.j+1
		if canGetToFromLeft(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}
	} else if curr == 'L' {
		ni, nj := pos.i-1, pos.j
		if canGetToFromBelow(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}

		ni, nj = pos.i, pos.j+1
		if canGetToFromLeft(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}
	} else if curr == 'J' {
		ni, nj := pos.i-1, pos.j
		if canGetToFromBelow(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}

		ni, nj = pos.i, pos.j-1
		if canGetToFromRight(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}
	} else if curr == '7' {
		ni, nj := pos.i+1, pos.j
		if canGetToFromAbove(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}

		ni, nj = pos.i, pos.j-1
		if canGetToFromRight(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}
	} else if curr == 'F' {
		ni, nj := pos.i+1, pos.j
		if canGetToFromAbove(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}

		ni, nj = pos.i, pos.j+1
		if canGetToFromLeft(ni, nj, tiles) {
			neighbors = append(neighbors, coordinate{ni, nj})
		}
	}

	return neighbors
}

func isInside(i, j int, tiles []string) bool {
	return i >= 0 && i < len(tiles) && j >= 0 && j < len(tiles[0])
}

func canGetToFromBelow(i, j int, tiles []string) bool {
	if !isInside(i, j, tiles) {
		return false
	}

	t := tiles[i][j]
	return t == '7' || t == 'F' || t == '|' || t == 'S'
}

func canGetToFromAbove(i, j int, tiles []string) bool {
	if !isInside(i, j, tiles) {
		return false
	}

	t := tiles[i][j]
	return t == 'L' || t == 'J' || t == '|' || t == 'S'
}

func canGetToFromLeft(i, j int, tiles []string) bool {
	if !isInside(i, j, tiles) {
		return false
	}

	t := tiles[i][j]
	return t == 'J' || t == '7' || t == '-' || t == 'S'
}

func canGetToFromRight(i, j int, tiles []string) bool {
	if !isInside(i, j, tiles) {
		return false
	}

	t := tiles[i][j]
	return t == 'L' || t == 'F' || t == '-' || t == 'S'
}

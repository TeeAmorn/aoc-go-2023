package day10

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

/*
 * PART 1
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 10, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	_, length := getLoop(lines)
	return strconv.Itoa(length / 2), nil
}

/*
 * PART 2
 */

type item struct {
	index int
	pipe  rune
}

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 10, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	path, _ := getLoop(lines)

	rows := make(map[int][]item)
	for _, t := range path {
		value := []rune(lines[t.i])[t.j]
		if value == 'S' {
			value = getStartPipeType(path)
		}
		pipe := item{t.j, value}
		if v, ok := rows[t.i]; ok {
			rows[t.i] = append(v, pipe)
		} else {
			rows[t.i] = []item{pipe}
		}
	}
	for _, r := range rows {
		sort.SliceStable(r, func(i, j int) bool {
			return r[i].index < r[j].index
		})
	}

	answer := 0
	for i := range len(lines) {
		for j := range len(lines[0]) {
			inside := isInsideLoop(coordinate{i, j}, rows)
			if inside {
				answer += 1
			}
		}
	}

	loop := make(map[coordinate]bool)
	for _, p := range path {
		loop[p] = true
	}

	return strconv.Itoa(answer), nil
}

/*
 * HELPER
 */

func getLoop(lines []string) ([]coordinate, int) {

	start := findStartTile(lines)
	possibleNeighbors := []coordinate{}

	if isInside(start.i-1, start.j, lines) && canGetToFromBelow(start.i-1, start.j, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i - 1, start.j})
	}
	if isInside(start.i+1, start.j, lines) && canGetToFromAbove(start.i+1, start.j, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i + 1, start.j})
	}
	if isInside(start.i, start.j-1, lines) && canGetToFromRight(start.i, start.j-1, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i, start.j - 1})
	}
	if isInside(start.i, start.j+1, lines) && canGetToFromLeft(start.i, start.j+1, lines) {
		possibleNeighbors = append(possibleNeighbors, coordinate{start.i, start.j + 1})
	}

	length := 0
	var path []coordinate
	for _, n := range possibleNeighbors {
		seen := make(map[coordinate]bool)
		found, steps, p := lengthToStart(n, lines, seen)
		if found {
			length = steps + 1
			path = append(p, n)
			break
		}
	}

	return path, length
}

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

func lengthToStart(pos coordinate, tiles []string, seen map[coordinate]bool) (bool, int, []coordinate) {
	if tiles[pos.i][pos.j] == 'S' {
		return true, 0, []coordinate{}
	}

	seen[pos] = true
	neighbors := getNeighbors(pos, tiles)
	for _, n := range neighbors {
		if !seen[n] {
			if len(seen) == 1 && tiles[n.i][n.j] == 'S' {
				continue
			}
			found, steps, path := lengthToStart(n, tiles, seen)
			path = append(path, n)
			return found, steps + 1, path
		}
	}

	return false, 0, []coordinate{}
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

func getStartPipeType(path []coordinate) rune {
	s := path[0]
	a := path[1]
	b := path[len(path)-1]

	x := deduceRelativeLocation(s, a)
	y := deduceRelativeLocation(s, b)
	locations := map[string]bool{x: true, y: true}

	if locations["left"] && locations["above"] {
		return 'J'
	} else if locations["left"] && locations["right"] {
		return '-'
	} else if locations["left"] && locations["below"] {
		return '7'
	} else if locations["above"] && locations["right"] {
		return 'L'
	} else if locations["above"] && locations["below"] {
		return '|'
	} else {
		return 'F'
	}
}

func deduceRelativeLocation(a coordinate, b coordinate) string {
	if a.i-b.i == 1 && a.j-b.j == 0 {
		return "above"
	} else if a.i-b.i == -1 && a.j-b.j == 0 {
		return "below"
	} else if a.i-b.i == 0 && a.j-b.j == 1 {
		return "left"
	} else {
		return "right"
	}
}

func printTileMap(lines []string, loop map[coordinate]bool, enclosed map[coordinate]bool) {
	for i := range len(lines) {
		toPrint := ""
		for j := range len(lines[0]) {
			pos := coordinate{i, j}
			if enclosed[pos] {
				toPrint += "X "
			} else if _, ok := loop[pos]; ok {
				toPrint += "* "
			} else {
				toPrint += ". "
			}
		}
		fmt.Println(toPrint)
	}
}

func isInsideLoop(pos coordinate, rows map[int][]item) bool {
	row, ok := rows[pos.i]

	if !ok {
		return false
	}

	start := 0
	for start < len(row) {
		if row[start].index > pos.j {
			break
		} else if row[start].index == pos.j {
			return false
		}
		start++
	}

	if start == len(row) {
		return false
	}

	isEven := true
	var previousAngledWall rune
	for start < len(row) {
		if (row[start].pipe == 'J' && previousAngledWall == 'F') || (row[start].pipe == '7' && previousAngledWall == 'L') || row[start].pipe == '|' {
			isEven = !isEven
		} else if row[start].pipe == 'F' || row[start].pipe == 'L' {
			previousAngledWall = row[start].pipe
		}
		start++
	}

	return !isEven
}

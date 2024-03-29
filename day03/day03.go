package day03

import (
	"strconv"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 3, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	values := make([][]int, len(lines))
	for i := range values {
		values[i] = make([]int, len(lines[0]))
	}

	symbols := make([][2]int, 0)

	for i, line := range lines {
		numStr := ""
		for j, char := range line {
			if char >= '0' && char <= '9' {
				numStr += string(char)
			} else {
				if char != '.' {
					pos := [2]int{i, j}
					symbols = append(symbols, pos)
				}

				if len(numStr) > 0 {
					num, err := strconv.Atoi(numStr)

					if err != nil {
						return "", err
					}

					for dy := j - len(numStr); dy < j; dy++ {
						values[i][dy] = num
					}

					numStr = ""
				}
			}
		}
		if len(numStr) > 0 {
			num, err := strconv.Atoi(numStr)

			if err != nil {
				return "", err
			}

			for dy := len(line) - len(numStr); dy < len(line); dy++ {
				values[i][dy] = num
			}

			numStr = ""
		}
	}

	sum := 0

	dxdys := [8][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	for _, symbol := range symbols {
		x := symbol[0]
		y := symbol[1]
		for _, dxdy := range dxdys {
			dx := dxdy[0]
			dy := dxdy[1]
			newx := x + dx
			newy := y + dy

			if newx < 0 || newx >= len(lines) || newy < 0 || newy >= len(lines[0]) {
				continue
			}

			if values[newx][newy] == 0 {
				continue
			}

			sum += values[newx][newy]
			removeIsland(newx, newy, values)
		}
	}

	return strconv.Itoa(sum), nil
}

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 3, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	islands := make([][]int, len(lines))
	for i := range islands {
		islands[i] = make([]int, len(lines[0]))
	}

	islandValues := make([]int, 0)
	islandValues = append(islandValues, 0)
	islandCount := 0

	gears := make([][2]int, 0)

	for i, line := range lines {
		numStr := ""
		for j, char := range line {
			if char >= '0' && char <= '9' {
				numStr += string(char)
			} else {
				if char == '*' {
					pos := [2]int{i, j}
					gears = append(gears, pos)
				}

				if len(numStr) > 0 {
					islandCount += 1

					for dy := j - len(numStr); dy < j; dy++ {
						islands[i][dy] = islandCount
					}

					num, err := strconv.Atoi(numStr)

					if err != nil {
						return "", err
					}

					islandValues = append(islandValues, num)

					numStr = ""
				}
			}
		}
		if len(numStr) > 0 {
			islandCount += 1

			for dy := len(line) - len(numStr); dy < len(line); dy++ {
				islands[i][dy] = islandCount
			}

			num, err := strconv.Atoi(numStr)

			if err != nil {
				return "", err
			}

			islandValues = append(islandValues, num)

			numStr = ""
		}
	}

	sum := 0

	dxdys := [8][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
	for _, gear := range gears {
		x := gear[0]
		y := gear[1]

		parts := make([]int, 0)
		seen := make(map[int]bool)

		for _, dxdy := range dxdys {
			dx := dxdy[0]
			dy := dxdy[1]
			newx := x + dx
			newy := y + dy

			if newx < 0 || newx >= len(lines) || newy < 0 || newy >= len(lines[0]) {
				continue
			}

			_, hasSeen := seen[islands[newx][newy]]
			if islands[newx][newy] == 0 || hasSeen {
				continue
			}

			seen[islands[newx][newy]] = true
			parts = append(parts, islands[newx][newy])
		}

		if len(parts) == 2 {
			sum += islandValues[parts[0]] * islandValues[parts[1]]
		}
	}

	return strconv.Itoa(sum), nil
}

func removeIsland(i int, j int, values [][]int) {
	if j < 0 || j >= len(values[0]) || values[i][j] == 0 {
		return
	}

	values[i][j] = 0
	removeIsland(i, j-1, values)
	removeIsland(i, j+1, values)
}

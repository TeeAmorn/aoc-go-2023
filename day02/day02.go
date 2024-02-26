package day02

import (
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 2, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	sum := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		sets := strings.Split(line, ":")
		sets = strings.Split(sets[1], ";")

		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, colorCount := range colors {
				colorCount := strings.Trim(colorCount, " ")
				colorCountArr := strings.Split(colorCount, " ")

				count, err := strconv.Atoi(colorCountArr[0])
				if err != nil {
					return "", err
				}

				color := colorCountArr[1]

				switch color {
				case "red":
					if count > 12 {
						goto done
					}
				case "green":
					if count > 13 {
						goto done
					}
				case "blue":
					if count > 14 {
						goto done
					}
				}
			}
		}

		sum += i + 1
	done:
	}

	return strconv.Itoa(sum), nil
}

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 2, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	sum := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		sets := strings.Split(line, ":")
		sets = strings.Split(sets[1], ";")

		maxCount := cubeCount{}
		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, colorCount := range colors {
				colorCount := strings.Trim(colorCount, " ")
				colorCountArr := strings.Split(colorCount, " ")

				count, err := strconv.Atoi(colorCountArr[0])
				if err != nil {
					return "", err
				}

				color := colorCountArr[1]

				switch color {
				case "red":
					maxCount.red = max(maxCount.red, count)
				case "green":
					maxCount.green = max(maxCount.green, count)
				case "blue":
					maxCount.blue = max(maxCount.blue, count)
				}
			}
		}

		sum += maxCount.red * maxCount.green * maxCount.blue
	}

	return strconv.Itoa(sum), nil
}

type cubeCount struct {
	red   int
	blue  int
	green int
}

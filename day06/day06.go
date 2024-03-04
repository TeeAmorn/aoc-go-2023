package day06

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

/*
 * ==================== NOTE ====================
 * We are basically solving this quadratic equation:
 * x ( time - x ) > distance
 * which is equivalent to this:
 * x^2 - time * x + distance < 0
 *
 * If roots are whole numbers, then we have x ( time - x ) = distance.
 * So we have to increment/decrement the roots.
 * The answer for each race is the number of discrete
 * roots that satisfy the equation above.
 * ==============================================
 */

/*
 * PART 1
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 6, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	// Parse inputs
	times := getFloatsFromLine(lines[0])
	distances := getFloatsFromLine(lines[1])

	answer := 1
	for i, time := range times {
		distance := distances[i]
		roots, _ := computeQuadraticRoots(1, -time, distance)
		x1r := math.Ceil(roots.x1)
		x2r := math.Floor(roots.x2)

		var x1, x2 int
		if roots.x1 == x1r {
			x1 = int(x1r + 1)
		} else {
			x1 = int(x1r)
		}
		if roots.x2 == x2r {
			x2 = int(x2r - 1)
		} else {
			x2 = int(x2r)
		}
		answer *= x2 - x1 + 1
	}

	return strconv.Itoa(answer), nil
}

func getFloatsFromLine(line string) []float64 {
	floats := []float64{}

	input := strings.Split(line, ":")[1]
	input = strings.TrimSpace(input)

	var n string
	for _, c := range input {
		if c < '0' || c > '9' {
			if n == "" {
				continue
			}
			time, _ := strconv.Atoi(n)
			floats = append(floats, float64(time))
			n = ""
		} else {
			n += string(c)
		}
	}
	number, _ := strconv.Atoi(n)
	floats = append(floats, float64(number))

	return floats
}

/*
 * PART 2
 */

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 6, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	// Parse inputs
	time := getSingleFloatFromLine(lines[0])
	distance := getSingleFloatFromLine(lines[1])

	roots, _ := computeQuadraticRoots(1, -time, distance)
	x1r := math.Ceil(roots.x1)
	x2r := math.Floor(roots.x2)

	var x1, x2 int
	if roots.x1 == x1r {
		x1 = int(x1r + 1)
	} else {
		x1 = int(x1r)
	}
	if roots.x2 == x2r {
		x2 = int(x2r - 1)
	} else {
		x2 = int(x2r)
	}

	answer := x2 - x1 + 1
	return strconv.Itoa(answer), nil
}

/*
 * HELPER
 */

type quadraticRoots struct {
	x1 float64
	x2 float64
}

func computeQuadraticRoots(a, b, c float64) (quadraticRoots, error) {
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		return quadraticRoots{}, errors.New("No roots")
	}

	x1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	return quadraticRoots{x1, x2}, nil
}

func getSingleFloatFromLine(line string) float64 {
	var n string
	for _, c := range line {
		if c >= '0' && c <= '9' {
			n += string(c)
		}
	}
	number, _ := strconv.Atoi(n)
	return float64(number)
}

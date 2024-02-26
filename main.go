package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TeeAmorn/aoc-go-2023/day01"
)

func main() {
	day, part := parseArgs()

	challenges := [][]func() (string, error){
		{day01.Part1, day01.Part2},
	}

	if day > len(challenges) {
		fmt.Printf("First argument (challenge day) must be an integer between 1 and %v (inclusive)\n", len(challenges))
		os.Exit(1)
	}

	if part > len(challenges[day-1]) {
		fmt.Printf("Second argument (challenge part) must be an integer between 1 and %v (inclusive)\n", len(challenges[day-1]))
		os.Exit(1)
	}

	challenge := challenges[day-1][part-1]
	fmt.Printf("========== Day %v Part %v ==========\n", day, part)

	startTime := time.Now()
	result, err := challenge()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)

	fmt.Printf("Elapsed Time: %v\n", elapsed)
	fmt.Println("----------------------------------")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println("==================================")
}

func parseArgs() (int, int) {

	var day, part int
	var err error

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Must provide 2 arguments: challenge day and challenge part!")
		os.Exit(1)
	}

	day, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("First argument (challenge day) must be an integer!")
		os.Exit(1)
	}

	part, err = strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Second argument (challenge part) must be an integer!")
		os.Exit(1)
	}

	return day, part
}

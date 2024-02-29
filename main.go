package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TeeAmorn/aoc-go-2023/day01"
	"github.com/TeeAmorn/aoc-go-2023/day02"
	"github.com/TeeAmorn/aoc-go-2023/day03"
	"github.com/TeeAmorn/aoc-go-2023/day04"
	"github.com/TeeAmorn/aoc-go-2023/day05"
	"github.com/TeeAmorn/aoc-go-2023/day06"
)

func main() {
	day, part, inputFileNumber := parseArgs()

	challenges := [][]func(int) (string, error){
		{day01.Part1, day01.Part2},
		{day02.Part1, day02.Part2},
		{day03.Part1, day03.Part2},
		{day04.Part1, day04.Part2},
		{day05.Part1, day05.Part2},
		{day06.Part1, day06.Part2},
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
	result, err := challenge(inputFileNumber)
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

func parseArgs() (int, int, int) {

	var day, part, inputFileNumber int
	var err error

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Must provide at least 2 arguments: challenge day and challenge part!")
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

	if len(args) > 2 {
		inputFileNumber, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Third argument (input file number) must be an integer!")
			os.Exit(1)
		}
	} else {
		inputFileNumber = 0
	}

	return day, part, inputFileNumber
}

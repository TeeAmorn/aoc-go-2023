package day08

import (
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

/*
 * PART 1
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 8, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	graph := buildMap(lines[2:])
	directions := lines[0]
	idx := 0
	count := 0
	curr := "AAA"

	for curr != "ZZZ" {
		if directions[idx] == 'L' {
			curr = graph[curr].left
		} else {
			curr = graph[curr].right
		}
		idx = (idx + 1) % len(directions)
		count++
	}

	return strconv.Itoa(count), nil
}

/*
 * PART 2
 */

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 8, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)
	graph := buildMap(lines[2:])
	directions := lines[0]

	startingNodes := []string{}
	for node := range graph {
		if node[len(node)-1] == 'A' {
			startingNodes = append(startingNodes, node)
		}
	}

	// Paths are periodic, meaning once we get from A to Z, path will repeat
	// Therefore, the answer will be the least common multiple of the steps for all paths from A to Z
	answer := amortizedSearch(startingNodes[0], directions, 0, graph, make(map[string]int))
	for i := 1; i < len(startingNodes); i++ {
		curr := amortizedSearch(startingNodes[i], directions, 0, graph, make(map[string]int))
		// Get least common multiple:
		// Formula is a * b / greatestCommonDenominator(a, b)
		answer = answer * curr / gcd(answer, curr)
	}

	return strconv.Itoa(answer), nil
}

/*
 * HELPER
 */

type neighbor struct {
	left  string
	right string
}

func buildMap(lines []string) map[string]neighbor {
	graph := make(map[string]neighbor)
	for _, l := range lines {
		input := strings.Split(l, " = ")
		node := input[0]
		input[1] = strings.Trim(input[1], "()")
		neighborInput := strings.Split(input[1], ", ")
		n := neighbor{neighborInput[0], neighborInput[1]}
		graph[node] = n
	}
	return graph
}

func amortizedSearch(node string, directions string, idx int, graph map[string]neighbor, seen map[string]int) int {
	if steps, ok := seen[node]; ok {
		return steps + 1
	}

	if node[len(node)-1] == 'Z' {
		seen[node] = 0
		return seen[node]
	}

	var next string
	if directions[idx] == 'L' {
		next = graph[node].left
	} else {
		next = graph[node].right
	}
	nextIdx := (idx + 1) % len(directions)
	seen[node] = amortizedSearch(next, directions, nextIdx, graph, seen) + 1
	return seen[node]
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

package day05

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 5, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	seedsString := strings.Split(lines[0], ": ")
	seedsString = strings.Split(seedsString[1], " ")
	seeds := make([]int, 0, len(seedsString))
	for _, seedString := range seedsString {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	allMaps := make([][][]int, 0)
	currentMap := make([][]int, 0)
	for i := 3; i < len(lines); {
		line := lines[i]

		if len(line) == 0 || line[0] < '0' || line[0] > '9' {
			allMaps = append(allMaps, currentMap)
			currentMap = make([][]int, 0)
			i += 2
			continue
		}

		numsString := strings.Split(line, " ")
		nums := make([]int, 0, 3)
		for _, numString := range numsString {
			num, _ := strconv.Atoi(numString)
			nums = append(nums, num)
		}
		tmp := nums[0]
		nums[0] = nums[1]
		nums[1] = tmp
		currentMap = append(currentMap, nums)
		i++
	}
	allMaps = append(allMaps, currentMap)

	for _, currentMap := range allMaps {
		sort.SliceStable(currentMap, func(i, j int) bool {
			return currentMap[i][0] < currentMap[j][0]
		})
	}

	location := math.MaxInt
	for _, val := range seeds {
		for _, currentMap := range allMaps {
			index := findLastLessThanOrEqualTo(currentMap, 0, len(currentMap)-1, val)
			if index != -1 {
				entryMap := currentMap[index]
				if val < entryMap[0]+entryMap[2] {
					val = entryMap[1] + (val - entryMap[0])
				}
			}
		}
		if val < location {
			location = val
		}
	}

	return strconv.Itoa(location), nil
}

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 5, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	seedsString := strings.Split(lines[0], ": ")
	seedsString = strings.Split(seedsString[1], " ")
	seeds := make([]int, 0, len(seedsString))
	for _, seedString := range seedsString {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	allMaps := make([][][]int, 0)
	currentMap := make([][]int, 0)
	for i := 3; i < len(lines); {
		line := lines[i]

		if len(line) == 0 || line[0] < '0' || line[0] > '9' {
			allMaps = append(allMaps, currentMap)
			currentMap = make([][]int, 0)
			i += 2
			continue
		}

		numsString := strings.Split(line, " ")
		nums := make([]int, 0, 3)
		for _, numString := range numsString {
			num, _ := strconv.Atoi(numString)
			nums = append(nums, num)
		}
		tmp := nums[0]
		nums[0] = nums[1]
		nums[1] = tmp
		currentMap = append(currentMap, nums)
		i++
	}
	allMaps = append(allMaps, currentMap)

	for _, currentMap := range allMaps {
		sort.SliceStable(currentMap, func(i, j int) bool {
			return currentMap[i][0] < currentMap[j][0]
		})
	}

	location := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedRange := seeds[i+1]
		for i := 0; i < seedRange; i++ {
			val := seedStart + i
			for _, currentMap := range allMaps {
				index := findLastLessThanOrEqualTo(currentMap, 0, len(currentMap)-1, val)
				if index != -1 {
					entryMap := currentMap[index]
					if val < entryMap[0]+entryMap[2] {
						val = entryMap[1] + (val - entryMap[0])
					}
				}
			}
			if val < location {
				location = val
			}
		}
	}

	return strconv.Itoa(location), nil
}

func findLastLessThanOrEqualTo(arr [][]int, l int, r int, n int) int {
	if l == r {
		if arr[l][0] > n {
			return -1
		}
		return r
	}

	m := (l + r + 1) / 2
	if arr[m][0] <= n {
		return findLastLessThanOrEqualTo(arr, m, r, n)
	} else {
		return findLastLessThanOrEqualTo(arr, l, m-1, n)
	}
}

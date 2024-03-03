package day05

import (
	"errors"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

/*
 * PART ONE
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 5, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	// Get seeds
	seedsString := strings.Split(lines[0], ": ")
	seedsString = strings.Split(seedsString[1], " ")
	seeds := make([]int, 0, len(seedsString))
	for _, seedString := range seedsString {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	// Build mapping from seed to location
	mapping := buildMap(lines)

	// Find lowest location
	lowestLocation := math.MaxInt
	for _, s := range seeds {
		for _, e := range mapping {
			if s >= e.from.start && s <= e.from.end {
				offset := e.to.start - e.from.start
				location := s + offset
				lowestLocation = min(location, lowestLocation)
			}
		}
	}

	return strconv.Itoa(lowestLocation), nil
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

/*
 * PART TWO
 */

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 5, FileNumber: inputFileNumber}
	lines, err := utils.ReadInput(opts)
	if err != nil {
		return "", err
	}

	// Get seeds
	seedsString := strings.Split(lines[0], ": ")
	seedsString = strings.Split(seedsString[1], " ")
	seeds := make([]int, 0, len(seedsString))
	for _, seedString := range seedsString {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}
	seedRanges := []window{}
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, window{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	// Build mapping from seed to location
	mapping := buildMap(lines)

	// Sort mapping based on 'to' ranges
	sort.SliceStable(mapping, func(i, j int) bool {
		return mapping[i].to.start < mapping[j].to.start
	})

	for _, e := range mapping {
		lowestLocation := math.MaxInt
		for _, s := range seedRanges {
			// if there is an overlap
			if s.start <= e.from.end && e.from.start <= s.end {
				seed := max(s.start, e.from.start)
				offset := e.to.start - e.from.start
				location := seed + offset
				lowestLocation = min(lowestLocation, location)
			}
		}
		if lowestLocation != math.MaxInt {
			return strconv.Itoa(lowestLocation), nil
		}
	}

	// This behavior is NOT correct. It is possible that the answer
	// has not been shifted at all. I just assumed that the answer
	// has to be from a seed number whose location has been shifted,
	// which was the case. If seed - location has not been shifted,
	// then will have to search here instead of returning error.
	return "", errors.New("Not able to find seed with overlap")
}

/*
 * HELPER
 */

type window struct {
	start int
	end   int
}

type mapEntry struct {
	from window
	to   window
}

type listMap []mapEntry

func buildMap(lines []string) listMap {
	mapping := listMap{}
	i := len(lines) - 1
	for ; lines[i][0] >= '0' && lines[i][0] <= '9'; i-- {
		entry := getMapEntryFromLine(lines[i])
		mapping = append(mapping, entry)
	}

	// Sort mapping based on 'from' ranges
	sort.SliceStable(mapping, func(i, j int) bool {
		return mapping[i].from.start < mapping[j].from.start
	})

	// Start at the last entry of 'temperature-to-humidity map'
	i -= 2

	j := 0

	currentMap := listMap{}
	for i >= 2 {
		line := lines[i]
		if line[0] >= '0' && line[0] <= '9' {
			// add entry to map
			entry := getMapEntryFromLine(line)
			currentMap = append(currentMap, entry)

			// move up one line
			i--
			continue
		}

		// merge mappings
		mapping = mergeMaps(currentMap, mapping)

		// reset map for next iteration
		currentMap = listMap{}
		i -= 2
		j++
	}

	return mapping
}

func getMapEntryFromLine(line string) mapEntry {
	// Parse entry from string into list of numbers
	unparsedNumbers := strings.Split(line, " ")
	numbers := []int{}
	for _, n := range unparsedNumbers {
		num, _ := strconv.Atoi(n)
		numbers = append(numbers, num)
	}

	// Populate mapping entry
	fromEntry := window{numbers[1], numbers[1] + numbers[2] - 1}
	toEntry := window{numbers[0], numbers[0] + numbers[2] - 1}
	return mapEntry{fromEntry, toEntry}
}

func mergeMaps(fromMap, toMap listMap) listMap {
	// Sort fromMap based on 'to' ranges
	sort.SliceStable(fromMap, func(i, j int) bool {
		return fromMap[i].to.start < fromMap[j].to.start
	})

	// Sort toMap based on 'from' ranges
	sort.SliceStable(toMap, func(i, j int) bool {
		return toMap[i].from.start < toMap[j].from.start
	})

	// Combine intervals
	combinedWindows := []window{}
	for i := 0; i < len(fromMap); i++ {
		fromWindow := fromMap[i].to
		combinedWindows = append(combinedWindows, fromWindow)
	}
	for i := 0; i < len(toMap); i++ {
		toWindow := toMap[i].from
		combinedWindows = append(combinedWindows, toWindow)
	}
	slices.SortFunc(combinedWindows, func(i, j window) int {
		return i.start - j.start
	})

	// Initialize checkpoints and intervals
	a, b := combinedWindows[0].start, combinedWindows[0].end
	checkpointSet := map[int]bool{
		a:     true,
		b + 1: true,
	}
	intervals := []window{{a, b + 1}}
	checkpoints := []int{}

	// Populate intervals
	for _, w := range combinedWindows[1:] {
		a, b := w.start, w.end
		checkpointSet[a] = true
		checkpointSet[b+1] = true
		if a <= intervals[len(intervals)-1].end {
			last := len(intervals) - 1
			intervals[last].end = max(intervals[last].end, b+1)
		} else {
			intervals = append(intervals, window{a, b + 1})
		}
	}

	// Populate checkpoints
	for k := range checkpointSet {
		checkpoints = append(checkpoints, k)
	}
	slices.Sort(checkpoints)

	// Build non-overlapping intervals
	intermediateWindows := []window{}
	intervalIndex := 0
	lastCheckpoint := checkpoints[0]
	for _, checkpoint := range checkpoints[1:] {
		end := checkpoint - 1
		for end >= intervals[intervalIndex].end {
			intervalIndex++
		}
		if end >= intervals[intervalIndex].start {
			w := window{lastCheckpoint, end}
			intermediateWindows = append(intermediateWindows, w)
		}
		lastCheckpoint = checkpoint
	}

	mergedMap := listMap{}
	fromIndex, toIndex := 0, 0
	for _, w := range intermediateWindows {
		for fromIndex < len(fromMap) && w.end > fromMap[fromIndex].to.end {
			fromIndex++
		}
		for toIndex < len(toMap) && w.end > toMap[toIndex].from.end {
			toIndex++
		}

		var fromWindow window
		if fromIndex >= len(fromMap) || w.start < fromMap[fromIndex].to.start {
			// default 'from' mapping
			fromWindow = window{w.start, w.end}
		} else {
			offset := fromMap[fromIndex].to.start - fromMap[fromIndex].from.start
			fromWindow = window{w.start - offset, w.end - offset}
		}

		var toWindow window
		if toIndex >= len(toMap) || w.start < toMap[toIndex].from.start {
			// default 'to' mapping
			toWindow = window{w.start, w.end}
		} else {
			offset := toMap[toIndex].to.start - toMap[toIndex].from.start
			toWindow = window{w.start + offset, w.end + offset}
		}

		mergedMap = append(mergedMap, mapEntry{
			fromWindow,
			toWindow,
		})
	}

	finalMap := listMap{}
	for _, e := range mergedMap {
		from := e.from
		to := e.to
		if from.start == to.start && from.end == to.end {
			continue
		}
		finalMap = append(finalMap, e)
	}

	return finalMap
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

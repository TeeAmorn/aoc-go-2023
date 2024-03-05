package day07

import (
	"sort"
	"strconv"
	"strings"

	"github.com/TeeAmorn/aoc-go-2023/utils"
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

/*
 * PART 1
 */

func Part1(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 7, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	s := compute(lines, false)
	return strconv.Itoa(s), nil
}

/*
 * PART 2
 */

func Part2(inputFileNumber int) (string, error) {
	opts := utils.InputOptions{Day: 7, FileNumber: inputFileNumber}
	lines, _ := utils.ReadInput(opts)

	s := compute(lines, true)
	return strconv.Itoa(s), nil
}

/*
 * HELPER
 */

func compute(lines []string, wildcardJoker bool) int {

	// Get type of each hand
	hands := make([]hand, 0, len(lines))
	for _, line := range lines {
		inputs := strings.Split(line, " ")
		cards := inputs[0]
		bid, _ := strconv.Atoi(inputs[1])

		h := hand{cards, getHandValue(cards, wildcardJoker), getHandType(cards, wildcardJoker), bid}
		hands = append(hands, h)
	}

	// Sort hand based on win
	sort.SliceStable(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		if a.hand != b.hand {
			return a.hand < b.hand
		}
		return a.value < b.value
	})

	// Compute rank
	s := 0
	for i, h := range hands {
		s += h.bid * (i + 1)
	}

	return s
}

type hand struct {
	cards string
	value int
	hand  int
	bid   int
}

func getHandType(cards string, wildcardJoker bool) int {
	count := make(map[rune]int)
	for _, c := range cards {
		count[c] += 1
	}
	c := getHighestCount(count)

	switch len(count) {
	case 1:
		return FiveOfAKind
	case 2:
		if wildcardJoker && count['J'] > 0 {
			return FiveOfAKind
		}

		if c == 4 {
			// 4 1
			return FourOfAKind
		} else {
			// 3 2
			return FullHouse
		}
	case 3:
		if c == 3 {
			// 3 1 1
			if wildcardJoker && count['J'] > 0 {
				return FourOfAKind
			}
			return ThreeOfAKind
		} else {
			// 2 2 1
			if wildcardJoker && count['J'] > 0 {
				if count['J'] == 2 {
					return FourOfAKind
				} else {
					return FullHouse
				}
			}
			return TwoPair
		}
	case 4:
		// 2 1 1 1
		if wildcardJoker && count['J'] > 0 {
			return ThreeOfAKind
		}
		return OnePair
	default:
		if wildcardJoker && count['J'] > 0 {
			return OnePair
		}
		return HighCard
	}
}

func getHighestCount(count map[rune]int) int {
	var n int
	for _, v := range count {
		if v > n {
			n = v
		}
	}
	return n
}

func getHandValue(cards string, wildcardJoker bool) int {
	v := 0
	for _, c := range cards {
		v *= 0x10
		if c >= '2' && c <= '9' {
			v += int(c-'2') + 0x2
		} else if c == 'T' {
			v += int('9'-'2') + 0x3
		} else if c == 'J' {
			if wildcardJoker {
				v += 0x1
			} else {
				v += int('9'-'2') + 0x4
			}
		} else if c == 'Q' {
			v += int('9'-'2') + 0x5
		} else if c == 'K' {
			v += int('9'-'2') + 0x6
		} else {
			v += int('9'-'2') + 0x7
		}
	}
	return v
}

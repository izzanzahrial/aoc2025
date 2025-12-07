package day5

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part2() (int, error) {
	freshRangesInput, err := util.Read("day5", "fresh_ranges.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read fresh ranges input day 5 part 1: %w", err)
	}

	parsedFreshRanges := bytes.Split(freshRangesInput, []byte("\n"))

	var freshRanges []freshRange
	for _, line := range parsedFreshRanges {
		fr := bytes.Split(line, []byte("-"))
		lowerBound, err := strconv.Atoi(string(fr[0]))
		if err != nil {
			return 0, fmt.Errorf("failed to convert lower bound to int: %w", err)
		}

		upperBound, err := strconv.Atoi(string(fr[1]))
		if err != nil {
			return 0, fmt.Errorf("failed to convert upper bound to int: %w", err)
		}

		freshRanges = append(freshRanges, freshRange{
			lowerBound: lowerBound,
			upperBound: upperBound,
		})
	}

	sorter := sorter{}
	sorter.Sort(freshRanges)

	removeOverlapping(&sorter.ranges)

	var possibleFreshIngredients int
	for _, r := range sorter.ranges {
		possibleFreshIngredients += r.upperBound - r.lowerBound + 1
	}

	return possibleFreshIngredients, nil
}

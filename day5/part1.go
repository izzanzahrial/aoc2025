package day5

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"

	"github.com/izzanzahrial/aoc2025/util"
)

type freshRange struct {
	lowerBound int
	upperBound int
}

type sorter struct {
	ranges []freshRange
}

func (s *sorter) Len() int {
	return len(s.ranges)
}

func (s *sorter) Sort(ranges []freshRange) {
	s.ranges = ranges
	sort.Sort(s)
}

func (s *sorter) Swap(i, j int) {
	s.ranges[i], s.ranges[j] = s.ranges[j], s.ranges[i]
}

func (s *sorter) Less(i, j int) bool {
	return s.ranges[i].lowerBound < s.ranges[j].lowerBound
}

func Part1() (int, error) {
	freshRangesInput, err := util.Read("day5", "fresh_ranges.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read fresh ranges input day 5 part 1: %w", err)
	}

	ingredientsInput, err := util.Read("day5", "ingredients.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read ingredients input day 5 part 1: %w", err)
	}

	parsedFreshRanges := bytes.Split(freshRangesInput, []byte("\n"))
	parsedIngredients := bytes.Split(ingredientsInput, []byte("\n"))

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

	var freshIngredients int
	for _, ingredient := range parsedIngredients {
		intIngredient, err := strconv.Atoi(string(ingredient))
		if err != nil {
			return 0, fmt.Errorf("failed to convert ingredient to int: %w", err)
		}

		if isFresh(intIngredient, sorter.ranges) {
			freshIngredients++
		}
	}

	return freshIngredients, nil
}

func isFresh(ingredient int, freshRanges []freshRange) bool {
	left := 0
	right := len(freshRanges) - 1

	for left <= right {
		mid := (left + right) / 2
		if ingredient >= freshRanges[mid].lowerBound && ingredient <= freshRanges[mid].upperBound {
			return true
		}

		if ingredient < freshRanges[mid].lowerBound {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func removeOverlapping(ranges *[]freshRange) {
	for i := 1; i < len(*ranges); i++ {
		// check if the upper bound of the prev value is greater than the lower bound of the current value
		if (*ranges)[i-1].upperBound >= (*ranges)[i].lowerBound {
			// then check if the upper bound of the current value is greater than the upper bound of the prev value
			// meaning that the current value is overlapping with the previous value
			// 10-14 and 12-18 overlaps, but 10-20 and 12-18 doesn't
			if (*ranges)[i].upperBound > (*ranges)[i-1].upperBound {
				(*ranges)[i-1].upperBound = (*ranges)[i].upperBound
			}
			*ranges = append((*ranges)[:i], (*ranges)[i+1:]...)
			i--
		}
	}
}

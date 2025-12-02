package day1

import (
	"fmt"
	"strings"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part1() (int, error) {
	b, err := util.Read("day1", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("day1 part1 failed to read input: %w", err)
	}

	strs := strings.Split(string(b), "\n")

	var result int
	startDial := 50
	for _, str := range strs {
		poiting, distance, err := parse(str)
		if err != nil {
			return 0, fmt.Errorf("day1 part1 failed to parse: %w", err)
		}

		if poiting == "left" {
			if startDial-distance < 0 {
				startDial = (startDial - distance) % 100
			} else {
				startDial -= distance
			}
		} else {
			if startDial+distance > 99 {
				startDial = (startDial + distance) % 100
			} else {
				startDial += distance
			}
		}

		if startDial == 0 {
			result++
		}
	}

	return result, nil
}

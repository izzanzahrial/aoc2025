package day1

import (
	"fmt"
	"strings"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part2() (int, error) {
	b, err := util.Read("day1", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("day1 part2 failed to read input: %w", err)
	}

	strs := strings.Split(string(b), "\n")

	var result int
	startDial := 50
	for _, str := range strs {
		poiting, distance, err := parse(str)
		if err != nil {
			return 0, fmt.Errorf("day1 part2 failed to parse: %w", err)
		}

		if poiting == "left" {
			if startDial == 0 {
				result += distance / 100
			} else if distance >= startDial {
				result += 1 + (distance-startDial)/100
			}
			startDial = ((startDial-distance)%100 + 100) % 100
		} else {
			result += (startDial + distance) / 100
			startDial = (startDial + distance) % 100
		}
	}

	return result, nil
}

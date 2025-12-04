package day3

import (
	"bytes"
	"fmt"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part1() (int, error) {
	input, err := util.Read("day3", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input day 3 part 1: %w", err)
	}

	var result int
	parsedInput := bytes.Split(input, []byte("\n"))
	for _, line := range parsedInput {
		var curr [2]byte
		for i := 0; i < len(line)-1; i++ {
			if curr[0] == '0' {
				curr[0] = line[i]
			}

			if line[i] > curr[0] {
				curr[0] = line[i]
				curr[1] = 0
			}

			for j := i + 1; j < len(line); j++ {
				if curr[1] == '0' {
					curr[1] = line[j]
				}

				if line[j] > curr[1] {
					curr[1] = line[j]
				}
			}
		}

		digit1 := int(curr[0] - '0')
		digit2 := int(curr[1] - '0')
		result += (digit1 * 10) + digit2
	}

	return result, nil
}

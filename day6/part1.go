package day6

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part1() (int, error) {
	input, err := util.Read("day6", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input day 6 part 1: %w", err)
	}

	parsedInput := bytes.Split(input, []byte("\n"))
	var mathProblems [][]string
	for _, line := range parsedInput {
		mathProblems = append(mathProblems, strings.Fields(string(line)))
	}

	var result int
	for i := 0; i < len(mathProblems[0]); i++ {
		num1, err := strconv.Atoi(mathProblems[0][i])
		if err != nil {
			return 0, fmt.Errorf("failed to convert first number to int: %w", err)
		}

		num2, err := strconv.Atoi(mathProblems[1][i])
		if err != nil {
			return 0, fmt.Errorf("failed to convert second number to int: %w", err)
		}

		num3, err := strconv.Atoi(mathProblems[2][i])
		if err != nil {
			return 0, fmt.Errorf("failed to convert third number to int: %w", err)
		}

		num4, err := strconv.Atoi(mathProblems[3][i])
		if err != nil {
			return 0, fmt.Errorf("failed to convert fourth number to int: %w", err)
		}

		if mathProblems[4][i] == "+" {
			result += num1 + num2 + num3 + num4
		} else {
			result += num1 * num2 * num3 * num4
		}
	}

	return result, nil
}

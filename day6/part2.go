package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part2() (int, error) {
	input, err := util.Read("day6", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input day 6 part 1: %w", err)
	}

	parsedInput := strings.Split(string(input), "\n")
	var result int
	var currProblem []int
	var currOp string
	for i := len(parsedInput[0]) - 1; i >= 0; i-- {
		var currNum strings.Builder
		isBlank := true
		for j := 0; j < len(parsedInput); j++ {
			char := parsedInput[j][i]
			if char == ' ' {
				continue
			}

			isBlank = false
			if char == '+' || char == '*' {
				currOp = string(char)
				continue
			}

			if err := currNum.WriteByte(char); err != nil {
				return 0, fmt.Errorf("failed to write to currNum: %w", err)
			}
		}

		// Column is separator, finish current problem
		if isBlank {
			if len(currProblem) > 0 {
				if currOp == "*" {
					result += multiplication(currProblem...)
				} else {
					result += addition(currProblem...)
				}
			}

			// reset state
			currProblem = nil
			currOp = ""
			continue
		}

		strNum := strings.Trim(currNum.String(), " ")
		if len(strNum) > 0 {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return 0, fmt.Errorf("failed to convert currNum to int: %w, currNum: %s", err, strNum)
			}
			currProblem = append(currProblem, num)
		}
	}

	if len(currProblem) > 0 {
		if currOp == "*" {
			result += multiplication(currProblem...)
		} else {
			result += addition(currProblem...)
		}
	}

	return result, nil
}

func addition(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func multiplication(nums ...int) int {
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}

package day4

import (
	"bytes"
	"fmt"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part2() (int, error) {
	input, err := util.Read("day4", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input day 4 part 2: %w", err)
	}

	parsedInput := bytes.Split(input, []byte("\n"))

	canBeAccessed := removeAble(parsedInput)

	return canBeAccessed, nil
}

func removeAble(parsedInput [][]byte) int {
	row := len(parsedInput)
	col := len(parsedInput[0])

	var canBeAccessed int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if parsedInput[i][j] != '@' {
				continue
			}

			var rollPaper int
			// check top
			if i > 0 {
				if parsedInput[i-1][j] == '@' {
					rollPaper++
				}
			}

			// check top right
			if i > 0 && j < col-1 {
				if parsedInput[i-1][j+1] == '@' {
					rollPaper++
				}
			}

			//check right
			if j < col-1 {
				if parsedInput[i][j+1] == '@' {
					rollPaper++
				}
			}

			// check bottom right
			if i < row-1 && j < col-1 {
				if parsedInput[i+1][j+1] == '@' {
					rollPaper++
				}
			}

			//check bottom
			if i < row-1 {
				if parsedInput[i+1][j] == '@' {
					rollPaper++
				}
			}

			// check bottom left
			if i < row-1 && j > 0 {
				if parsedInput[i+1][j-1] == '@' {
					rollPaper++
				}
			}

			// check left
			if j > 0 {
				if parsedInput[i][j-1] == '@' {
					rollPaper++
				}
			}

			// check top left
			if i > 0 && j > 0 {
				if parsedInput[i-1][j-1] == '@' {
					rollPaper++
				}
			}

			if rollPaper < 4 {
				canBeAccessed++
				parsedInput[i][j] = 'x'
			}
		}
	}

	if canBeAccessed == 0 {
		return 0
	}

	return canBeAccessed + removeAble(parsedInput)
}

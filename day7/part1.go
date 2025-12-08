package day7

import (
	"bytes"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part1() (int, error) {
	input, err := util.Read("day7", "input.txt")
	if err != nil {
		return 0, err
	}

	var totalSplits int
	parsedInput := bytes.Split(input, []byte("\n"))
	row, col := len(parsedInput)-1, len(parsedInput[0])-1
	for i := 0; i <= row; i++ {
		for j := 0; j <= col; j++ {
			char := parsedInput[i][j]
			switch char {
			case 'S':
				if inBounds(i+1, j, row, col) {
					parsedInput[i+1][j] = '|'
				}
			case '|':
				down := inBounds(i+1, j, row, col)
				right := inBounds(i+1, j+1, row, col)
				left := inBounds(i+1, j-1, row, col)

				if down && right {
					if parsedInput[i+1][j] == '^' {
						if left {
							parsedInput[i+1][j-1] = '|'
						}
						if right {
							parsedInput[i+1][j+1] = '|'
						}
						totalSplits++
						continue
					}
				}

				if down {
					parsedInput[i+1][j] = '|'
				}
			}
		}
	}

	return totalSplits, nil
}

func inBounds(i, j, row, col int) bool {
	return i >= 0 && i <= row && j >= 0 && j <= col
}

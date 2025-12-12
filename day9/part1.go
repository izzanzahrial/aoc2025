package day9

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part1() (int, error) {
	input, err := util.Read("day9", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input day 9 part 1: %w", err)
	}

	parsedInput := bytes.Split(input, []byte("\n"))
	var largest int
	for i := 0; i < len(parsedInput)-1; i++ {
		nodeA := bytes.Split(parsedInput[i], []byte(","))
		rowA, err := strconv.Atoi(string(nodeA[0]))
		if err != nil {
			return 0, fmt.Errorf("failed to parse row A: %w", err)
		}
		colA, err := strconv.Atoi(string(nodeA[1]))
		if err != nil {
			return 0, fmt.Errorf("failed to parse col A: %w", err)
		}
		for j := i + 1; j < len(parsedInput); j++ {
			nodeB := bytes.Split(parsedInput[j], []byte(","))
			rowB, err := strconv.Atoi(string(nodeB[0]))
			if err != nil {
				return 0, fmt.Errorf("failed to parse row B: %w", err)
			}
			colB, err := strconv.Atoi(string(nodeB[1]))
			if err != nil {
				return 0, fmt.Errorf("failed to parse col B: %w", err)
			}

			totalTiles := TotalTiles(rowA, colA, rowB, colB)
			largest = max(totalTiles, largest)
		}
	}

	return largest, nil
}

func TotalTiles(rowA, colA, rowB, colB int) int {
	totalRow := int(math.Abs(float64(rowA-rowB))) + 1
	totalCol := int(math.Abs(float64(colA-colB))) + 1
	return totalRow * totalCol
}

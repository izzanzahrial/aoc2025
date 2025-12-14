package day11

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part1() (int, error) {
	input, err := util.Read("day11", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read input day 11 part 1: %w", err)
	}

	parsedInput := bytes.Split(input, []byte("\n"))
	rackMap := make(map[string][]string, len(parsedInput))
	var totalPaths int
	var startingCodes []string
	for _, line := range parsedInput {
		rackCode := strings.Split(string(line), ":")
		from := rackCode[0]
		to := strings.Split(rackCode[1], " ")

		if from == "you" {
			startingCodes = to
		}
		rackMap[from] = to
	}

	visited := make(map[string]int, len(startingCodes))
	for _, code := range startingCodes {
		totalPaths += solve(code, visited, rackMap)
	}

	return totalPaths, nil
}

func solve(code string, visited map[string]int, rackMap map[string][]string) int {
	if code == "out" {
		return 1
	}

	if v, ok := visited[string(code)]; ok {
		return v
	}

	var totalPath int
	for _, c := range rackMap[code] {
		totalPath += solve(c, visited, rackMap)
	}

	visited[code] = totalPath
	return totalPath
}

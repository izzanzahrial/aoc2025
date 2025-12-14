package day11

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/izzanzahrial/aoc2025/util"
)

type state struct {
	node string
	fft  bool
	dac  bool
}

func Part2() (int, error) {
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

		if from == "svr" {
			startingCodes = to
		}
		rackMap[from] = to
	}

	visited := make(map[state]int, len(startingCodes))
	for _, code := range startingCodes {
		totalPaths += solve2(code, false, false, visited, rackMap)
	}

	return totalPaths, nil
}

func solve2(code string, hasFFT, hasDAC bool, visited map[state]int, rackMap map[string][]string) int {
	if code == "out" {
		if hasFFT && hasDAC {
			return 1
		}
		return 0
	}

	key := state{code, hasFFT, hasDAC}
	if v, ok := visited[key]; ok {
		return v
	}

	var totalPath int
	for _, c := range rackMap[code] {
		nFFT := hasFFT || c == "fft"
		nDAC := hasDAC || c == "dac"
		totalPath += solve2(c, nFFT, nDAC, visited, rackMap)
	}

	visited[key] = totalPath
	return totalPath
}

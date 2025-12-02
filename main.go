package main

import (
	"fmt"
	"log/slog"

	"github.com/izzanzahrial/aoc2025/day1"
)

func main() {
	day1Part1, err := day1.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	day1Part2, err := day1.Part2()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 1 Part 1 result: %d", day1Part1))
	slog.Info(fmt.Sprintf("Day 1 Part 2 result: %d", day1Part2))
}

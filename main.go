package main

import (
	"fmt"
	"log/slog"

	"github.com/izzanzahrial/aoc2025/day1"
	"github.com/izzanzahrial/aoc2025/day2"
	"github.com/izzanzahrial/aoc2025/day3"
	"github.com/izzanzahrial/aoc2025/day4"
	"github.com/izzanzahrial/aoc2025/day5"
	"github.com/izzanzahrial/aoc2025/day6"
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

	day2Part1, err := day2.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	day2Part2, err := day2.Part2()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 2 Part 1 result: %d", day2Part1))
	slog.Info(fmt.Sprintf("Day 2 Part 2 result: %d", day2Part2))

	day3Part1, err := day3.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 3 Part 1 result: %d", day3Part1))

	day4Part1, err := day4.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	day4Part2, err := day4.Part2()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 4 Part 1 result: %d", day4Part1))
	slog.Info(fmt.Sprintf("Day 4 Part 2 result: %d", day4Part2))

	day5Part1, err := day5.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	day5Part2, err := day5.Part2()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 5 Part 1 result: %d", day5Part1))
	slog.Info(fmt.Sprintf("Day 5 Part 2 result: %d", day5Part2))

	day6Part1, err := day6.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	day6Part2, err := day6.Part2()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 6 Part 1 result: %d", day6Part1))
	slog.Info(fmt.Sprintf("Day 6 Part 2 result: %d", day6Part2))
}

package main

import (
	"fmt"
	"log/slog"

	"github.com/izzanzahrial/aoc2025/day1"
	"github.com/izzanzahrial/aoc2025/day11"
	"github.com/izzanzahrial/aoc2025/day2"
	"github.com/izzanzahrial/aoc2025/day3"
	"github.com/izzanzahrial/aoc2025/day4"
	"github.com/izzanzahrial/aoc2025/day5"
	"github.com/izzanzahrial/aoc2025/day6"
	"github.com/izzanzahrial/aoc2025/day7"
	"github.com/izzanzahrial/aoc2025/day8"
	"github.com/izzanzahrial/aoc2025/day9"
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

	day7Part1, err := day7.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 7 Part 1 result: %d", day7Part1))

	day8Part1, err := day8.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	day8Part2, err := day8.Part2()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 8 Part 1 result: %d", day8Part1))
	slog.Info(fmt.Sprintf("Day 8 Part 2 result: %d", day8Part2))

	day9Part1, err := day9.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 9 Part 1 result: %d", day9Part1))

	day11Part1, err := day11.Part1()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info(fmt.Sprintf("Day 11 Part 1 result: %d", day11Part1))
}

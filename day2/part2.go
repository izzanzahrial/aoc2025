package day2

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part2() (int, error) {
	b, err := util.Read("day2", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("day2 part2 failed to read input: %w", err)
	}

	input, err := parse(string(b))
	if err != nil {
		return 0, fmt.Errorf("day2 part2 failed to parse: %w", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var result int
	for _, num := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var curr int
			for i := num[0]; i <= num[1]; i++ {
				numStr := strconv.Itoa(i)

				if invalid(numStr) {
					curr += i
				}
			}

			mu.Lock()
			result += curr
			mu.Unlock()
		}()
	}

	wg.Wait()

	return result, nil
}

func invalid(str string) bool {
	n := len(str)

	// We only need to check pattern lengths up to half the string.
	// (e.g., if string is length 10, max pattern is length 5 repeated twice).
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		// if the length is not mod 0 of the patternLen we skip it
		if n%patternLen != 0 {
			continue
		}

		pattern := str[:patternLen]
		isMatch := true

		// Check every chunk of size patternLen against the pattern
		for i := patternLen; i < n; i += patternLen {
			curr := str[i : i+patternLen]
			if curr != pattern {
				isMatch = false
				break
			}
		}

		// If the inner loop finished without setting isMatch to false, we found it.
		if isMatch {
			return true
		}
	}

	return false
}

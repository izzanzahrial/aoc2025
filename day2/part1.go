package day2

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/izzanzahrial/aoc2025/util"
)

func Part1() (int, error) {
	b, err := util.Read("day2", "input.txt")
	if err != nil {
		return 0, fmt.Errorf("day2 part1 failed to read input: %w", err)
	}

	input, err := parse(string(b))
	if err != nil {
		return 0, fmt.Errorf("day2 part1 failed to parse: %w", err)
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

				if len(numStr)%2 != 0 {
					continue
				}

				mid := len(numStr) / 2
				firstHalf := numStr[:mid]
				secondHalf := numStr[mid:]

				if firstHalf == secondHalf {
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

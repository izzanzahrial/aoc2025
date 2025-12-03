package day2

import (
	"strconv"
	"strings"
)

func parse(str string) ([][2]int, error) {
	var result [][2]int
	for _, strNum := range strings.Split(str, ",") {
		var nums [2]int
		for i, n := range strings.Split(strNum, "-") {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			nums[i] = num
		}
		result = append(result, nums)
	}

	return result, nil
}

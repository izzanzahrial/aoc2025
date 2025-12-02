package day1

import "strconv"

func parse(str string) (string, int, error) {
	var poiting string
	var distance int
	if str[0] == 'L' {
		poiting = "left"
	} else {
		poiting = "right"
	}

	distance, err := strconv.Atoi(str[1:])
	if err != nil {
		return "", 0, err
	}

	return poiting, distance, nil
}

package day5

func isFresh(ingredient int, freshRanges []freshRange) bool {
	left := 0
	right := len(freshRanges) - 1

	for left <= right {
		mid := (left + right) / 2
		if ingredient >= freshRanges[mid].lowerBound && ingredient <= freshRanges[mid].upperBound {
			return true
		}

		if ingredient < freshRanges[mid].lowerBound {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func removeOverlapping(ranges *[]freshRange) {
	for i := 1; i < len(*ranges); i++ {
		// check if the upper bound of the prev value is greater than the lower bound of the current value
		if (*ranges)[i-1].upperBound >= (*ranges)[i].lowerBound {
			// then check if the upper bound of the current value is greater than the upper bound of the prev value
			// meaning that the current value is overlapping with the previous value
			// 10-14 and 12-18 overlaps, but 10-20 and 12-18 doesn't
			if (*ranges)[i].upperBound > (*ranges)[i-1].upperBound {
				(*ranges)[i-1].upperBound = (*ranges)[i].upperBound
			}
			*ranges = append((*ranges)[:i], (*ranges)[i+1:]...)
			i--
		}
	}
}

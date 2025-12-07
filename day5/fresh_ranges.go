package day5

import "sort"

type freshRange struct {
	lowerBound int
	upperBound int
}

type sorter struct {
	ranges []freshRange
}

func (s *sorter) Len() int {
	return len(s.ranges)
}

func (s *sorter) Sort(ranges []freshRange) {
	s.ranges = ranges
	sort.Sort(s)
}

func (s *sorter) Swap(i, j int) {
	s.ranges[i], s.ranges[j] = s.ranges[j], s.ranges[i]
}

func (s *sorter) Less(i, j int) bool {
	return s.ranges[i].lowerBound < s.ranges[j].lowerBound
}

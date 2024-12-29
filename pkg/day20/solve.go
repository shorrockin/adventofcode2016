package day20

import (
	"adventofcode2016/pkg/utils"
	slice "adventofcode2016/pkg/utils/slices"
	"slices"
	"strings"
)

const MAX_VALUE = 4294967295

type Range [2]int

func PartOne(path string) int {
	return parse(path)[0][0]
}

func PartTwo(path string) int {
	return slice.Reduce(parse(path), 0, func(acc int, r Range) int {
		return acc + r[1] - r[0] + 1
	})
}

func parse(path string) []Range {
	ranges := []Range{{0, MAX_VALUE}}

	for _, line := range utils.MustReadInput(path) {
		parts := strings.Split(line, "-")
		blocked := [2]int{utils.MustAtoi(parts[0]), utils.MustAtoi(parts[1])}

		for idx := 0; idx < len(ranges); idx++ {
			allowed := ranges[idx]

			// there are 4 scenarios where we need to modify the allowed list, it either:
			// 1. overlaps completely, and elimanets this allowed range
			// 2. overlaps within, and splits the allowed range in two
			// 3. overlaps the start, and shifts the start of the allowed range
			// 4. overlaps the end, and shifts the end of the allowed range
			if blocked[0] <= allowed[0] && blocked[1] >= allowed[1] {
				ranges = append(ranges[:idx], ranges[idx+1:]...)
			} else if blocked[0] > allowed[0] && blocked[1] < allowed[1] {
				ranges = append(ranges, Range{blocked[1] + 1, allowed[1]})
				ranges[idx][1] = blocked[0] - 1
			} else if blocked[0] <= allowed[0] && blocked[1] >= allowed[0] && blocked[1] < allowed[1] {
				ranges[idx][0] = blocked[1] + 1
			} else if blocked[0] > allowed[0] && blocked[0] <= allowed[1] && blocked[1] >= allowed[1] {
				ranges[idx][1] = blocked[0] - 1
			}
		}
	}

	slices.SortFunc(ranges, func(left, right Range) int {
		return left[0] - right[0]
	})

	return ranges
}

package day09

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"strings"
)

type Part int

const (
	PartOne Part = iota
	PartTwo
)

func Solve(path string, part Part) int {
	lines := utils.MustReadInput(path)
	assert.Equal(1, len(lines), "incorrect number of input lines")
	return Decompress(lines[0], part)
}

func Decompress(input string, part Part) int {
	count := 0

	for idx := 0; idx < len(input); idx++ {
		char := input[idx]
		if char == '(' {
			end := strings.Index(input[idx:], ")")
			assert.NotEqual(-1, end, "could not find end of marker", "idx", idx, "input", input)

			marker := input[idx+1 : idx+end]
			parts := strings.Split(marker, "x")
			assert.Equal(2, len(parts), "incorrect number of parts in marker", "parts", parts, "marker", marker)

			repeat := utils.MustAtoi(parts[1])
			length := utils.MustAtoi(parts[0])

			if part == PartOne {
				count += repeat * length
				idx += end + length
			} else {
				count += repeat * Decompress(input[idx+end+1:idx+end+1+length], PartTwo)
				idx += end + length
			}
		} else {
			count += 1
		}
	}

	return count
}

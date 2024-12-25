package day03

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/slices"
	"strings"
)

func PartOne(path string) int {
	lines := utils.MustReadInput(path)
	values := slices.Map(lines, func(line string) []int {
		fields := strings.Fields(line)
		return slices.Map(fields, utils.MustAtoi)
	})
	return len(slices.Filter(values, IsValidTriangle))
}

func PartTwo(path string) int {
	lines := utils.MustReadInput(path)
	var values [][]int

	for i := 0; i < len(lines); i = i + 3 {
		first := strings.Fields(lines[i])
		second := strings.Fields(lines[i+1])
		third := strings.Fields(lines[i+2])

		values = append(values, []int{utils.MustAtoi(first[0]), utils.MustAtoi(second[0]), utils.MustAtoi(third[0])})
		values = append(values, []int{utils.MustAtoi(first[1]), utils.MustAtoi(second[1]), utils.MustAtoi(third[1])})
		values = append(values, []int{utils.MustAtoi(first[2]), utils.MustAtoi(second[2]), utils.MustAtoi(third[2])})
	}
	return len(slices.Filter(values, IsValidTriangle))
}

func IsValidTriangle(values []int) bool {
	assert.Assert(3 == len(values), "values must contain 3 ints")
	a := values[0]
	b := values[1]
	c := values[2]
	return a+b > c && a+c > b && b+c > a
}

package day02

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/grid"
	"adventofcode2016/pkg/utils"
	"strings"
)

var PART_ONE_GRID = [][]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

var PART_TWO_GRID = [][]string{
	{" ", " ", "1", " ", " "},
	{" ", "2", "3", "4", " "},
	{"5", "6", "7", "8", "9"},
	{" ", "A", "B", "C", " "},
	{" ", " ", "R", " ", " "},
}

func PartOne(path string) int {
	return utils.MustAtoi(processGrid(path, PART_ONE_GRID, grid.Coordinate{X: 1, Y: 1}))
}

func PartTwo(path string) string {
	return processGrid(path, PART_TWO_GRID, grid.Coordinate{X: 0, Y: 2})
}

func processGrid(path string, grid [][]string, position grid.Coordinate) string {
	lines := utils.MustReadInput(path)
	var code strings.Builder

	for _, line := range lines {
		for _, character := range line {
			next := position
			switch character {
			case 'U':
				next = position.North()
			case 'D':
				next = position.South()
			case 'L':
				next = position.West()
			case 'R':
				next = position.East()
			default:
				assert.Fail("unexpected character in line", character)
			}

			if next.X < 0 || next.Y < 0 || next.X == len(grid[0]) || next.Y == len(grid) {
				continue
			}
			if grid[next.Y][next.X] == " " {
				continue
			}

			position = next
		}
		code.WriteString(grid[position.Y][position.X])

	}
	return code.String()
}

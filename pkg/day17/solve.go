package day17

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/astar"
	"adventofcode2016/pkg/utils/collections"
	"adventofcode2016/pkg/utils/grid"
	"adventofcode2016/pkg/utils/logger"
	"adventofcode2016/pkg/utils/slices"
	"fmt"
)

type Position struct {
	coords   grid.Coord
	passcode string
	hash     string
}

var log = logger.New("day17")
var end = grid.At(3, 3)

func PartOne(passcode string) string {
	log.Reset(fmt.Sprintf("part one: %s", passcode))
	tail := slices.Tail(path(passcode, true))
	return logger.LogReturn(&log, tail.passcode[len(passcode):])
}

func PartTwo(passcode string) int {
	log.Reset(fmt.Sprintf("part two: %s", passcode))
	solution := path(passcode, false)
	return logger.LogReturn(&log, len(solution)-1)
}

func path(passcode string, shortest bool) []Position {
	path := astar.AStar(
		Position{grid.At(0, 0), passcode, utils.Md5(passcode)},
		isComplete,
		neighbors,
		utils.Ternary(shortest, shortestHeuristic, longestHeuristic),
	)
	assert.NotEqual(0, len(path), "no path found")
	return path
}

func shortestHeuristic(node Position, from *collections.PqNode[Position]) float64 {
	return float64(node.coords.Distance(end))
}

func longestHeuristic(node Position, from *collections.PqNode[Position]) float64 {
	// if we're moving to the last possible position, we want to provide a heuristic
	// such that the longest path will be prioritized, we can use the passcode length
	// as a proxy for path length
	if node.coords == end {
		// TODO: can't use math.MaxFloat64 here, not sure why
		return 10000 - float64(len(node.passcode))
	}
	return 0
}

func neighbors(node *collections.PqNode[Position]) []Position {
	position := node.Contents

	directions := make([]grid.Coord, 0, 4)
	if isOpen(position.hash[0]) {
		directions = append(directions, grid.North.WithLabel("U"))
	}
	if isOpen(position.hash[1]) {
		directions = append(directions, grid.South.WithLabel("D"))
	}
	if isOpen(position.hash[2]) {
		directions = append(directions, grid.West.WithLabel("L"))
	}
	if isOpen(position.hash[3]) {
		directions = append(directions, grid.East.WithLabel("R"))
	}

	directions = slices.Filter(directions, func(to grid.Coord) bool {
		np := position.coords.Offset(to)
		return np.X >= 0 && np.Y >= 0 && np.X <= end.X && np.Y <= end.Y
	})

	return slices.Map(directions, func(to grid.Coord) Position {
		return Position{
			position.coords.Offset(to),
			position.passcode + to.Label,
			utils.Md5(position.passcode + to.Label),
		}
	})
}

func isComplete(node *collections.PqNode[Position]) bool {
	return node.Contents.coords == end
}

func isOpen(char byte) bool {
	return char >= 'b' && char <= 'f'
}

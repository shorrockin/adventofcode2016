package day13

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/graph"
	"adventofcode2016/pkg/grid"
	"adventofcode2016/pkg/utils"
	"math/bits"
)

func StepsToPosition(seed int, starting grid.Coordinate, target grid.Coordinate) int {
	maze := NewMaze(seed)
	path, ok := graph.BFS(
		starting,
		func(at grid.Coordinate) []grid.Coordinate { return maze.Neighbors(at) },
		func(at grid.Coordinate) bool { return at == target },
	)

	if !ok {
		assert.Fail("could not find path to position")
	}

	return len(path) - 1
}

// effectively a copy of our BFS algorithm which limits itself to a path
// size of 50, returning the number of unique positions visited
func CountUniquePositions(seed int, starting grid.Coordinate, steps int) int {
	maze := NewMaze(seed)

	queue := utils.NewQueue[[]grid.Coordinate]()
	queue.Enqueue([]grid.Coordinate{starting})

	visited := utils.NewSet[grid.Coordinate]()
	visited.Add(starting)

	for !queue.IsEmpty() {
		currentPath, ok := queue.Dequeue()
		if !ok {
			assert.Fail("failed to dequeue, expected value to be on queue")
		}

		if len(currentPath) > steps {
			continue
		}

		tail := currentPath[len(currentPath)-1]

		for _, neighbor := range maze.Neighbors(tail) {
			if !visited.Contains(neighbor) {
				var newPath []grid.Coordinate = make([]grid.Coordinate, len(currentPath)+1)
				copy(newPath, currentPath)
				newPath[len(currentPath)] = neighbor

				queue.Enqueue(newPath)
				visited.Add(neighbor)
			}
		}
	}

	return len(visited)
}

type Maze struct {
	seed  int
	cache map[grid.Coordinate]bool
}

func NewMaze(seed int) Maze {
	return Maze{seed, make(map[grid.Coordinate]bool)}
}

func (maze *Maze) Neighbors(at grid.Coordinate) []grid.Coordinate {
	var neighbors []grid.Coordinate
	for _, cardinal := range at.Cardinals() {
		if maze.IsOpen(cardinal) {
			neighbors = append(neighbors, cardinal)
		}
	}
	return neighbors
}

func (maze *Maze) IsOpen(pos grid.Coordinate) bool {
	if open, ok := maze.cache[pos]; ok {
		return open
	}

	x := pos.X
	y := pos.Y

	if x < 0 || y < 0 {
		return false
	}

	n := x*x + 3*x + 2*x*y + y + y*y + maze.seed
	open := bits.OnesCount(uint(n))%2 == 0
	maze.cache[pos] = open
	return open
}

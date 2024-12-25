package day13

import (
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/bfs"
	"adventofcode2016/pkg/utils/collections"
	"adventofcode2016/pkg/utils/grid"
	"math/bits"
)

func StepsToPosition(seed int, starting grid.Coord, target grid.Coord) int {
	maze := NewMaze(seed)
	path, ok := bfs.BFS(
		starting,
		func(at grid.Coord) []grid.Coord { return maze.Neighbors(at) },
		func(at grid.Coord) bool { return at == target },
	)

	if !ok {
		assert.Fail("could not find path to position")
	}

	return len(path) - 1
}

// effectively a copy of our BFS algorithm which limits itself to a path
// size of 50, returning the number of unique positions visited
func CountUniquePositions(seed int, starting grid.Coord, steps int) int {
	maze := NewMaze(seed)

	queue := collections.NewQueue[[]grid.Coord]()
	queue.Push([]grid.Coord{starting})

	visited := collections.NewSet[grid.Coord]()
	visited.Add(starting)

	for !queue.IsEmpty() {
		currentPath, ok := queue.Pop()
		if !ok {
			assert.Fail("failed to dequeue, expected value to be on queue")
		}

		if len(currentPath) > steps {
			continue
		}

		tail := currentPath[len(currentPath)-1]

		for _, neighbor := range maze.Neighbors(tail) {
			if !visited.Contains(neighbor) {
				var newPath []grid.Coord = make([]grid.Coord, len(currentPath)+1)
				copy(newPath, currentPath)
				newPath[len(currentPath)] = neighbor

				queue.Push(newPath)
				visited.Add(neighbor)
			}
		}
	}

	return len(visited)
}

type Maze struct {
	seed  int
	cache map[grid.Coord]bool
}

func NewMaze(seed int) Maze {
	return Maze{seed, make(map[grid.Coord]bool)}
}

func (maze *Maze) Neighbors(at grid.Coord) []grid.Coord {
	var neighbors []grid.Coord
	for _, cardinal := range at.Cardinals() {
		if maze.IsOpen(cardinal) {
			neighbors = append(neighbors, cardinal)
		}
	}
	return neighbors
}

func (maze *Maze) IsOpen(pos grid.Coord) bool {
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

package bfs

import (
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/collections"
)

func BFS[T comparable](start T, neighbors func(from T) []T, complete func(at T) bool) ([]T, bool) {
	queue := collections.NewQueue[[]T]()
	queue.Push([]T{start})

	visited := collections.NewSet[T]()
	visited.Add(start)

	for !queue.IsEmpty() {
		currentPath, ok := queue.Pop()
		if !ok {
			assert.Fail("failed to dequeue, expected value to be on queue")
		}

		tail := currentPath[len(currentPath)-1]

		if complete(tail) {
			return currentPath, true
		}

		for _, neighbor := range neighbors(tail) {
			if !visited.Contains(neighbor) {
				var newPath []T = make([]T, len(currentPath)+1)
				copy(newPath, currentPath)
				newPath[len(currentPath)] = neighbor

				queue.Push(newPath)
				visited.Add(neighbor)
			}
		}
	}

	return make([]T, 0), false
}

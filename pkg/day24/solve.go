package day24

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/bfs"
	"adventofcode2016/pkg/utils/graph"
	"adventofcode2016/pkg/utils/grid"
	"adventofcode2016/pkg/utils/logger"
	"adventofcode2016/pkg/utils/slices"
	"math"
)

type Checkpoint struct {
	label string
	coord grid.Coord
}

var log = logger.New("day24")

func Solve(path string, partOne bool) int {
	log.Reset("starting")

	var checkpoints []Checkpoint
	maze := grid.Parse(utils.MustReadInput(path), func(value rune, x, y int) bool {
		switch value {
		case '.':
			return false
		case '#':
			return true
		default:
			log.Log("found checkpoint", "label", string(value), "x", x, "y", y)
			checkpoints = append(checkpoints, Checkpoint{string(value), grid.At(x, y)})
			return false
		}
	})

	distances := graph.NewGraph[Checkpoint]()
	for _, from := range checkpoints {
		for _, to := range checkpoints {
			if distances.Exists(from, to) {
				continue
			}

			path, ok := bfs.BFS(
				from.coord,
				func(target grid.Coord) []grid.Coord {
					neighbors := target.Cardinals()
					neighbors = slices.Filter(neighbors, func(coord grid.Coord) bool {
						value, ok := maze.GetContents(coord)
						if !ok {
							return false
						}
						return value == false
					})
					return neighbors
				},
				func(at grid.Coord) bool {
					return at == to.coord
				},
			)
			assert.True(ok, "could not find path between checkpoints", "from", from, "to", to)
			distances.AddBidirectionalEdge(from, to, float64(len(path)))
		}
	}

	paths := slices.Permutations(checkpoints)
	shortest := math.MaxInt

	for _, path := range paths {
		if path[0].label != "0" {
			continue
		}

		sum := 0
		for i := 1; i < len(path); i++ {
			distance, _ := distances.Distance(path[i-1], path[i])
			sum += int(distance) - 1
		}

		if !partOne {
			distance, _ := distances.Distance(path[len(path)-1], path[0])
			sum += int(distance) - 1
		}

		if sum < shortest {
			shortest = sum
		}

		log.Log("testing path", "path", slices.Map(path, func(c Checkpoint) string { return c.label }), "sum", sum)
	}

	return logger.Return(&log, shortest)
}

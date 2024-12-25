package grid

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/graph"
	"adventofcode2016/pkg/utils/slices"
	"math"
	"strings"
)

type Node[T any] struct {
	Coordinate Coord
	Contents   T
}

type Grid[T any] map[Coord]Node[T]

func NewGrid[T any]() Grid[T] {
	return make(Grid[T])
}

func Parse[T any](lines []string, init func(value rune, x int, y int) T) Grid[T] {
	grid := NewGrid[T]()
	for y, line := range lines {
		for x, char := range line {
			grid.InsertAt(x, y, init(char, x, y))
		}
	}
	return grid
}

func ParseRune(char rune, x, y int) rune {
	return char
}

func (g Grid[T]) At(x int, y int) (Node[T], bool) {
	return g.Get(At(x, y))
}

func (g Grid[T]) MustAt(x int, y int) Node[T] {
	return g.MustGet(At(x, y))
}

func (g Grid[T]) Get(coordinate Coord) (Node[T], bool) {
	value, ok := g[coordinate]
	return value, ok
}

func (g Grid[T]) GetContents(coordinate Coord) (T, bool) {
	value, ok := g[coordinate]
	return value.Contents, ok
}

func (g Grid[T]) MustGetContents(coordinate Coord) T {
	g.AssertPopulated(coordinate)
	return g[coordinate].Contents
}

func (g Grid[T]) Contains(coordinate Coord) bool {
	_, ok := g[coordinate]
	return ok
}

func (g Grid[T]) MustGet(coordinate Coord) Node[T] {
	g.AssertPopulated(coordinate)
	return g[coordinate]
}

func (g Grid[T]) GetAll(coordinates []Coord) []Node[T] {
	values := slices.MaybeMap(coordinates, func(coordinate Coord) (Node[T], bool) {
		value, ok := g[coordinate]
		return value, ok
	})

	return values
}

func (g Grid[T]) InsertAt(x, y int, value T) {
	g.Insert(At(x, y), value)
}

func (g Grid[T]) Insert(coordinate Coord, value T) {
	g.AssertEmpty(coordinate)
	g[coordinate] = Node[T]{
		Coordinate: coordinate,
		Contents:   value,
	}
}

func (g Grid[T]) ReplaceAt(x, y int, value T) {
	g.Replace(At(x, y), value)
}

func (g Grid[T]) Replace(coordinate Coord, value T) {
	g[coordinate] = Node[T]{
		Coordinate: coordinate,
		Contents:   value,
	}
}

func (g Grid[T]) String(empty rune, mapper func(contents T) rune) string {
	var builder strings.Builder
	minWidth, maxWidth := g.Width()
	minHeight, maxHeight := g.Height()

	for y := minHeight; y <= maxHeight; y++ {
		if y != minHeight {
			builder.WriteString("\n")
		}
		for x := minWidth; x <= maxWidth; x++ {
			node, ok := g.At(x, y)
			if !ok {
				builder.WriteRune(empty)
				continue
			}
			builder.WriteRune(mapper(node.Contents))
		}
	}
	return builder.String()
}

func (g Grid[T]) Width() (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for coordinate := range g {
		max = utils.Max(coordinate.X, max)
		min = utils.Min(coordinate.X, min)
	}
	return min, max
}

func (g Grid[T]) Height() (int, int) {
	max := math.MinInt
	min := math.MaxInt
	for coordinate := range g {
		max = utils.Max(coordinate.Y, max)
		min = utils.Min(coordinate.Y, min)
	}
	return min, max
}

func (g Grid[T]) UpdateAt(coordinate Coord, value T) {
	g.AssertPopulated(coordinate)
	g[coordinate] = Node[T]{coordinate, value}
}

func (g Grid[T]) AssertPopulated(coordinate Coord) {
	_, ok := g[coordinate]
	assert.True(ok, "expected value to be populated at coordinate", coordinate)
}

func (g Grid[T]) AssertEmpty(coordinate Coord) {
	_, ok := g[coordinate]
	assert.False(ok, "expected value to not be populated at coordinate", "coordinate", coordinate)
}

func (g Grid[T]) ToGraph(validator func(node Node[T]) bool) *graph.Graph[Coord] {
	grph := graph.NewGraph[Coord]()
	for coordinate, node := range g {
		if !validator(node) {
			continue
		}

		for _, neighbor := range coordinate.Cardinals() {
			if neighborNode, ok := g.Get(neighbor); ok && validator(neighborNode) {
				grph.AddBidirectionalEdge(coordinate, neighbor, 1)
			}
		}
	}
	return grph
}

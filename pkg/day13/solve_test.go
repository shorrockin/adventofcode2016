package day13

import (
	"adventofcode2016/pkg/grid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMazeIsOpenCalculation(t *testing.T) {
	maze := NewMaze(10)
	assert.True(t, maze.IsOpen(grid.At(1, 1)))
	assert.True(t, maze.IsOpen(grid.At(4, 4)))
	assert.False(t, maze.IsOpen(grid.At(0, 2)))
	assert.False(t, maze.IsOpen(grid.At(9, 1)))
	assert.True(t, maze.IsOpen(grid.At(9, 5)))
	assert.True(t, maze.IsOpen(grid.At(3, 3)))
}

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 11, StepsToPosition(10, grid.At(1, 1), grid.At(7, 4)))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 86, StepsToPosition(1364, grid.At(1, 1), grid.At(31, 39)))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 127, CountUniquePositions(1364, grid.At(1, 1), 50))
}

package grid

import "testing"
import "github.com/stretchr/testify/assert"

func TestGridAccess(t *testing.T) {
	grid := NewGrid[string]()
	grid.InsertAt(1, 1, "Center")
	grid.InsertAt(0, 0, "NorthWest")
	grid.InsertAt(1, 2, "South")
	grid.InsertAt(2, 1, "East")

	coordinates := At(1, 1).Compass()
	assert.Equal(t, 8, len(coordinates))

	neighbors := grid.GetAll(coordinates)
	assert.Equal(t, 3, len(neighbors))
	assert.Equal(t, "East", neighbors[0].Contents)
	assert.Equal(t, "South", neighbors[1].Contents)
	assert.Equal(t, "NorthWest", neighbors[2].Contents)

	_, ok := grid.At(1, 1)
	assert.True(t, ok)

	_, ok = grid.At(-1, 1)
	assert.False(t, ok)

	grid.UpdateAt(At(1, 1), "Chris")
	assert.Equal(t, "Chris", grid.MustGet(At(1, 1)).Contents)
}

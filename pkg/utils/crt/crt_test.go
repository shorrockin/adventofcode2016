package crt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChineeseRemainderTheorum(t *testing.T) {
	assert.Equal(t, 23, CRT([]int{2, 3, 2}, []int{3, 5, 7}))
	assert.Equal(t, int64(5), CRT([]int64{0, 1}, []int64{5, 2}))
}

package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, "DDRRRD", PartOne("ihgpwlah"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, "DUDRDLRRRD", PartOne("edjrjqaa"))
}

func TestPartTwoExample(t *testing.T) {
	assert.Equal(t, 370, PartTwo("ihgpwlah"))
	assert.Equal(t, 492, PartTwo("kglvqrro"))
	assert.Equal(t, 830, PartTwo("ulqzkmiv"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 502, PartTwo("edjrjqaa"))
}

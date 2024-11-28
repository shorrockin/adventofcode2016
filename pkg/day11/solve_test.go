package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloorDefinitionParser(t *testing.T) {
	items := parseFloor(1, "The first floor contains a polonium-compatible microchip, a thulium-compatible microchip, a thulium generator, a promethium generator")
	assert.Equal(t, 4, len(items))
	// assert.Equal(t, chips[0], "polonium")
	// assert.Equal(t, chips[1], "thulium")
	// assert.Equal(t, generators[0], "thulium")
	// assert.Equal(t, generators[1], "promethium")
}

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 11, Solve("input.example.txt"))
}

func TestFaciltyHashing(t *testing.T) {
	one := Facility{elevator: 1, items: []Item{NewChip("a", 1), NewChip("b", 2), NewGenerator("c", 3), NewGenerator("d", 4)}}
	two := Facility{elevator: 1, items: []Item{NewChip("a", 1), NewChip("b", 2), NewGenerator("c", 3), NewGenerator("d", 4)}}
	assert.Equal(t, one.hash(), two.hash())
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 47, Solve("input.partone.txt"))
}

func TestPartTwoActual(t *testing.T) {
	assert.Equal(t, 71, Solve("input.parttwo.txt"))
}

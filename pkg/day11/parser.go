package day11

import (
	"adventofcode2016/pkg/assert"
	"adventofcode2016/pkg/utils"
	"regexp"
)

var CHIP_REGEXP = regexp.MustCompile(`a ([a-z]+)-compatible microchip`)
var GENERATOR_REGEXP = regexp.MustCompile(`a ([a-z]+) generator`)

func parseFloor(floor int, definition string) []Item {
	mapper := func(match []string) string { return match[1] }
	chips := utils.Map(
		CHIP_REGEXP.FindAllStringSubmatch(definition, -1),
		mapper,
	)
	generators := utils.Map(
		GENERATOR_REGEXP.FindAllStringSubmatch(definition, -1),
		mapper,
	)

	return append(
		utils.Map(chips, func(name string) Item { return NewChip(name, floor) }),
		utils.Map(generators, func(name string) Item { return NewGenerator(name, floor) })...,
	)
}

func parse(path string) Facility {
	facility := Facility{elevator: 1, items: []Item{}}
	lines := utils.MustReadInput(path)
	items := []Item{}

	for idx, line := range lines {
		items = append(items, parseFloor(idx+1, line)...)
	}
	facility.items = items

	// sanity check to make sure all chips have generators and visa versa
	// for chip := range facility.items {
	// 	if _, ok := facility.generators[chip]; !ok {
	// 		assert.Fail("chip without generator", "chip", chip)
	// 	}
	// }
	// for generator := range facility.generators {
	// 	if _, ok := facility.chips[generator]; !ok {
	// 		assert.Fail("generator without chip", "generator", generator)
	// 	}
	// }

	// facility should start in a valid state
	if !facility.valid() {
		assert.Fail("invalid starting state", "facility", facility)
	}

	return facility
}

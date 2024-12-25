package day12

import "adventofcode2016/pkg/utils/assert"

type State struct {
	registers  []int
	operations []Operation
	current    int
}

func (s *State) read(register string) int {
	switch register {
	case "a":
		return s.registers[0]
	case "b":
		return s.registers[1]
	case "c":
		return s.registers[2]
	case "d":
		return s.registers[3]
	default:
		assert.Fail("unable to read register value", "register", register)
	}
	return 0
}

func (s *State) write(register string, value int) {
	switch register {
	case "a":
		s.registers[0] = value
	case "b":
		s.registers[1] = value
	case "c":
		s.registers[2] = value
	case "d":
		s.registers[3] = value
	default:
		assert.Fail("unable to write register value", register)
	}
}

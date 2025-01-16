package day23

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"strconv"
	"strings"
)

type Operation func(state *State)

func Solve(path string, c int) int {
	state := parse(path)
	state.write("c", c)

	for {
		if state.current < 0 || state.current >= len(state.operations) {
			return state.registers[0]
		}
		state.operations[state.current](state)
		state.current += 1
	}
}

func parse(path string) *State {
	var operations []Operation
	for _, line := range utils.MustReadInput(path) {
		fields := strings.Fields(line)
		switch fields[0] {
		case "cpy":
			assert.Equal(3, len(fields), "expected 3 fields for cpy", "line", line)
			operations = append(operations, cpy(fields[1], fields[2]))
		case "inc":
			assert.Equal(2, len(fields), "expected 2 fields for inc", "line", line)
			operations = append(operations, add(fields[1], 1))
		case "dec":
			assert.Equal(2, len(fields), "expected 2 fields for dec", "line", line)
			operations = append(operations, add(fields[1], -1))
		case "jnz":
			assert.Equal(3, len(fields), "expected 3 fields for jnz", "line", line)
			operations = append(operations, jnz(fields[1], utils.MustAtoi(fields[2])-1))
		default:
			assert.Fail("unknown field while parsing", "fields", fields)
		}
	}
	return &State{
		registers:  make([]int, 4),
		operations: operations,
		current:    0,
	}
}

func cpy(from string, to string) Operation {
	return func(state *State) {
		literal, err := strconv.Atoi(from)
		if err != nil {
			literal = state.read(from)
		}
		state.write(to, literal)
	}
}

func add(target string, amount int) Operation {
	return func(state *State) {
		state.write(target, state.read(target)+amount)
	}
}

func jnz(target string, amount int) Operation {
	return func(state *State) {
		literal, err := strconv.Atoi(target)
		if err != nil {
			literal = state.read(target)
		}

		if literal != 0 {
			state.current += amount
		}
	}
}

package day23

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"adventofcode2016/pkg/utils/logger"
	"strings"
)

var log = logger.New("day23").Disable()

type Operation struct {
	cmd  string
	args []string
}

func (op *Operation) eval(s *State) {
	log.Log("eval", "cmd", op.cmd, "args", op.args, "state", s.registers, "current", s.current)
	switch op.cmd {
	case "cpy":
		if s.isRegister(op.args[1]) {
			assert.Equal(2, len(op.args), "expected 2 args for cpy", "args", op.args)
			s.write(op.args[1], s.literal(op.args[0]))
		} else {
			log.Log("ignoring copy command", "args", op.args)
		}
	case "inc":
		assert.Equal(1, len(op.args), "expected 1 args for inc", "args", op.args)
		s.write(op.args[0], s.read(op.args[0])+1)
	case "dec":
		assert.Equal(1, len(op.args), "expected 1 args for dec", "args", op.args)
		s.write(op.args[0], s.read(op.args[0])-1)
	case "jnz":
		assert.Equal(2, len(op.args), "expected 2 args for jez", "args", op.args)
		literal := s.literal(op.args[0])
		if literal != 0 {
			s.current += s.literal(op.args[1]) - 1
		}
	case "tgl":
		assert.Equal(1, len(op.args), "expected 1 args for tgl", "args", op.args)
		offset := s.current + s.literal(op.args[0])
		assert.NotEqual(0, offset, "offset should not be 0", "offset", offset)

		if offset >= 0 && offset < len(s.operations) {
			toggled := &s.operations[offset]
			if len(toggled.args) == 1 {
				toggled.cmd = utils.Ternary(toggled.cmd == "inc", "dec", "inc")
			} else if len(toggled.args) == 2 {
				toggled.cmd = utils.Ternary(toggled.cmd == "jnz", "cpy", "jnz")
			}
		}

	// custom commands below added for optimization of logic
	case "add":
		assert.Equal(2, len(op.args), "expected 2 args for add", "args", op.args)
		s.write(op.args[0], s.read(op.args[0])+s.literal(op.args[1]))
	case "mul":
		assert.Equal(2, len(op.args), "expected 2 args for mul", "args", op.args)
		s.write(op.args[0], s.read(op.args[0])*s.literal(op.args[1]))
	case "noop":
		// do nothing

	default:
		assert.Fail("unknown operation", "op", op)
	}

	s.current += 1
}

func Solve(path string, a int) int {
	state := parse(path)
	state.write("a", a)

	for {
		if state.current < 0 || state.current >= len(state.operations) {
			return state.registers[0]
		}
		state.operations[state.current].eval(state)
	}
}

func parse(path string) *State {
	var operations []Operation
	for _, line := range utils.MustReadInput(path) {
		// support added for code comments in input files to allow for debugging
		if strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		operations = append(operations, Operation{cmd: fields[0], args: fields[1:]})
	}
	return &State{
		registers:  make([]int, 4),
		operations: operations,
		current:    0,
	}
}

package day21

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/assert"
	"slices"
	"strings"
)

func Solve(path, from string, partOne bool) string {
	password := []rune(from)

	lines := utils.MustReadInput(path)
	if !partOne {
		slices.Reverse(lines)
	}

	for _, line := range lines {
		parts := strings.Fields(line)
		switch parts[0] {
		case "swap":
			if parts[1] == "position" {
				swapPosition(utils.MustAtoi(parts[2]), utils.MustAtoi(parts[5]), &password)
			} else {
				swapLetter(rune(parts[2][0]), rune(parts[5][0]), &password)
			}
		case "rotate":
			if parts[1] == "based" {
				rotateLetter(rune(parts[6][0]), &password, partOne)
			} else {
				steps := utils.MustAtoi(parts[2])
				if !partOne {
					parts[1] = utils.Ternary(parts[1] == "left", "right", "left")
				}

				if parts[1] == "left" {
					rotateLeft(steps, &password)
				} else {
					rotateRight(steps, &password)
				}
			}
		case "reverse":
			reverse(utils.MustAtoi(parts[2]), utils.MustAtoi(parts[4]), &password)
		case "move":
			if partOne {
				move(utils.MustAtoi(parts[2]), utils.MustAtoi(parts[5]), &password)
			} else {
				move(utils.MustAtoi(parts[5]), utils.MustAtoi(parts[2]), &password)
			}
		default:
			panic(assert.Fail("unknown instruction", "instruction", parts[0]))
		}
	}

	return string(password)
}

func swapPosition(from, to int, password *[]rune) {
	(*password)[from], (*password)[to] = (*password)[to], (*password)[from]
}

func swapLetter(from, to rune, password *[]rune) {
	for idx, char := range *password {
		if char == from {
			(*password)[idx] = to
		} else if char == to {
			(*password)[idx] = from
		}
	}
}

func rotateLeft(steps int, password *[]rune) {
	steps = steps % len(*password)
	*password = append((*password)[steps:], (*password)[:steps]...)
}

func rotateRight(steps int, password *[]rune) {
	steps = steps % len(*password)
	*password = append((*password)[len(*password)-steps:], (*password)[:len(*password)-steps]...)
}

func rotateLetter(letter rune, password *[]rune, partOne bool) {
	idx := slices.Index(*password, letter)

	if partOne {
		rotateRight(utils.Ternary(idx >= 4, idx+2, idx+1), password)
	} else {
		rotateLeft(inverseRotateLetter(len(*password))[idx], password)
	}
}

func reverse(from, to int, password *[]rune) {
	for from < to {
		(*password)[from], (*password)[to] = (*password)[to], (*password)[from]
		from++
		to--
	}
}

func move(from, to int, password *[]rune) {
	char := (*password)[from]
	*password = append((*password)[:from], (*password)[from+1:]...)
	*password = append((*password)[:to], append([]rune{char}, (*password)[to:]...)...)
}

var cache = utils.NewCache[int, map[int]int]()

// TODO: probably should just do this with math instead of memoization working
// backwards, but couldn't quite pin down a generalized formula as there are
// some password lengths (e.g. 5) which can't be reversed.
func inverseRotateLetter(length int) map[int]int {
	return cache.Memoize(length, func() map[int]int {
		inverse := make(map[int]int)
		for i := 0; i < length; i++ {
			steps := 1 + i
			if i >= 4 {
				steps++
			}
			position := (i + steps) % length
			if _, exists := inverse[position]; exists {
				panic("duplicate inverse position exists, cannot inverse for this password length")
			}
			inverse[position] = steps
		}

		return inverse
	})
}

package day18

import (
	"adventofcode2016/pkg/utils"
	"adventofcode2016/pkg/utils/logger"
	"strings"
)

var log = logger.New("day18")
var cache = utils.NewCache[string, string]()

func Solve(input string, iterations int) int {
	log.Reset("starting", "input", input, "iterations", iterations)
	safe := utils.CountCharacters(input, '.')
	for range iterations {
		input = next(input)
		safe += utils.CountCharacters(input, '.')
	}
	return logger.LogReturn(&log, safe, "cache_hits", len(cache))
}

func next(previous string) string {
	return cache.Memoize(previous, func() string {
		var out strings.Builder
		for idx := range previous {
			left := idx == 0 || previous[idx-1] == '.'
			center := previous[idx] == '.'
			right := idx == len(previous)-1 || previous[idx+1] == '.'

			if !left && !center && right {
				out.WriteByte('^')
			} else if left && !center && !right {
				out.WriteByte('^')
			} else if !left && center && right {
				out.WriteByte('^')
			} else if left && center && !right {
				out.WriteByte('^')
			} else {
				out.WriteByte('.')
			}
		}

		return out.String()
	})
}

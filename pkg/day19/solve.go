package day19

import (
	"adventofcode2016/pkg/utils/logger"
	"math"
)

var log = logger.New("day19")

func PartOne(value int) int {
	return logger.Return(
		&log,
		2*(value-int(math.Pow(2, math.Floor(math.Log2(float64(value))))))+1,
		"from",
		value,
	)
}

func PartTwo(value int) int {
	log3 := math.Floor(math.Log(float64(value)) / math.Log(3))
	pow := int(math.Pow(3, log3))
	result := 2*value - 3*pow

	if value == pow {
		result = value
	} else if value-pow <= pow {
		result = value - pow
	}

	return logger.Return(&log, result, "from", value)
}

package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneExample(t *testing.T) {
	assert.Equal(t, 6, CountLights(7, 3, "input.example.txt"))
}

func TestPartOneActual(t *testing.T) {
	assert.Equal(t, 110, CountLights(50, 6, "input.txt"))
}

func TestPartTwoActual(t *testing.T) {
	// ZJHRKCPLYJ
	DisplayLights(50, 6, "input.txt")
	// log.Fatalf("intentional failure to log output")
}

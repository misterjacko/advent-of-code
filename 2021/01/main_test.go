package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}, LoadInput("input_test.txt"))
}

func TestIsGreater(t *testing.T) {
	assert.Equal(t, true, IsGreater(199, 200))
	assert.Equal(t, false, IsGreater(210, 207))
}

func TestPart1(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 7, Part1(input))
}

func TestCombine3(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	output := []int{
		607,
		618,
		618,
		617,
		647,
		716,
		769,
		792,
	}
	assert.Equal(t, output, Combine3(input))
}

func TestPart2(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 5, Part2(input))
}

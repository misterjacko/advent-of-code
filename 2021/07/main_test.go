package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []string{"3", "4", "3", "1", "2"}, MakeSlice("3,4,3,1,2"))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, map[int]int{
		0: 1, 1: 2, 2: 3, 4: 1, 7: 1, 14: 1, 16: 1,
	}, LoadInput("input_test.txt"))
}

func TestOrderedSlice(t *testing.T) {
	input := map[int]int{
		0: 1, 1: 2, 2: 3, 4: 1, 7: 1, 14: 1, 16: 1,
	}
	assert.Equal(t, []int{0, 1, 2, 4, 7, 14, 16}, OrderedSlice(input))
}

func TestPart1(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 37, Part1(input))
}

func TestPart2(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 168, Part2(input))
}

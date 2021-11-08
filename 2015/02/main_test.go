package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, [][]int{{2, 3, 4}, {1, 1, 10}}, LoadInput("input_test.txt"))
}

func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 58+43, Part1(data))
}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 34+14, Part2(data))
}

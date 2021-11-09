package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"}, LoadInput("input_test.txt"))
}

func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 2, Part1(data))
}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 11, Part2(data))
}

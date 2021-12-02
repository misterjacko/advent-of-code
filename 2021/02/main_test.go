package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}, LoadInput("input_test.txt"))
}

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []string{"forward", "5"}, MakeSlice("forward 5"))
	assert.Equal(t, []string{"down", "5"}, MakeSlice("down 5"))
	assert.Equal(t, []string{"up", "3"}, MakeSlice("up 3"))
}

func TestPart1(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 150, Part1(input))
}

func TestPart2(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 900, Part2(input))
}

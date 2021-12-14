package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	dotInput, instructions := LoadInput("input_test_short.txt")
	dotResult := map[Coord]int{
		Coord{
			col: 6,
			row: 10,
		}: 1,
		Coord{
			col: 0,
			row: 14,
		}: 1,
	}
	instructionResult := []Instruction{
		Instruction{
			fold:     "y",
			location: 7,
		},
		Instruction{
			fold:     "x",
			location: 5,
		},
	}

	assert.Equal(t, instructionResult, instructions)
	assert.Equal(t, dotResult, dotInput)
}

func TestPart1(t *testing.T) {
	dotInput, instructions := LoadInput("input_test.txt")

	assert.Equal(t, 17, Part1(dotInput, instructions, 2))
}

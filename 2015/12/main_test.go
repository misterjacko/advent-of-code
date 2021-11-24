package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	assert.Equal(t, []string{"[", "1", ",", "2", ",", "3", "]"}, LoadInput("input_test.txt"))
	assert.Equal(t, []string{
		"{", "\"", "a", "\"", ":", "2", ",", "\"", "b", "\"", ":", "4", "}",
	}, LoadInput("input_test2.txt"))
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, 3, StringToInt("3"))
}

func TestMakeNeg(t *testing.T) {
	assert.Equal(t, -3, MakeNeg(3))
}

func TestIsNegative(t *testing.T) {
	assert.Equal(t, true, IsNegative("-"))
	assert.Equal(t, false, IsNegative("{"))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 6, Part1(LoadInput("input_test.txt")))
	assert.Equal(t, 6, Part1(LoadInput("input_test2.txt")))
	assert.Equal(t, 13, Part1(LoadInput("input_test3.txt")))
}

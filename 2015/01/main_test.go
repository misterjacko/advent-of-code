package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{")", ")", "(", "(", "(", "(", "("}, LoadInput("input_test.txt"))
}

func TestPart1(t *testing.T) {
	// test_input := LoadInput("input_test.txt")
	assert.Equal(t, 3, Part1(LoadInput("input_test.txt")))
	assert.Equal(t, 3, Part1(LoadInput("input_test2.txt")))
	assert.Equal(t, -3, Part1(LoadInput("input_test3.txt")))
}

func TestPart2(t *testing.T) {
	// test_input := LoadInput("input_test.txt")
	assert.Equal(t, 1, Part2([]string{")"}))
	assert.Equal(t, 5, Part2([]string{"(", ")", "(", ")", ")"}))
}

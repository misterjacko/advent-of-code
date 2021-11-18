package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	assert.Equal(t, 1, StringToInt("1"))
}

func TestLookAndSay(t *testing.T) {
	assert.Equal(t, []int{1, 1}, LookAndSay([]int{1}))
	assert.Equal(t, []int{2, 1}, LookAndSay([]int{1, 1}))
	assert.Equal(t, []int{1, 2, 1, 1}, LookAndSay([]int{2, 1}))
	assert.Equal(t, []int{1, 1, 1, 2, 2, 1}, LookAndSay([]int{1, 2, 1, 1}))
}

func TestPart1(t *testing.T) {
	test_input := "1"
	assert.Equal(t, 2, Part1(test_input, 1))
	assert.Equal(t, 2, Part1(test_input, 2))
	assert.Equal(t, 4, Part1(test_input, 3))
	assert.Equal(t, 6, Part1(test_input, 4))
	assert.Equal(t, 6, Part1(test_input, 5))
}

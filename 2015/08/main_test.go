package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	stringResult := []string{
		"\"\"",
		"\"abc\"",
		"\"aaa\\\"aaa\"",
		"\"\\x27\"",
	}
	assert.Equal(t, stringResult, LoadInput("input_test.txt"))
}

func TestCodeLength(t *testing.T) {
	assert.Equal(t, 2, CodeLength("\"\""))
	assert.Equal(t, 5, CodeLength("\"abc\""))
	assert.Equal(t, 10, CodeLength("\"aaa\\\"aaa\""))
	assert.Equal(t, 6, CodeLength("\"\\x27\""))
}

func TestEscape(t *testing.T) {
	assert.Equal(t, false, Escape("\""))
	assert.Equal(t, false, Escape("a"))
	assert.Equal(t, true, Escape("\\"))
}

func TestSkipNumber(t *testing.T) {
	assert.Equal(t, 1, SkipNumber("\\"))
	assert.Equal(t, 1, SkipNumber("\""))
	assert.Equal(t, 3, SkipNumber("x"))
}

func TestStringLength(t *testing.T) {
	assert.Equal(t, 0, StringLength("\"\""))
	assert.Equal(t, 3, StringLength("\"abc\""))
	assert.Equal(t, 7, StringLength("\"aaa\\\"aaa\""))
	assert.Equal(t, 1, StringLength("\"\\x27\""))
}
func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 12, Part1(data))
}
func TestNeedsEscape(t *testing.T) {
	assert.Equal(t, true, NeedsEscape("\\"))
	assert.Equal(t, true, NeedsEscape("\""))
	assert.Equal(t, false, NeedsEscape("x"))
}

func TestStringLengthAdded(t *testing.T) {
	assert.Equal(t, 6, StringLengthAdded("\"\""))
	assert.Equal(t, 9, StringLengthAdded("\"abc\""))
	assert.Equal(t, 16, StringLengthAdded("\"aaa\\\"aaa\""))
	assert.Equal(t, 11, StringLengthAdded("\"\\x27\""))
}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 19, Part2(data))
}

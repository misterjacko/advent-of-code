package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, map[string][]string{
		"start": {"A", "b"},
		"A":     {"start", "c", "b", "end"},
		"b":     {"start", "A", "d", "end"},
		"c":     {"A"},
		"d":     {"b"},
		"end":   {"A", "b"},
	}, LoadInput("input_test1.txt"))
}

func TestIsUpper(t *testing.T) {
	assert.Equal(t, false, IsUpper("start"))
	assert.Equal(t, false, IsUpper("end"))
	assert.Equal(t, true, IsUpper("A"))
	assert.Equal(t, true, IsUpper("HN"))
	assert.Equal(t, false, IsUpper("sa"))
	assert.Equal(t, false, IsUpper("b"))
}

func TestPart1(t *testing.T) {
	data1 := LoadInput("input_test1.txt")
	data2 := LoadInput("input_test2.txt")
	data3 := LoadInput("input_test3.txt")
	assert.Equal(t, 10, Part1(data1))
	assert.Equal(t, 19, Part1(data2))
	assert.Equal(t, 226, Part1(data3))
}

// func TestMakeLowerList(t *testing.T) { // may not be in the right order
// 	assert.Equal(t, []string{"b", "c", "d"}, MakeLowerList(LoadInput("input_test1.txt")))
// 	assert.Equal(t, []string{"sa", "dc", "kj"}, MakeLowerList(LoadInput("input_test2.txt")))
// }

func TestPart2(t *testing.T) {
	data1 := LoadInput("input_test1.txt")
	data2 := LoadInput("input_test2.txt")
	data3 := LoadInput("input_test3.txt")
	assert.Equal(t, 36, Part2(data1)) //b +12 c +6 d +0  but both add 4
	assert.Equal(t, 103, Part2(data2))
	assert.Equal(t, 3509, Part2(data3))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	stringResult := []string{
		"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
		"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
	}
	assert.Equal(t, stringResult, LoadInput("input_test.txt"))
}

func TestArray(t *testing.T) {
	arrayResult := []string{
		"Comet", "can", "fly", "14", "km/s", "for", "10", "seconds,", "but", "then", "must", "rest", "for", "127", "seconds.",
	}
	assert.Equal(t, arrayResult, Array("Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds."))
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, -999, StringToInt("-999"))
	assert.Equal(t, 11, StringToInt("11"))
	assert.Equal(t, 0, StringToInt("0"))
}

func TestTotalCycles(t *testing.T) {
	assert.Equal(t, []int{7, 10}, TotalCycles(10, 127, 1000))
	assert.Equal(t, []int{5, 11}, TotalCycles(11, 162, 1000))
	assert.Equal(t, []int{1, 1}, TotalCycles(50, 50, 101))
}

func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 1120, Part1(data, 1000))
}

func TestPt2Cycles(t *testing.T) {
	assert.Equal(t, 8, Pt2Cycles(10, 127, 1000))
	assert.Equal(t, 6, Pt2Cycles(11, 162, 1000))
	assert.Equal(t, 2, Pt2Cycles(50, 50, 101))
}

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []int{
		0, 10, 20, 20, 20, 30, 40, 40, 40, 50, 60, 60, 60,
	}, MakeSlice(10, 2, 2, 3))
}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 689, Part2(data, 1000))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []string{"3", "4", "3", "1", "2"}, MakeSlice("34312"))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}, LoadInput("input_test.txt"))
}

func TestIsLowest(t *testing.T) {
	assert.Equal(t, false, IsLowest(2, []int{3, 1}))
	assert.Equal(t, true, IsLowest(1, []int{2, 9, 9}))
	assert.Equal(t, false, IsLowest(7, []int{9, 8, 8, 6}))

}

func TestMapEnds(t *testing.T) {
	data := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	answer1, _ := MapEnds(data[0], data[1])
	answer2, _ := MapEnds(data[len(data)-1], data[len(data)-2])
	assert.Equal(t, 3, answer1) // first line
	assert.Equal(t, 6, answer2) // last line
}

func TestMapCenter(t *testing.T) {
	data := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	answer1, _ := MapCenter(data[1], data[0], data[2])
	answer2, _ := MapCenter(data[2], data[1], data[3])
	assert.Equal(t, 0, answer1)
	assert.Equal(t, 6, answer2)
}

func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	answer, _ := Part1(data)
	assert.Equal(t, 15, answer)
}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test.txt")
	_, lowestPoints := Part1(data)
	answer := Part2(data, lowestPoints)
	assert.Equal(t, 1134, answer)
}

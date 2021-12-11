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
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}, LoadInput("input_test.txt"))
}

// func TestGrowOne(t *testing.T) {
// 	input := [][]int{
// 		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
// 		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
// 		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
// 		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
// 		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
// 		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
// 		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
// 		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
// 		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
// 		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
// 	}

// 	assert.Equal(t, input, GrowOne(input)) // i know it will fail
// }

// func TestGrowNeighbors(t *testing.T) {
// 	input := [][]int{
// 		{2, 2, 2, 2, 2},
// 		{2, 10, 10, 10, 2},
// 		{2, 10, 2, 10, 2},
// 		{2, 10, 10, 10, 2},
// 		{2, 2, 2, 2, 2},
// 	}
// 	response := [][]int{
// 		{3, 4, 5, 4, 3},
// 		{4, 0, 0, 0, 4},
// 		{5, 0, 0, 0, 5},
// 		{4, 0, 0, 0, 4},
// 		{3, 4, 5, 4, 3},
// 	}
// 	assert.Equal(t, response, GrowNeighbors(input))
// }

// func TestMapCenter(t *testing.T) {
// 	data := [][]int{
// 		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
// 		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
// 		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
// 		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
// 		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
// 	}
// 	answer1, _ := MapCenter(data[1], data[0], data[2])
// 	answer2, _ := MapCenter(data[2], data[1], data[3])
// 	assert.Equal(t, 0, answer1)
// 	assert.Equal(t, 6, answer2)
// }

func TestPart1(t *testing.T) {
	data1 := LoadInput("input_test.txt")
	data2 := LoadInput("input_test.txt")
	data3 := LoadInput("input_test.txt")
	assert.Equal(t, 0, Part1(data1, 1))
	assert.Equal(t, 35, Part1(data2, 2))
	assert.Equal(t, 204, Part1(data3, 10))
}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 195, Part2(data))
}

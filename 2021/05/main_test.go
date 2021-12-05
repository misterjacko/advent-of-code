package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []string{"0", "9", "5", "9"}, MakeSlice("0,9 -> 5,9"))
}

func TestMakeIntSlice(t *testing.T) {
	assert.Equal(t, []int{0, 9, 5, 9}, MakeIntSlice([]string{"0", "9", "5", "9"}))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, [][]int{
		{0, 9, 5, 9},
		{8, 0, 0, 8},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{6, 4, 2, 0},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
		{0, 0, 8, 8},
		{5, 5, 8, 2},
	}, LoadInput("input_test.txt"))
}

func TestMake10Array(t *testing.T) {
	result := [10][10]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	assert.Equal(t, result, Make10Array())
}

// is it horiz

func TestIsHorizontal(t *testing.T) {
	assert.Equal(t, true, IsHorizontal([]int{0, 9, 5, 9}))  //horiz
	assert.Equal(t, false, IsHorizontal([]int{8, 0, 0, 8})) // diag /
	assert.Equal(t, false, IsHorizontal([]int{2, 2, 2, 1})) // vert
	assert.Equal(t, false, IsHorizontal([]int{0, 0, 5, 5})) // diag \
}

// is it vert
func TestIsVertical(t *testing.T) {
	assert.Equal(t, false, IsVertical([]int{0, 9, 5, 9})) //horiz
	assert.Equal(t, false, IsVertical([]int{8, 0, 0, 8})) // diag /
	assert.Equal(t, true, IsVertical([]int{2, 2, 2, 1}))  // vert
	assert.Equal(t, false, IsVertical([]int{0, 0, 5, 5})) // diag \
}

// do horiz
func TestMarkHorizontal(t *testing.T) {
	array := Make10Array()
	result1 := [10][10]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
	}
	assert.Equal(t, result1, MarkHorizontal([]int{0, 9, 5, 9}, array)) //horiz
	result2 := [10][10]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	assert.Equal(t, result2, MarkHorizontal([]int{0, 0, 0, 0}, array)) //horiz
}

// do vert
func TestMarkVertical(t *testing.T) {
	array := Make10Array()
	result := [10][10]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	assert.Equal(t, result, MarkVertical([]int{2, 2, 2, 1}, array)) //horiz
}

func TestPart1(t *testing.T) {
	array := Make10Array()
	input := LoadInput("input_test.txt")

	assert.Equal(t, 5, Part1(input, array))
}

func TestMarkDiagonal(t *testing.T) {
	array := Make10Array()
	result1 := [10][10]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	assert.Equal(t, result1, MarkDiagonal([]int{1, 1, 2, 2}, array))
	assert.Equal(t, result1, MarkDiagonal([]int{2, 2, 1, 1}, array))
	result2 := [10][10]int{
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	assert.Equal(t, result2, MarkDiagonal([]int{0, 2, 2, 0}, array))
	assert.Equal(t, result2, MarkDiagonal([]int{2, 0, 0, 2}, array))
	result3 := [10][10]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	assert.Equal(t, result3, MarkDiagonal([]int{1, 2, 7, 8}, array))
}

func TestPart2(t *testing.T) {
	array := Make10Array()
	input := LoadInput("input_test.txt")
	assert.Equal(t, 12, Part2(input, array))
}

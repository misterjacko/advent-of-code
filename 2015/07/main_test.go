package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	stringResult := []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	}
	assert.Equal(t, stringResult, LoadInput("input_test.txt"))
}

func TestArray(t *testing.T) {
	arrayResult := []string{
		"123", "->", "x",
	}
	assert.Equal(t, arrayResult, Array("123 -> x"))

	arrayResult1 := []string{
		"x", "AND", "y", "->", "d",
	}
	assert.Equal(t, arrayResult1, Array("x AND y -> d"))

	arrayResult2 := []string{
		"NOT", "x", "->", "h",
	}
	assert.Equal(t, arrayResult2, Array("NOT x -> h"))
}

// only numbers after LSHIFT or RSHIFT
// or len(slice) = 3

// func TestStringToInt(t *testing.T) {
// 	assert.Equal(t, -999, StringToInt("-999"))
// 	assert.Equal(t, 11, StringToInt("11"))
// 	assert.Equal(t, 0, StringToInt("0"))
// }

// func TestIntToBinary(t *testing.T) {
// 	assert.Equal(t, "1111011", IntToBinary(123))
// 	assert.Equal(t, "111001000", IntToBinary(456))
// }

func TestAND(t *testing.T) {
	assert.Equal(t, 1, AND(5, 3))
	assert.Equal(t, 2, AND(3, 2))
}

func TestOR(t *testing.T) {
	assert.Equal(t, 7, OR(5, 3))
	assert.Equal(t, 10, OR(8, 2))
}
func TestLSHIFT(t *testing.T) {
	assert.Equal(t, 240, LSHIFT(60, 2))
}
func TestRSHIFT(t *testing.T) {
	assert.Equal(t, 15, RSHIFT(60, 2))
}

func TestNOT(t *testing.T) {
	assert.Equal(t, 65528, NOT(7))
	assert.Equal(t, 65364, NOT(171))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 123, Part1(LoadInput("input_test.txt"), "x"))
	assert.Equal(t, 456, Part1(LoadInput("input_test.txt"), "y"))
	assert.Equal(t, 72, Part1(LoadInput("input_test.txt"), "d"))
	assert.Equal(t, 507, Part1(LoadInput("input_test.txt"), "e"))
	assert.Equal(t, 492, Part1(LoadInput("input_test.txt"), "f"))
	assert.Equal(t, 114, Part1(LoadInput("input_test.txt"), "g"))
	assert.Equal(t, 65412, Part1(LoadInput("input_test.txt"), "h"))
	assert.Equal(t, 65079, Part1(LoadInput("input_test.txt"), "i"))

}

// func TestPart2(t *testing.T) {
// 	data := LoadInput("input_test2.txt")
// 	assert.Equal(t, 1+2000000, Part2(data))
// }

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}, LoadInput("input_test.txt"))
}

func TestInitSlice(t *testing.T) {
	assert.Equal(t, []int{0, 0, 0, 0, 0}, InitSlice("00100"))
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, InitSlice("101111101000"))
}

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []string{"0", "0", "1", "0", "0"}, MakeSlice("00100"))
	assert.Equal(t, []string{"1", "1", "1", "0", "0"}, MakeSlice("11100"))
}

func TestConvertBinary(t *testing.T) {
	assert.Equal(t, 22, ConvertBinary("10110"))
	assert.Equal(t, 9, ConvertBinary("01001"))
}

func TestPart1(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 198, Part1(input))
}

func TestMostCommon(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, "1", MostCommon(input, 0))
	assert.Equal(t, "0", MostCommon(input, 1))
	assert.Equal(t, "1", MostCommon([]string{"1011", "1010", "0011", "0110"}, 0))
	assert.Equal(t, "0", MostCommon([]string{"1011", "1010", "0011", "0110", "0111"}, 0))
}

func TestLeastCommon(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, "0", LeastCommon(input, 0))
	assert.Equal(t, "1", LeastCommon(input, 1))
	assert.Equal(t, "0", LeastCommon([]string{"1011", "1010", "0011", "0110"}, 0))
}
func TestKeepOnly(t *testing.T) {
	input := LoadInput("input_test.txt")
	output1 := []string{
		"11110",
		"10110",
		"10111",
		"10101",
		"11100",
		"10000",
		"11001",
	}
	output2 := []string{
		"10110",
		"10111",
		"10101",
		"10000",
	}
	input2 := KeepOnly(input, "1", 0)
	assert.Equal(t, output1, KeepOnly(input, "1", 0))
	assert.Equal(t, output2, KeepOnly(input2, "0", 1))
}

func TestPart2(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 230, Part2(input))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// to test
// multiply, add, halt overwwrite
// read file

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3", "4"}, LoadInput("input_test.txt"))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 21, Add(10, 11))
	assert.Equal(t, 21, Multiply(3, 7))

}

// func TestPart2(t *testing.T) {
// 	var floatValue float64 = 543
// 	testVal := ConvertIntToFlo64(543)

// 	assert.Equal(t, floatValue, testVal)
// }

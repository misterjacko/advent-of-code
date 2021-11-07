package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, []int{1234, 65432, 1}, LoadInput("input_test.txt"))
}

// func TestPart2(t *testing.T) {
// 	var floatValue float64 = 543
// 	testVal := ConvertIntToFlo64(543)

// 	assert.Equal(t, floatValue, testVal)
// }

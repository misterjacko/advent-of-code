package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadInput(t *testing.T) {
	assert.Equal(t, []int{1234, 65432, 1}, LoadInput("input_test.txt"))
}

func Test_ConvertStrToFlo64(t *testing.T) {
	var floatValue float64 = 543
	testVal := ConvertIntToFlo64(543)
	assert.Equal(t, floatValue, testVal)
}

func Test_FindFuel(t *testing.T) {
	assert.Equal(t, 2, FindFuel(12))
	assert.Equal(t, 2, FindFuel(14))
	assert.Equal(t, 654, FindFuel(1969))
	assert.Equal(t, 33583, FindFuel(100756))
}

func Test_FuelForFuel(t *testing.T) {
	assert.Equal(t, 2, FuelForFuel(14))
	assert.Equal(t, 966, FuelForFuel(1969))
	assert.Equal(t, 50346, FuelForFuel(100756))
}

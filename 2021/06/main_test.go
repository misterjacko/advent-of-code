package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInt(t *testing.T) {
	assert.Equal(t, 123, MakeInt("123"))
}

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []string{"3", "4", "3", "1", "2"}, MakeSlice("3,4,3,1,2"))
}

func TestLoadInput(t *testing.T) {
	assert.Equal(t, map[int]int{
		0: 0, 1: 1, 2: 1, 3: 2, 4: 1, 5: 0, 6: 0, 7: 0, 8: 0,
	}, LoadInput("input_test.txt"))
}

func TestNextDay(t *testing.T) {
	assert.Equal(t, map[int]int{
		0: 1, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0,
	}, NextDay(map[int]int{
		0: 0, 1: 1, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0,
	}))
	assert.Equal(t, map[int]int{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 1, 7: 0, 8: 1,
	}, NextDay(map[int]int{
		0: 1, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0,
	}))
	assert.Equal(t, map[int]int{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 2, 7: 0, 8: 1,
	}, NextDay(map[int]int{
		0: 1, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 1, 8: 0,
	}))
	// assert.Equal(t, map[int]int{
	// 	0: 1, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0,
	// }, NextDay(map[int]int{
	// 	0: 0, 1: 1, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0,
	// }))
}

func TestPart1(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 26, Part1(input, 18))
	assert.Equal(t, 5934, Part1(input, 80))
	assert.Equal(t, 26984457539, Part1(input, 256))
}

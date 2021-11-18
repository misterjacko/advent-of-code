package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	stringResult := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	assert.Equal(t, stringResult, LoadInput("input_test.txt"))
}

func TestArray(t *testing.T) {
	arrayResult := []string{
		"London", "to", "Dublin", "=", "464",
	}
	assert.Equal(t, arrayResult, Array("London to Dublin = 464"))
}

func TestMakeSet(t *testing.T) {
	testInput := [][]string{
		{"London", "to", "Dublin", "=", "464"},
		{"London", "to", "Belfast", "=", "518"},
		{"Dublin", "to", "Belfast", "=", "141"},
	}
	// arrayResult := []string{
	// 	"London", "Dublin", "Belfast",
	// }
	assert.Equal(t, 3, len(MakeSet(testInput)))
}

func TestMakeCombinations(t *testing.T) {
	testInput := []string{"London", "Dublin", "Belfast"}
	assert.Equal(t, 6, len(MakeCombinations(testInput)))
}

func TestMakeMapKey(t *testing.T) {
	assert.Equal(t, []string{"London-Dublin", "Dublin-London"}, MakeMapKey("London", "Dublin"))
}

func TestMakeRoutes(t *testing.T) {
	testInput := [][]string{
		{"London", "to", "Dublin", "=", "464"},
		{"London", "to", "Belfast", "=", "518"},
	}
	testResult := map[string]int{
		"London-Dublin":  464,
		"Dublin-London":  464,
		"London-Belfast": 518,
		"Belfast-London": 518,
	}
	assert.Equal(t, testResult, MakeRoutes(testInput))
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, -999, StringToInt("-999"))
	assert.Equal(t, 11, StringToInt("11"))
	assert.Equal(t, 0, StringToInt("0"))
}

// func TestPart1(t *testing.T) {
// 	data := LoadInput("input_test.txt")
// 	assert.Equal(t, 605, Part1(data))
// }

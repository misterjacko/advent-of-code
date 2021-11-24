package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	stringResult := []string{
		"Alice would gain 54 happiness units by sitting next to Bob.",
		"Alice would lose 79 happiness units by sitting next to Carol.",
		"Alice would lose 2 happiness units by sitting next to David.",
		"Bob would gain 83 happiness units by sitting next to Alice.",
		"Bob would lose 7 happiness units by sitting next to Carol.",
		"Bob would lose 63 happiness units by sitting next to David.",
		"Carol would lose 62 happiness units by sitting next to Alice.",
		"Carol would gain 60 happiness units by sitting next to Bob.",
		"Carol would gain 55 happiness units by sitting next to David.",
		"David would gain 46 happiness units by sitting next to Alice.",
		"David would lose 7 happiness units by sitting next to Bob.",
		"David would gain 41 happiness units by sitting next to Carol.",
	}
	assert.Equal(t, stringResult, LoadInput("input_test.txt"))
}

func TestArray(t *testing.T) {
	arrayResult := []string{
		"Alice", "would", "gain", "54", "happiness", "units", "by", "sitting", "next", "to", "Bob",
	}
	assert.Equal(t, arrayResult, Array("Alice would gain 54 happiness units by sitting next to Bob."))
}

func TestMakeSet(t *testing.T) {
	testInput := [][]string{
		{"Alice", "would", "gain", "54", "happiness", "units", "by", "sitting", "next", "to", "Bob"},
		{"Alice", "would", "lose", "79", "happiness", "units", "by", "sitting", "next", "to", "Carol"},
		{"Alice", "would", "lose", "2", "happiness", "units", "by", "sitting", "next", "to", "David"},
		{"Bob", "would", "gain", "83", "happiness", "units", "by", "sitting", "next", "to", "Alice"},
		{"Bob", "would", "lose", "7", "happiness", "units", "by", "sitting", "next", "to", "Carol"},
		{"Bob", "would", "lose", "63", "happiness", "units", "by", "sitting", "next", "to", "David"},
		{"Carol", "would", "lose", "62", "happiness", "units", "by", "sitting", "next", "to", "Alice"},
		{"Carol", "would", "gain", "60", "happiness", "units", "by", "sitting", "next", "to", "Bob"},
		{"Carol", "would", "gain", "55", "happiness", "units", "by", "sitting", "next", "to", "David"},
		{"David", "would", "gain", "46", "happiness", "units", "by", "sitting", "next", "to", "Alice"},
		{"David", "would", "lose", "7", "happiness", "units", "by", "sitting", "next", "to", "Bob"},
		{"David", "would", "gain", "41", "happiness", "units", "by", "sitting", "next", "to", "Carol"},
	}
	// arrayResult := []string{
	// 	"London", "Dublin", "Belfast",
	// }
	assert.Equal(t, 4, len(MakeSet(testInput)))
}

func TestMakeCombinations(t *testing.T) {
	testInput := []string{"Alice", "Bob", "Carol", "David"}
	assert.Equal(t, 24, len(MakeCombinations(testInput)))
}

func TestMakeMapKey(t *testing.T) {
	assert.Equal(t, []string{"Alice-Bob", "Bob-Alice"}, MakeMapKey("Alice", "Bob"))
}

func TestMakeRelationships(t *testing.T) {
	testInput := [][]string{
		{"Alice", "would", "gain", "54", "happiness", "units", "by", "sitting", "next", "to", "Bob"},
		{"Alice", "would", "lose", "79", "happiness", "units", "by", "sitting", "next", "to", "Carol"},
		{"Bob", "would", "gain", "83", "happiness", "units", "by", "sitting", "next", "to", "Alice"},
	}
	testResult := map[string]int{
		"Alice-Bob":   137,
		"Bob-Alice":   137,
		"Alice-Carol": -79,
		"Carol-Alice": -79,
	}
	assert.Equal(t, testResult, MakeRelationships(testInput))
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, -999, StringToInt("-999"))
	assert.Equal(t, 11, StringToInt("11"))
	assert.Equal(t, 0, StringToInt("0"))
}

func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 330, Part1(data))
}

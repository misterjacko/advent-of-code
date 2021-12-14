package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	input, inserts, _, _ := LoadInput("input_test_short.txt")
	inputResult := map[string]int{
		"NN": 1, "NC": 1, "CB": 1,
	}
	insertsResult := map[string]PolyPlan{
		"CH": PolyPlan{
			pair1: "CB",
			pair2: "BH",
		},
		"HH": PolyPlan{
			pair1: "HN",
			pair2: "NH",
		},
	}
	assert.Equal(t, inputResult, input)
	assert.Equal(t, insertsResult, inserts)
}

func TestCalcString(t *testing.T) {
	input, _, beginning, end := LoadInput("input_test_short.txt")
	result := map[string]int{
		"N": 2, "C": 1, "B": 1,
	}
	assert.Equal(t, result, CalcString(input, beginning, end))
}

func TestTakeTurn(t *testing.T) {
	input, inserts, _, _ := LoadInput("input_test.txt")
	result := map[string]int{
		"NC": 1,
		"CN": 1,
		"NB": 1,
		"BC": 1,
		"CH": 1,
		"HB": 1,
	}
	assert.Equal(t, result, TakeTurn(input, inserts))
}

func TestPart1(t *testing.T) {
	input, inserts, beginning, end := LoadInput("input_test.txt")
	assert.Equal(t, 1588, Part1(input, inserts, beginning, end, 10))
}

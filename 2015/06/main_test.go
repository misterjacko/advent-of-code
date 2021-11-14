package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	stringResult := []string{
		"turn on 0,0 through 999,999",
		"toggle 0,0 through 999,0",
		"turn off 499,499 through 500,500",
	}
	assert.Equal(t, stringResult, LoadInput("input_test.txt"))
}

func TestArray(t *testing.T) {
	arrayResult := []string{
		"on", "0,0", "999,999",
	}
	assert.Equal(t, arrayResult, Array("turn on 0,0 through 999,999"))

	arrayResult1 := []string{
		"toggle", "0,0", "999,0",
	}
	assert.Equal(t, arrayResult1, Array("toggle 0,0 through 999,0"))
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, -999, StringToInt("-999"))
	assert.Equal(t, 11, StringToInt("11"))
	assert.Equal(t, 0, StringToInt("0"))
}

func TestMakeInstruction(t *testing.T) {
	arrayResult := &Instruction{
		Command:  "on",
		RowStart: 0,
		RowEnd:   999,
		ColStart: 0,
		ColEnd:   999,
	}
	assert.Equal(t, arrayResult, MakeInstruction([]string{"on", "0,0", "999,999"}))

	arrayResult1 := &Instruction{
		Command:  "toggle",
		RowStart: 0,
		RowEnd:   999,
		ColStart: 0,
		ColEnd:   0,
	}
	assert.Equal(t, arrayResult1, MakeInstruction([]string{"toggle", "0,0", "999,0"}))
}

func TestChangeLight(t *testing.T) {
	assert.Equal(t, true, ChangeLight(false, "on"))
	assert.Equal(t, true, ChangeLight(false, "toggle"))
	assert.Equal(t, false, ChangeLight(false, "off"))
	assert.Equal(t, true, ChangeLight(true, "on"))
	assert.Equal(t, false, ChangeLight(true, "toggle"))
	assert.Equal(t, false, ChangeLight(true, "off"))
}

func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 1000000-1000-4, Part1(data))
}

func TestChangeLightBrightness(t *testing.T) {
	assert.Equal(t, 4, ChangeLightBrightness(3, "on"))
	assert.Equal(t, 5, ChangeLightBrightness(3, "toggle"))
	assert.Equal(t, 2, ChangeLightBrightness(3, "off"))
	assert.Equal(t, 0, ChangeLightBrightness(0, "off"))

}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test2.txt")
	assert.Equal(t, 1+2000000, Part2(data))
}

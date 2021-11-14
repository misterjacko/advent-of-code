package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	stringResult := []string{
		"ugknbfddgicrmopn",
		"aaa",
		"jchzalrnumimnmhp",
		"haegwjzuvuyypxyu",
		"dvszwmarrgswjxmb",
	}
	assert.Equal(t, stringResult, LoadInput("input_test.txt"))
}

func TestNoForbiddenSubString(t *testing.T) {
	assert.Equal(t, true, NoForbiddenSubString("ugknbfddgicrmopn"))
	assert.Equal(t, false, NoForbiddenSubString("haegwjzuvuyypxyu"))
}

func TestIsVowel(t *testing.T) {
	assert.Equal(t, 1, IsVowel("a"))
	assert.Equal(t, 1, IsVowel("u"))
	assert.Equal(t, 0, IsVowel("t"))
}

func TestIsDupe(t *testing.T) {
	assert.Equal(t, true, IsDupe("a", "a"))
	assert.Equal(t, false, IsDupe("", "a"))
	assert.Equal(t, false, IsDupe("a", "b"))
}

func TestIterate(t *testing.T) {
	assert.Equal(t, true, Iterate("ugknbfddgicrmopn"))
	assert.Equal(t, false, Iterate("dvszwmarrgswjxmb"))
	assert.Equal(t, false, Iterate("jchzalrnumimnmhp"))
}

func TestPart1(t *testing.T) {
	data := LoadInput("input_test.txt")
	assert.Equal(t, 2, Part1(data))
}

func TestArray(t *testing.T) {
	assert.Equal(t, []string{"q", "j", "h", "v"}, Array("qjhv"))
}

func TestSandwich(t *testing.T) {
	assert.Equal(t, true, Sandwich("qjhvhtzxzqqjkmpb"))
	assert.Equal(t, false, Sandwich("uurcxstgmygtbstg"))
}

func TestTwoRepeat(t *testing.T) {
	assert.Equal(t, true, TwoRepeat("qjhvhtzxzqqjkmpb"))
	assert.Equal(t, false, TwoRepeat("ieodomkazucvgmuy"))
}

func TestPart2(t *testing.T) {
	data := LoadInput("input_test2.txt")
	assert.Equal(t, 2, Part2(data))
}

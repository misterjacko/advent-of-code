package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	result := [][][]string{
		{{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, {"fdgacbe", "cefdb", "cefbgd", "gcbe"}},
		{{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, {"fcgedb", "cgb", "dgebacf", "gc"}},
		{{"fgaebd", "cg", "bdaec", "gdafb", "agbcfd", "gdcbef", "bgcad", "gfac", "gcb", "cdgabef"}, {"cg", "cg", "fdcagb", "cbg"}},
		{{"fbegcd", "cbd", "adcefb", "dageb", "afcb", "bc", "aefdc", "ecdab", "fgdeca", "fcdbega"}, {"efabcd", "cedba", "gadfec", "cb"}},
		{{"aecbfdg", "fbg", "gf", "bafeg", "dbefa", "fcge", "gcbea", "fcaegb", "dgceab", "fcbdga"}, {"gecf", "egdcabf", "bgf", "bfgea"}},
		{{"fgeab", "ca", "afcebg", "bdacfeg", "cfaedg", "gcfdb", "baec", "bfadeg", "bafgc", "acf"}, {"gebdcfa", "ecba", "ca", "fadegcb"}},
		{{"dbcfg", "fgd", "bdegcaf", "fgec", "aegbdf", "ecdfab", "fbedc", "dacgb", "gdcebf", "gf"}, {"cefg", "dcbef", "fcge", "gbcadfe"}},
		{{"bdfegc", "cbegaf", "gecbf", "dfcage", "bdacg", "ed", "bedf", "ced", "adcbefg", "gebcd"}, {"ed", "bcgafe", "cdgba", "cbgef"}},
		{{"egadfb", "cdbfeg", "cegd", "fecab", "cgb", "gbdefca", "cg", "fgcdab", "egfdb", "bfceg"}, {"gbdfcae", "bgc", "cg", "cgb"}},
		{{"gcafb", "gcf", "dcaebfg", "ecagb", "gf", "abcdeg", "gaef", "cafbge", "fdbac", "fegbdc"}, {"fgae", "cfgab", "fg", "bagce"}},
	}
	assert.Equal(t, result, LoadInput("input_test.txt"))
}

func TestPart1(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 26, Part1(input))
}

func TestFourMinusOne(t *testing.T) {
	assert.Equal(t, "ef", FourMinusOne("eafb", "ab"))
}

func TestSortSignalPattern(t *testing.T) {
	input := []string{
		"be",
		"cfbegad",
		"cbdgef",
		"fgaecd",
		"cgeb",
		"fdcge",
		"agebfd",
		"fecdb",
		"fabcd",
		"edb",
	}
	output := []string{
		"be",
		"edb",
		"cgeb",
		"fecdb",
		"fabcd",
		"fdcge",
		"agebfd",
		"cbdgef",
		"fgaecd",
		"cfbegad",
	}

	assert.Equal(t, output, SortSignalPattern(input))
}

func TestHasOne(t *testing.T) {
	assert.Equal(t, false, HasOne("gcdfa", "ab"))
	assert.Equal(t, true, HasOne("fbcad", "ab"))
	assert.Equal(t, false, HasOne("cdfbe", "ab"))
}

func TestHasFourMinusOne(t *testing.T) {
	assert.Equal(t, false, HasFourMinusOne("gcdfa", "ef"))
	// assert.Equal(t, true, HasFourMinusOne("fbcad", "ab")) // ruled out
	assert.Equal(t, true, HasFourMinusOne("cdfbe", "ef"))
}

func TestSortKeys(t *testing.T) {
	assert.Equal(t, "acdfg", SortKeys("gcdfa"))
}

func TestCalculate(t *testing.T) {
	inputCodes := []string{
		"cdfeb", "fcadb", "cdfeb", "cdbaf",
	}
	inputMap := map[string]string{
		"bcdef":   "5",
		"acdfg":   "2",
		"abcdf":   "3",
		"abd":     "7",
		"abcdef":  "9",
		"abcdefg": "8",
		"bcdefg":  "6",
		"abef":    "4",
		"abcdeg":  "0",
		"ab":      "1",
	}
	assert.Equal(t, 5353, Calculate(inputMap, inputCodes))
}

func TestPart2(t *testing.T) {
	input := LoadInput("input_test.txt")
	assert.Equal(t, 61229, Part2(input))
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}, LoadInput("input_test.txt"))
}

func TestScore(t *testing.T) {
	assert.Equal(t, 3, Score(")"))
	assert.Equal(t, 57, Score("]"))
	assert.Equal(t, 1197, Score("}"))
	assert.Equal(t, 25137, Score(">"))
}
func TestRemove(t *testing.T) {
	assert.Equal(t, "{([(<[}>{{[(", Remove("{([(<{}[<>[]}>{[]{[(<()>"))
	assert.Equal(t, "[[<[())<([[[]]", Remove("[[<[([]))<([[{}[[()]]]"))
	assert.Equal(t, "[{[{(]}([{[{{}(", Remove("[{[{({}]{}}([{[{{{}}([]"))
}

func TestRemoveAllPairs(t *testing.T) {
	assert.Equal(t, "{([(<[}>{{[(", RemoveAllPairs("{([(<{}[<>[]}>{[]{[(<()>"))
	assert.Equal(t, "[[<[)<([", RemoveAllPairs("[[<[([]))<([[{}[[()]]]"))
	assert.Equal(t, "[{[{(]}([{[{(", RemoveAllPairs("[{[{({}]{}}([{[{{{}}([]"))
}

func TestCloser(t *testing.T) {
	assert.Equal(t, true, Closer(")"))
	assert.Equal(t, true, Closer("]"))
	assert.Equal(t, true, Closer("}"))
	assert.Equal(t, true, Closer(">"))
	assert.Equal(t, false, Closer("("))
	assert.Equal(t, false, Closer("{"))
	assert.Equal(t, false, Closer("["))
	assert.Equal(t, false, Closer("<"))
}

func TestPart1(t *testing.T) {
	p1, p2 := Part1(LoadInput("input_test.txt"))
	assert.Equal(t, 26397, p1)
	assert.Equal(t, 288957, p2)
}

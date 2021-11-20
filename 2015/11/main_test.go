package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotForbidden(t *testing.T) {
	assert.Equal(t, true, NotForbidden("a"))
	assert.Equal(t, false, NotForbidden("i"))
	assert.Equal(t, false, NotForbidden("o"))
	assert.Equal(t, false, NotForbidden("l"))
	assert.Equal(t, false, NotForbidden("hijklmmn"))
	assert.Equal(t, true, NotForbidden("abbceffg"))
	assert.Equal(t, true, NotForbidden("ghjaabcc"))
}

func TestHasPairs(t *testing.T) {
	assert.Equal(t, false, HasPairs("hijklmmn"))
	assert.Equal(t, true, HasPairs("abbceffg"))
	assert.Equal(t, false, HasPairs("abcdefgh"))
	assert.Equal(t, true, HasPairs("ghjaabcc"))
	assert.Equal(t, true, HasPairs("ghjaabaa"))
	assert.Equal(t, false, HasPairs("ghjaaagc"))
}

func TestHasStraight(t *testing.T) {
	assert.Equal(t, true, HasStraight("hijklmmn"))
	assert.Equal(t, false, HasStraight("abbceffg"))
	assert.Equal(t, true, HasStraight("abcdefgh"))
	assert.Equal(t, true, HasStraight("ghjaabcc"))
	assert.Equal(t, false, HasStraight("ghjaabaa"))
}

func TestIncrement(t *testing.T) {
	assert.Equal(t, "b", Increment("a"))
	assert.Equal(t, "a", Increment("z"))
	assert.Equal(t, "j", Increment("i"))
	assert.Equal(t, "r", Increment("q"))
}

func TestIncrementPass(t *testing.T) {
	assert.Equal(t, "aaab", IncrementPass("aaaa"))
	assert.Equal(t, "aaba", IncrementPass("aaaz"))
	assert.Equal(t, "baaa", IncrementPass("azzz"))

}

func TestPart1(t *testing.T) {
	assert.Equal(t, "abcdffaa", Part1("abcdefgh"))
	assert.Equal(t, "ghjaabcc", Part1("ghijklmn"))

}

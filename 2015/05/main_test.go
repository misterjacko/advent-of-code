package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 609043, Part1("abcdef"))
	assert.Equal(t, 1048970, Part1("pqrstuv"))
}

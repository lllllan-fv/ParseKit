package parsekit

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestParse(t *testing.T) {
	c := qt.New(t)

	c.Assert(1, qt.Equals, 1)
}

func TestIsBracketed(t *testing.T) {
	c := qt.New(t)

	var strs = []string{"{}", "{1}", "{}1", "1{}", "1", "1[]", "[]1", "[1]", "[]"}
	var isBracketeds = []bool{true, true, false, false, false, false, false, false, false}
	var isSquareBracketeds = []bool{false, false, false, false, false, false, false, true, true}

	for index, str := range strs {
		c.Assert(isBracketed(str), qt.Equals, isBracketeds[index])
		c.Assert(isSquareBracketed(str), qt.Equals, isSquareBracketeds[index])
	}
}

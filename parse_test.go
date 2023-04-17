package parsekit

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestParse(t *testing.T) {
	c := qt.New(t)

	c.Assert(1, qt.Equals, 1)
}

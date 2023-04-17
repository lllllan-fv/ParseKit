package parsekit

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

type bracketedRule struct {
	Rule               string
	IsBracketed        bool
	ParseBracket       []string
	IsSquareBracketed  bool
	parseSquareBracket []string
}

func createTestRules() []bracketedRule {
	return []bracketedRule{
		{"{}", true, []string{""}, false, nil},
		{"{1}", true, []string{"1"}, false, nil},
		{"{1,2}", true, []string{"1", "2"}, false, nil},
		{"{1,2,}", true, []string{"1", "2", ""}, false, nil},
		{"[]", false, nil, true, []string{""}},
		{"[1]", false, nil, true, []string{"1"}},
		{"[1,2]", false, nil, true, []string{"1", "2"}},
		{"[1,2,]", false, nil, true, []string{"1", "2", ""}},
	}
}

func TestParse(t *testing.T) {
	c := qt.New(t)

	c.Assert(1, qt.Equals, 1)
}

func TestParseBracket(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		rules, _ := parseBracket(rule.Rule)
		c.Assert(rules, qt.DeepEquals, rule.ParseBracket)
	}
}

func TestParseSquareBracket(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		rules, _ := parseSquareBracket(rule.Rule)
		c.Assert(rules, qt.DeepEquals, rule.parseSquareBracket)
	}
}


func TestIsBracketed(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		c.Assert(isBracketed(rule.Rule), qt.Equals, rule.IsBracketed)
	}
}

func TestIsSquareBracketed(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		c.Assert(isSquareBracketed(rule.Rule), qt.Equals, rule.IsSquareBracketed)
	}
}

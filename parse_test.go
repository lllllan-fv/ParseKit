package parsekit

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

type bracketedRule struct {
	rule               string
	isBracketed        bool
	parseBracket       []string
	isSquareBracketed  bool
	parseSquareBracket []string
	ans                [][]int
}

func createTestRules() []bracketedRule {
	return []bracketedRule{
		{"{}", true, []string{""}, false, nil, nil},
		{"{1}", true, []string{"1"}, false, nil, [][]int{{1}}},
		{"{1,2}", true, []string{"1", "2"}, false, nil, [][]int{{1, 2}, {2, 1}}},
		{"{1,2,}", true, []string{"1", "2", ""}, false, nil, nil},
		{"[]", false, nil, true, []string{""}, nil},
		{"[1]", false, nil, true, []string{"1"}, [][]int{{1}}},
		{"[1,2]", false, nil, true, []string{"1", "2"}, [][]int{{1, 2}}},
		{"[1,2,]", false, nil, true, []string{"1", "2", ""}, nil},
		{"1{}", false, nil, false, nil, nil},
		{"{}1", false, nil, false, nil, nil},
		{"1[]", false, nil, false, nil, nil},
		{"[]1", false, nil, false, nil, nil},
		{"[1,[2]]", false, nil, true, []string{"1", "[2]"}, [][]int{{1, 2}}},
		{"[[1],[2]]", false, nil, true, []string{"[1]", "[2]"}, [][]int{{1, 2}}},
		{"[1,[2,3]]", false, nil, true, []string{"1", "[2,3]"}, [][]int{{1, 2, 3}}},
		{"[1,{2,3}]", false, nil, true, []string{"1", "{2,3}"}, [][]int{{1, 2, 3}, {1, 3, 2}}},
		{"{1,[2,3]}", true, []string{"1", "[2,3]"}, false, nil, [][]int{{1, 2, 3}, {2, 3, 1}}},
	}
}

func TestParse(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		ans, _ := Parse(rule.rule)
		err := qt.Commentf("rule = %v", rule.rule)
		c.Assert(ans, qt.DeepEquals, rule.ans, err)
	}
}

func TestPermute(t *testing.T) {
	c := qt.New(t)

	var a1 = [][]int{{1, 2}}
	var a2 = [][]int{{3, 4}}
	var a3 = [][]int{{5, 6}, {7, 8}}

	c.Assert(permute([][][]int{a1, a2, a3}), qt.DeepEquals, [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 7, 8},
		{1, 2, 5, 6, 3, 4},
		{1, 2, 7, 8, 3, 4},
		{3, 4, 1, 2, 5, 6},
		{3, 4, 1, 2, 7, 8},
		{3, 4, 5, 6, 1, 2},
		{3, 4, 7, 8, 1, 2},
		{5, 6, 1, 2, 3, 4},
		{5, 6, 3, 4, 1, 2},
		{7, 8, 1, 2, 3, 4},
		{7, 8, 3, 4, 1, 2},
	})
}

func TestCombine(t *testing.T) {
	c := qt.New(t)

	var a1 = [][]int{{1, 2}, {2, 1}}
	var a2 = [][]int{{3, 4}, {4, 3}}
	var a3 = [][]int{{5, 6, 7}, {5, 7, 6}, {7, 5, 6}}

	c.Assert(combine([][][]int{a1, a2, a3}), qt.DeepEquals, [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, 7, 6},
		{1, 2, 3, 4, 7, 5, 6},
		{1, 2, 4, 3, 5, 6, 7},
		{1, 2, 4, 3, 5, 7, 6},
		{1, 2, 4, 3, 7, 5, 6},
		{2, 1, 3, 4, 5, 6, 7},
		{2, 1, 3, 4, 5, 7, 6},
		{2, 1, 3, 4, 7, 5, 6},
		{2, 1, 4, 3, 5, 6, 7},
		{2, 1, 4, 3, 5, 7, 6},
		{2, 1, 4, 3, 7, 5, 6},
	})
}

func TestParseBracket(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		rules, _ := parseBracket(rule.rule)
		err := qt.Commentf("rule = %v", rule.rule)
		c.Assert(rules, qt.DeepEquals, rule.parseBracket, err)
	}
}

func TestParseSquareBracket(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		rules, _ := parseSquareBracket(rule.rule)
		err := qt.Commentf("rule = %v", rule.rule)
		c.Assert(rules, qt.DeepEquals, rule.parseSquareBracket, err)
	}
}

func TestIsBracketed(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		err := qt.Commentf("rule = %v", rule.rule)
		c.Assert(isBracketed(rule.rule), qt.Equals, rule.isBracketed, err)
	}
}

func TestIsSquareBracketed(t *testing.T) {
	c := qt.New(t)

	for _, rule := range createTestRules() {
		err := qt.Commentf("rule = %v", rule.rule)
		c.Assert(isSquareBracketed(rule.rule), qt.Equals, rule.isSquareBracketed, err)
	}
}

func TestSplitString(t *testing.T) {
	c := qt.New(t)

	c.Assert([]string{""}, qt.DeepEquals, split(""))
	c.Assert([]string{"", ""}, qt.DeepEquals, split(","))
	c.Assert([]string{"[]"}, qt.DeepEquals, split("[]"))
	c.Assert([]string{"", "[]"}, qt.DeepEquals, split(",[]"))
	c.Assert([]string{"", "[,]"}, qt.DeepEquals, split(",[,]"))
	c.Assert([]string{"", "{,}"}, qt.DeepEquals, split(",{,}"))
	c.Assert([]string{"", "{[}]"}, qt.DeepEquals, split(",{[}]"))
	c.Assert([]string{"", "{}", "[,]"}, qt.DeepEquals, split(",{},[,]"))
	c.Assert([]string{"", "[{[][]}]"}, qt.DeepEquals, split(",[{[][]}]"))
}

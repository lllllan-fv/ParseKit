package parsekit

import (
	"errors"
	"strings"
)

func Parse(rule string) (ans [][]int, err error) {

	return
}

func parseBracket(rule string) (rules []string, err error) {
	if !isBracketed(rule) {
		return nil, errors.New("rule not bracketed")
	}
	return strings.Split(rule[1:len(rule)-1], ","), nil
}

func parseSquareBracket(rule string) (rules []string, err error) {
	if !isSquareBracketed(rule) {
		return nil, errors.New("rule not square bracketed")
	}
	return strings.Split(rule[1:len(rule)-1], ","), nil
}

func isBracketed(str string) bool {
	return strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}")
}

func isSquareBracketed(str string) bool {
	return strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")
}

package parsekit

import "strings"

func Parse(rule string) (ans [][]int, err error) {

	return
}

func isBracketed(str string) bool {
	return strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}")
}

func isSquareBracketed(str string) bool {
	return strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")
}

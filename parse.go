package parsekit

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

func Parse(rule string) (ans [][]int, err error) {
	if rules, err := parseSquareBracket(rule); err == nil {
		var mulAns [][][]int
		for _, rule := range rules {
			subAns, err := Parse(rule)
			if err != nil {
				return nil, err
			}
			mulAns = append(mulAns, subAns)
		}

		return combine(mulAns), nil
	} else if rules, err := parseBracket(rule); err == nil {
		fmt.Printf("rules: %v\n", rules)
		return nil, nil
	}

	i, err := cast.ToIntE(rule)
	if err != nil {
		return nil, err
	}

	return [][]int{{i}}, nil
}

func combine(arrs [][][]int) (res [][]int) {
	combineDFS(arrs, 0, []int{}, &res)
	return
}

func combineDFS(arrs [][][]int, pos int, cur []int, res *[][]int) {
	if pos == len(arrs) {
		var tmp = make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for _, arr := range arrs[pos] {
		combineDFS(arrs, pos+1, append(cur, arr...), res)
	}
}

func parseBracket(rule string) (rules []string, err error) {
	if !isBracketed(rule) {
		return nil, errors.New("rule not bracketed")
	}
	return split(rule[1 : len(rule)-1]), nil
}

func parseSquareBracket(rule string) (rules []string, err error) {
	if !isSquareBracketed(rule) {
		return nil, errors.New("rule not square bracketed")
	}
	return split(rule[1 : len(rule)-1]), nil
}

func split(s string) []string {
	var result []string

	// 记录左括号出现的数量
	left := 0

	// 记录当前部分的起始位置
	start := 0

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch c {
		case ',':
			// 如果当前是逗号，而且左括号数量为 0，说明这个部分已经结束
			if left == 0 {
				result = append(result, s[start:i])
				start = i + 1
			}
		case '{', '[':
			// 如果当前是左括号，增加数量
			left++
		case '}', ']':
			// 如果当前是右括号，减少数量
			left--
		}
	}

	// 处理最后一个部分
	if start <= len(s) {
		result = append(result, s[start:])
	}

	return result
}

func isBracketed(str string) bool {
	return strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}")
}

func isSquareBracketed(str string) bool {
	return strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")
}

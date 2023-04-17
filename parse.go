package parsekit

import (
	"errors"
	"strings"

	"github.com/spf13/cast"
)

// 字符串解析
// 字符串内容包含：数字和字符 ',[]{}'
// [] 内表示固定顺序的内容，{} 内表示无序的内容，解析字符串获得所有可能的数字排序
// eg: [1,2,3] -> [1,2,3]
// eg: {1,2,3} -> [1,2,3] [1,3,2] [2,1,3] [2,3,1] [3,1,2] [3,2,1]
// eg: {1,[2,3]} -> [1,2,3] [2,3,1]
// eg: [1,{2,3}] -> [1,2,3] [1,3,2]
func Parse(rule string) (ans [][]int, err error) {
	if rules, err := parseSquareBracket(rule); err == nil {
		// 如果当前字符串格式为 [...]，则对其中内容进行组合拼接
		var mulAns [][][]int
		// 对其所有子部分进行递归解析
		for _, rule := range rules {
			subAns, err := Parse(rule)
			if err != nil {
				return nil, err
			}
			mulAns = append(mulAns, subAns)
		}

		return combine(mulAns), nil
	} else if rules, err := parseBracket(rule); err == nil {
		// 如果当前字符串格式为 {...}，则对其中内容进行全排列
		var mulAns [][][]int
		for _, rule := range rules {
			subAns, err := Parse(rule)
			if err != nil {
				return nil, err
			}
			mulAns = append(mulAns, subAns)
		}

		return permute(mulAns), nil
	}

	// 不符合上述格式，则作为单个数字看待
	i, err := cast.ToIntE(rule)
	if err != nil {
		return nil, err
	}

	return [][]int{{i}}, nil
}

// 三维数组全排列组成二维数组，示例见 TestPermute
func permute(arrs [][][]int) (res [][]int) {
	var vis = make([]int, len(arrs))
	permuteDFS(arrs, 0, vis, []int{}, &res)
	return res
}

// 回溯实现全排列
func permuteDFS(arrs [][][]int, dep int, vis, cur []int, res *[][]int) {
	if dep == len(arrs) {
		var tmp = make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for index, arr := range arrs {
		if vis[index] == 1 {
			continue
		}

		vis[index] = 1
		for _, sub := range arr {
			permuteDFS(arrs, dep+1, vis, append(cur, sub...), res)
		}
		vis[index] = 0
	}
}

// 三维数组组合拼接成二维数组，示例见 TestCombine
func combine(arrs [][][]int) (res [][]int) {
	combineDFS(arrs, 0, []int{}, &res)
	return
}

// 递归实现拼接
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

// 字符串满足 {...} 格式时按照最大颗粒拆分
func parseBracket(rule string) (rules []string, err error) {
	if !isBracketed(rule) {
		return nil, errors.New("rule not bracketed")
	}
	return split(rule[1 : len(rule)-1]), nil
}

// 字符串满足 [...] 格式时按照最大颗粒拆分
func parseSquareBracket(rule string) (rules []string, err error) {
	if !isSquareBracketed(rule) {
		return nil, errors.New("rule not square bracketed")
	}
	return split(rule[1 : len(rule)-1]), nil
}

// 字符串拆分，',' 为分隔符， '[] {}' 看作完整单位
func split(s string) []string {
	var left = 0
	var start = 0
	var result []string

	for i, c := range s {
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

// 是否 {} 包裹
func isBracketed(str string) bool {
	return strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}")
}

// 是否 [] 包裹
func isSquareBracketed(str string) bool {
	return strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")
}

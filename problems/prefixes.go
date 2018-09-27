package answers

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var (
		b   strings.Builder
		min = shortestWord(strs)
	)
	for i := 0; i < min; i++ {
		ok := isCommon(strs, i)
		if !ok {
			return b.String()
		}
		b.WriteString(string(strs[0][i]))
	}
	return b.String()
}

func isCommon(strs []string, idx int) bool {
	cur := strs[0][idx]
	for i := 1; i < len(strs); i++ {
		char := strs[i][idx]
		if char != cur {
			return false
		}
	}
	return true
}

func shortestWord(strs []string) int {
	min := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	return min
}

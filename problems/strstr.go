package answers

import "strings"

func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	return strings.Index(haystack, needle)
}

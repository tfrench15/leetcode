package answers

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tt := []struct {
		str string
		ans int
	}{
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"", 0},
		{"au", 2},
		{"bwf", 3},
	}
	for _, tc := range tt {
		ret := lengthOfLongestSubstring(tc.str)
		if ret != tc.ans {
			t.Errorf("Error: expected %v, got %v", tc.ans, ret)
		}
	}
}

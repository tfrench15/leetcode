package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		ans    []int
	}{
		{[]int{2, 7, 5, 11}, 9, []int{0, 1}},
		{[]int{1, 2, 3}, 10, nil},
	}
	for _, tc := range tt {
		actual := twoSum(tc.nums, tc.target)
		expected := tc.ans
		if (actual == nil) != (expected == nil) {
			t.Errorf("Error: actual is %v, expected is %v", actual, expected)
		}
		if len(actual) != len(expected) {
			t.Errorf("Error: length of actual is %v, length of expected is %v", len(actual), len(expected))
		}
		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("Error: actual is equal to %v, expected is equal to %v", actual, expected)
			}
		}
	}
}

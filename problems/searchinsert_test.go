package answers

import "testing"

func TestSearchInsert(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}
	for _, tc := range tt {
		got := searchInsert(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("Error: expected %v, got %v", tc.want, got)
		}
	}
}

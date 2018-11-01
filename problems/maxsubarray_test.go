package answers

import "testing"

func TestMaxSubArray(t *testing.T) {
	tt := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}, 6},
		{[]int{-2, 1}, 1},
		{[]int{-2, -1}, -1},
		{[]int{-2, -3, -1}, -1},
	}
	for _, tc := range tt {
		if got := maxSubArray(tc.nums); got != tc.want {
			t.Errorf("Error: expected %v, got %v", tc.want, got)
		}
	}
}

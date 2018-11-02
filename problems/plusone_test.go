package answers

import "testing"

func TestPlusOne(t *testing.T) {
	tt := []struct {
		digits []int
		want   []int
	}{
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{2, 9, 9, 9}, []int{3, 0, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{8, 9}, []int{9, 0}},
		{[]int{9, 8, 2, 1, 3, 3, 1, 8, 1, 4, 4, 7, 2, 7, 2, 0, 5, 6, 8, 9, 7, 7, 4, 3}, []int{9, 8, 2, 1, 3, 3, 1, 8, 1, 4, 4, 7, 2, 7, 2, 0, 5, 6, 8, 9, 7, 7, 4, 4}},
	}
	for _, tc := range tt {
		got := plusOne(tc.digits)
		for i, num := range got {
			if num != tc.want[i] {
				t.Errorf("Error: expected %v, got %v", tc.want, got)
			}
		}
	}
}

package answers

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		heights []int
		area    int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}
	for _, test := range tests {
		ans := maxArea(test.heights)
		if ans != test.area {
			t.Errorf("Error: expected %v, got %v", test.area, ans)
		}
	}
}

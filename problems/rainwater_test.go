package answers

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		heights []int
		vol     int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for _, test := range tests {
		ans := trap(test.heights)
		if ans != test.vol {
			t.Errorf("Error: expected %v, got %v", test.vol, ans)
		}
	}
}

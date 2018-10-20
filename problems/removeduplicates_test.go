package answers

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for _, test := range tests {
		got := removeDuplicates(test.input)
		if got != test.want {
			t.Errorf("Error: expected %d, got %d", test.want, got)
		}
	}
}

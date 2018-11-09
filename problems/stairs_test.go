package answers

import "testing"

func TestClimbStairs(t *testing.T) {
	tt := []struct {
		stairs int
		want   int
	}{
		{2, 2},
		{3, 3},
	}
	for _, tc := range tt {
		got := climbStairs(tc.stairs)
		if got != tc.want {
			t.Errorf("Error: expected %v, got %v", tc.want, got)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tt := []struct {
		num  int
		want int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
	}
	for _, tc := range tt {
		if got := fibonacci(tc.num); got != tc.want {
			t.Errorf("Error: expected %v, got %v", tc.want, got)
		}
	}
}

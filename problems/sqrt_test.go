package answers

import "testing"

func TestMySqrt(t *testing.T) {
	tt := []struct {
		num  int
		want int
	}{
		{4, 2},
		{8, 2},
	}
	for _, tc := range tt {
		got := mySqrt(tc.num)
		if got != tc.want {
			t.Errorf("Error: expected %v, got %v", tc.want, got)
		}
	}
}

package answers

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tt := []struct {
		str  string
		want int
	}{
		{"Hello, World", 5},
		{"a ", 1},
	}
	for _, tc := range tt {
		got := lengthOfLastWord(tc.str)
		if got != tc.want {
			t.Errorf("Error: expected %v, got %v", tc.want, got)
		}
	}
}

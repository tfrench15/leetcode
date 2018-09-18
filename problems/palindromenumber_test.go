package answers

import "testing"

func TestIsNumPalindrome(t *testing.T) {
	tests := []struct {
		num int
		exp bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}
	for _, test := range tests {
		ans := isNumPalindrome(test.num)
		if ans != test.exp {
			t.Error("Error: incorrect answer")
		}
	}
}

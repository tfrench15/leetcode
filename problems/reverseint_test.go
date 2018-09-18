package answers

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		num int
		exp int
	}{
		{123, 321},
		{1, 1},
		{-456, -654},
		{12345, 54321},
	}
	for _, test := range tests {
		if ans := reverse(test.num); ans != test.exp {
			t.Errorf("Error: expected %v, got %v", test.exp, ans)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		str string
		exp string
	}{
		{"abc", "cba"},
		{"hello", "olleh"},
		{"h", "h"},
		{"12345", "54321"},
	}
	for _, test := range tests {
		if ans := reverseString(test.str); ans != test.exp {
			t.Errorf("Error: expected %v, got %v", test.exp, ans)
		}
	}
}

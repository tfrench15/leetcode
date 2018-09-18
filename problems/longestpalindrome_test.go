package answers

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		str string
		exp string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"abb", "bb"},
		{"ac", "a"},
	}
	for _, test := range tests {
		ans := longestPalindrome(test.str)
		if ans != test.exp {
			t.Errorf("Error: expected %v, got %v", test.exp, ans)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s string
		b bool
	}{
		{"hannah", true},
		{"car", false},
		{"a", true},
		{"racecar", true},
		{"potato", false},
		{"bb", true},
	}
	for _, test := range tests {
		ans := isPalindrome(test.s)
		if ans != test.b {
			t.Error("Error: incorrect bool")
		}
	}
}

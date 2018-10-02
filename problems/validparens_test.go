package answers

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		str string
		exp bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, test := range tests {
		if ok := isValid(test.str); ok != test.exp {
			t.Errorf("Error: expected %v, got %v", test.exp, ok)
		}
	}
}

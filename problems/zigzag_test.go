package answers

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		str  string
		rows int
		ans  string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}
	for _, test := range tests {
		exp := convert(test.str, test.rows)
		if exp != test.ans {
			t.Errorf("Error: expected %v, got %v", test.ans, exp)
		}
	}
}

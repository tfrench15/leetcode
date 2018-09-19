package answers

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		rom string
		exp int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, test := range tests {
		ans := romanToInt(test.rom)
		if ans != test.exp {
			t.Errorf("Error: expected %v, got %v", test.exp, ans)
		}
	}
}
func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num int
		exp string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}
	for _, test := range tests {
		ans := intToRoman(test.num)
		if ans != test.exp {
			t.Errorf("Error: expected %v, got %v", test.exp, ans)
		}
	}
}

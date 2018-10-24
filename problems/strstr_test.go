package answers

import "testing"

func TestStrStr(t *testing.T) {
	tt := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"asdfasdf", "", 0},
	}
	for _, tc := range tt {
		got := strStr(tc.haystack, tc.needle)
		if got != tc.want {
			t.Errorf("Error: expected %v, got %v", tc.want, got)
		}
	}
}

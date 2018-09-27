package answers

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs   []string
		prefix string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}
	for _, test := range tests {
		ans := longestCommonPrefix(test.strs)
		if ans != test.prefix {
			t.Errorf("Error: expected %v, got %v", test.prefix, ans)
		}
	}
}

func TestIsCommon(t *testing.T) {
	tests := []struct {
		slc []string
		idx int
		b   bool
	}{
		{[]string{"hello", "helix", "help"}, 0, true},
		{[]string{"hello", "helix", "help"}, 1, true},
		{[]string{"hello", "helix", "help"}, 3, false},
	}
	for _, test := range tests {
		ok := isCommon(test.slc, test.idx)
		if ok != test.b {
			t.Errorf("Error: expected %v, got %v", test.b, ok)
		}
	}
}

func TestShortestWord(t *testing.T) {
	tests := []struct {
		slc []string
		exp int
	}{
		{[]string{"hello", "hi", "hey"}, 2},
		{[]string{"what", "where"}, 4},
		{[]string{"howdy"}, 5},
	}
	for _, test := range tests {
		ans := shortestWord(test.slc)
		if ans != test.exp {
			t.Errorf("Error: expected %v, got %v", test.exp, ans)
		}
	}
}

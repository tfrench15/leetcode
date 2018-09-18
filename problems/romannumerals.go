package answers

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	conv := 0
	for i, ch := range s {
		if i < len(s)-1 {
			if string(ch) == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				conv--
				continue
			}
			if string(ch) == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				conv -= m["X"]
				continue
			}
			if string(ch) == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
				conv -= m["C"]
				continue
			}
		}
		conv += m[string(ch)]
	}
	return conv
}

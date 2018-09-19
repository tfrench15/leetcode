package answers

import (
	"strings"
)

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

func intToRoman(num int) string {
	var b strings.Builder
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	thousands := num / 1000
	num -= thousands * 1000
	for thousands > 0 {
		b.WriteString(m[1000])
		thousands--
	}

	hundreds := num / 100
	num -= hundreds * 100
	for {
		if hundreds == 9 {
			b.WriteString(m[900])
			break
		}
		if hundreds < 9 && hundreds >= 5 {
			b.WriteString(m[500])
			for hundreds > 5 {
				b.WriteString(m[100])
				hundreds--
			}
			break
		}
		if hundreds == 4 {
			b.WriteString(m[400])
			break
		}
		if hundreds <= 3 {
			for hundreds > 0 {
				b.WriteString(m[100])
				hundreds--
			}
			break
		}
	}

	tens := num / 10
	num -= tens * 10
	for {
		if tens == 9 {
			b.WriteString(m[90])
			break
		}
		if tens < 9 && tens >= 5 {
			b.WriteString(m[50])
			for tens > 5 {
				b.WriteString(m[10])
				tens--
			}
			break
		}
		if tens == 4 {
			b.WriteString(m[40])
			break
		}
		if tens <= 3 {
			for tens > 0 {
				b.WriteString(m[10])
				tens--
			}
			break
		}
	}

	ones := num
	for {
		if ones == 9 {
			b.WriteString(m[9])
			break
		}
		if ones < 9 && ones >= 5 {
			b.WriteString(m[5])
			for ones > 5 {
				b.WriteString(m[1])
				ones--
			}
			break
		}
		if ones == 4 {
			b.WriteString(m[4])
			break
		}
		if ones <= 3 {
			for ones > 0 {
				b.WriteString(m[1])
				ones--
			}
			break
		}
	}
	return b.String()
}

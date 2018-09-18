package answers

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	if x < 0 {
		s := strconv.Itoa(x)
		s = strings.Replace(s, "-", "", 1)
		reversed := reverseString(s)
		ans, err := strconv.Atoi(reversed)
		if err != nil {
			log.Fatalf("Error converting to int: %v", err)
		}
		if (float64(ans) * (-1)) < math.Pow(-2, 31) {
			return 0
		}
		return (-1) * ans
	}
	s := strconv.Itoa(x)
	reversed := reverseString(s)
	ans, err := strconv.Atoi(reversed)
	if err != nil {
		log.Fatalf("Error converting to int: %v", err)
	}
	if float64(ans) > math.Pow(2, 31)-1 {
		return 0
	}
	return ans
}

func reverseString(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

package answers

import (
	"fmt"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	num := digitsToNum(digits)
	num++
	result := numToDigits(num)
	return result
}

func digitsToNum(digits []int) int {
	result := 0
	multiplier := 1
	for i := len(digits) - 1; i >= 0; i-- {
		result += (digits[i] * multiplier)
		multiplier *= 10
	}
	return result
}

func numToDigits(num int) []int {
	s := strconv.Itoa(num)
	slc := strings.Split(s, "")
	fmt.Println(slc)
	digits := []int{}
	for _, ch := range slc {
		digit, err := strconv.Atoi(ch)
		if err != nil {
			panic(err)
		}
		digits = append(digits, digit)
	}
	return digits
}

func secondAttempt(digits []int) []int {
	i := len(digits) - 1
	// carry := 0
	for i >= 0 {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		if digits[i] == 9 && i == 0 {
			digits[i] = 1
			digits = append(digits, 0)
			return digits
		}
		digits[i] = 0
		i--
	}
	return digits
}

package answers

import "strconv"

func isNumPalindrome(x int) bool {
	num := strconv.Itoa(x)
	i, j := 0, len(num)-1
	for i <= j {
		if num[i] != num[j] {
			return false
		}
		i++
		j--
	}
	return true
}

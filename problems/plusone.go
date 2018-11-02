package answers

func plusOne(digits []int) []int {
	i := len(digits) - 1
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

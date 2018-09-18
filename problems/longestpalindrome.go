package answers

func longestPalindrome(s string) string {
	// Easy cases
	if len(s) == 0 || isPalindrome(s) {
		return s
	}
	var ret string
	i, k := 0, len(s)-1
	for k >= 1 {
		j := k
		for j <= len(s)-1 {
			pal := s[i : j+1]
			if isPalindrome(pal) {
				ret = pal
				return ret
			}
			i++
			j++
		}
		i = 0
		k--
	}
	if ret == "" {
		ret = string(s[0])
	}
	return ret
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

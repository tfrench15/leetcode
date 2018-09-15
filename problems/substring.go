package answers

func lengthOfLongestSubstring(s string) int {
	// Easy cases
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var lengths []int
	var length int
	for i := 0; i < len(s)-1; i++ {
		seen := make(map[byte]bool)
		length = 0
		for j := i; j <= len(s)-1; j++ {
			_, ok := seen[s[j]]
			if !ok {
				length++
				seen[s[j]] = true
				if j == len(s)-1 {
					lengths = append(lengths, length)
				}
			} else {
				lengths = append(lengths, length)
				break
			}
		}
	}
	if length != 0 {
		lengths = append(lengths, length)
	}
	if len(lengths) == 0 {
		return 0
	}

	max := maxInt(lengths)
	return max
}

func maxInt(slc []int) int {
	max := slc[0]
	for _, v := range slc {
		if v > max {
			max = v
		}
	}
	return max
}

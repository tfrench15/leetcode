package answers

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	slc := strings.Split(s, " ")
	for i := len(slc) - 1; i >= 0; i-- {
		if len(slc[i]) != 0 {
			return len(slc[i])
		}
		continue
	}
	return 0
}

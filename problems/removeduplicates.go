package answers

import "fmt"

func removeDuplicates(nums []int) int {
	for i := 1; i < len(nums); i++ {
		cur, prev := nums[i], nums[i-1]
		if cur == prev {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	fmt.Println(nums)
	return len(nums)
}

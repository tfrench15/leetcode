package answers

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

package answers

func maxSubArray(nums []int) int {
	var (
		max   = minimum(nums)
		count int
	)
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j <= len(nums); j++ {
			if i == j {
				count = nums[j]
			} else {
				count = sum(nums[i:j])
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}

func minimum(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func sum(nums []int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	return count
}

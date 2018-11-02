package answers

func twoSum(nums []int, target int) []int {
	var ret []int
	idx := 0
	for idx < len(nums)-1 {
		for i, num := range nums {
			if idx == i {
				continue
			}
			if nums[idx]+num == target {
				ret = append(ret, idx, i)
				return ret
			}
		}
		idx++
	}
	return nil
}

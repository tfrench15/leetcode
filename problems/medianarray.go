package answers

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	arr := joinSortedArrays(nums1, nums2)
	med := findMedian(arr)
	return med
}

func findMedian(nums []int) float64 {
	if len(nums)%2 == 0 {
		idx := len(nums) / 2
		n1 := nums[idx-1]
		n2 := nums[idx]
		med := (float64(n1) + float64(n2)) / 2.0
		return med
	}
	n1 := nums[len(nums)/2]
	return float64(n1)
}

func joinSortedArrays(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	var ret []int
	i, j := 0, 0
	for {
		if i == len(nums1) && j == len(nums2) {
			break
		}
		if i == len(nums1) && j < len(nums2) {
			ret = append(ret, nums2[j])
			j++
			continue
		}
		if i < len(nums1) && j == len(nums2) {
			ret = append(ret, nums1[i])
			i++
			continue
		} else {
			if nums1[i] >= nums2[j] {
				ret = append(ret, nums2[j])
				j++
				continue
			}
			if nums1[i] < nums2[j] {
				ret = append(ret, nums1[i])
				i++
				continue
			}
		}
	}
	return ret
}

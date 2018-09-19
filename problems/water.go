package answers

func maxArea(heights []int) int {
	var (
		area    int
		maxArea = 0
	)
	for b1, h1 := range heights {
		for b2, h2 := range heights {
			base := b2 - b1
			height := minHeight(h1, h2)
			area = base * height
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func minHeight(h1, h2 int) int {
	if h1 <= h2 {
		return h1
	}
	return h2
}

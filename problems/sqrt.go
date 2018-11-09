package answers

import "math"

func mySqrt(num int) int {
	flt := float64(num)
	ans := math.Floor(math.Sqrt(flt))
	return int(ans)
}

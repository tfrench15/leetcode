package answers

func climbStairs(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return n
	}
	return fibonacci(n)
}

func fibonacci(n int) int {
	var next int
	cur, prev := 1, 1
	for i := 2; i <= n; i++ {
		next = cur + prev
		prev = cur
		cur = next
	}
	return next
}

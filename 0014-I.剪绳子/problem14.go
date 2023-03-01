package _14_剪绳子

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		current_max := 0
		for j := 1; j < i; j++ {
			current_max = max(current_max, max(j*(i-j), j*dp[i-j]))
		}

		dp[i] = current_max
	}

	return dp[n]
}

package dynamicProgramming

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	// dp := make([]int, n + 1)
	// dp[0], dp[1] = 0, 0
	prev, curr := 0, 0
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	for i := 2; i <= n; i++ {
		prev, curr = curr, min(curr+cost[i-1], prev+cost[i-2])
		// dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return curr
}

package dynamicProgramming

func min(nums ...int) int {
	res := nums[0]
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	first, second, third := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(2*dp[first], 3*dp[second], 5*dp[third])
	}
	return 0
}

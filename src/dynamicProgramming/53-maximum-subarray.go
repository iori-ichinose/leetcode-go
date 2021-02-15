package dynamicProgramming

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	res := nums[0]
	for i := 1; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

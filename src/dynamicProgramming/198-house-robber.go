/*
 * https://leetcode-cn.com/problems/house-robber/
 * 2020.11.24
 * Golang 0ms(100.00%), 2MB(16.97%)
 */
package dynamicProgramming

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

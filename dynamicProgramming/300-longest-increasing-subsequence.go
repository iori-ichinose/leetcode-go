/*
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/
 * 2020.12.18
 * Golang 4ms(82.87%), 3.5MB(18.92%)
 */
package dynamicProgramming

// O(n^2)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	res := 1
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// O(nlogn)
func lengthOfLIS2(nums []int) int {
	length, n := 1, len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[length] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > dp[length] {
			length++
			dp[length] = nums[i]
		} else {
			l, r, pos := 1, length, 0
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[pos+1] = nums[i]
		}
	}
	return length
}

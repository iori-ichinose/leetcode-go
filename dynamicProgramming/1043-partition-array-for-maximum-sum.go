/*
 * https://leetcode-cn.com/problems/partition-array-for-maximum-sum/
 * 2020.12.3
 * Golang 4ms(96.43%), 2.8MB(28.57%)
 */
package dynamicProgramming

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i <= n; i++ {
		maxN := -1
		for j := i - 1; j >= max(i-k, 0); j-- {
			maxN = max(maxN, arr[j])
			dp[i] = max(dp[i], dp[j]+maxN*(i-j))
		}
	}
	return dp[n]
}

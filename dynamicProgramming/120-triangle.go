/*
 * https://leetcode-cn.com/problems/triangle/
 * 2020.11.25
 * Golang 4ms(94.76%), 3MB(83.54%)
 */
package dynamicProgramming

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	dp := make([][]int, n)
	dp[0] = make([]int, 1)
	dp[0][0] = 1

	for i := 1; i < n; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	res := 1<<31 - 1
	for _, n := range dp[n-1] {
		res = min(res, n)
	}
	return res
}

func minimumTotalSimplified(triangle [][]int) int {
	n := len(triangle)
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	dp := make([]int, n)
	dp[0] = triangle[0][0]

	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}

	res := 1<<31 - 1
	for _, n := range dp {
		res = min(res, n)
	}
	return res
}

/*
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 * 2020.11.29
 * Golang 4ms(96.29%), 3MB(74.12%)
 */
package dynamicProgramming

func maxProfitDP(prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i][1], -prices[i])
	}
	return dp[n-1][0]
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	minPrice, maxPrice := prices[0], 0
	for i := 1; i < n; i++ {
		maxPrice = max(maxPrice, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return maxPrice
}

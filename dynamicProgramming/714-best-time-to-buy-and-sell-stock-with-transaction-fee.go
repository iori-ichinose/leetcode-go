/*
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
 * 2020.12.17
 * Golang 104ms(82.74%), 7.8MB(62.25%)
 */
package dynamicProgramming

func maxProfitWithFee(prices []int, fee int) int {
	m := len(prices)
	if m == 0 || m == 1 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	dp := [2]int{0, -prices[0]}
	for i := 1; i < m; i++ {
		dp[0], dp[1] = max(dp[0], dp[1]+prices[i]-fee), max(dp[1], dp[0]-prices[i])
	}

	return dp[0]
}

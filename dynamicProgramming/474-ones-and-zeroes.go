/*
 * https://leetcode-cn.com/problems/ones-and-zeroes/
 * 2020.11.30
 * Golang 32ms(69.64%), 3.7MB(47.35%)
 */
package dynamicProgramming

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	count := func(s string) []int {
		c := make([]int, 2)
		for i := 0; i < len(s); i++ {
			c[s[i]-'0']++
		}
		return c
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	for _, ch := range strs {
		c := count(ch)
		for zero := m; zero >= c[0]; zero-- {
			for one := n; one >= c[1]; one-- {
				dp[zero][one] = max(1+dp[zero-c[0]][one-c[1]], dp[zero][one])
			}
		}
	}
	return dp[m][n]
}

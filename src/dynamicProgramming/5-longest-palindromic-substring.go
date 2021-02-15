/*
 * https://leetcode-cn.com/problems/longest-palindromic-substring/
 * 2020.11.30
 * 动态规划：
 * Golang 168ms(22.78%), 7.2MB(20.47%)
 */
package dynamicProgramming

func longestPalindrome(s string) string {
	n, res := len(s), ""
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n-i)
	}

	// 次序不可颠倒
	for j := 0; j < n; j++ {
		for i := 0; i < n-j; i++ {
			if 0 == j {
				dp[i][j] = true
			} else if 1 == j {
				dp[i][j] = s[i] == s[j+i]
			} else {
				dp[i][j] = s[i] == s[j+i] && dp[i+1][j-2]
			}
			if dp[i][j] && j+1 > len(res) {
				res = s[i : i+j+1]
			}
		}
	}
	return res
}

func LongestPalindromeTb() {
	print("Testbench longestPalindrome: ")
	println(longestPalindrome("babad"))
}

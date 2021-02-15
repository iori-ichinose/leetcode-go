/*
 * https://leetcode-cn.com/problems/longest-palindrome/
 * 2020.11.24
 * Golang 0ms(100.00%), 2.1MB(31.25%)
 */
package qstring

func longestPalindrome(s string) int {
	count, res := map[rune]int{}, 0
	for _, n := range s {
		count[n]++
	}
	for _, n := range count {
		if n&1 == 0 {
			res += n
		} else {
			res += n - 1
			if res&1 == 0 {
				res++
			}
		}
	}
	return res
}

/*
 * https://leetcode-cn.com/problems/palindrome-permutation-lcci/
 * 2020.11.29
 * Golang 0ms(100.00%), 1.9MB(69.05%)
 */
package qstring

func canPermutePalindrome(s string) bool {
	count := map[rune]int{}
	for _, ch := range s {
		if count[ch] != 0 {
			count[ch] = 0
		} else {
			count[ch] = 1
		}
	}

	size := 0
	for _, n := range count {
		size += n
	}
	return size <= 1
}

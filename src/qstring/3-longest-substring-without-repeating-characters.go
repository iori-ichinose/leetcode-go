/*
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 * 2020.11.25
 * Golang 0ms(100.00%), 2.5MB(87.10%)
 */
package qstring

func lengthOfLongestSubstring(s string) int {
	// appear存放出现过的字符，因为只存在空格、数字、字母，所以256（ASCII）足够
	appear, n := [256]int{}, len(s)
	// 左右指针
	rightPtr, res := -1, 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	for leftPtr := 0; leftPtr < n; leftPtr++ {
		// 左指针右移时，其左侧不再出现，因此从appear移除
		if leftPtr != 0 {
			appear[s[leftPtr-1]] = 0
		}
		// 右指针不断右移，直到遇到重复的字符，停止
		for rightPtr < n-1 && appear[s[1+rightPtr]] == 0 {
			appear[s[rightPtr+1]]++
			rightPtr++
		}
		// 对于每个左指针能得到的最长长度都比较，取最大值
		res = max(res, rightPtr-leftPtr+1)
	}
	return res
}

/*
 * https://leetcode-cn.com/problems/wiggle-subsequence/
 * 2020.12.12
 * Golang 0ms(100.00%), 2.1MB(72.41%)
 */
package greedy

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	up, down := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else {
			down = up + 1
		}
	}
	if up > down {
		return up
	} else {
		return down
	}
}

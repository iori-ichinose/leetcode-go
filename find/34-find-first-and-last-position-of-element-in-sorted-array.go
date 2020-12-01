/*
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
 * 2020.11.30
 * Golang 8ms(90.40%), 3.9MB(100.00%)
 */
package find

import "sort"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	a, b := sort.SearchInts(nums, target), sort.SearchInts(nums, target+1)
	if a == n || nums[a] != target {
		return []int{-1, -1}
	} else {
		return []int{a, b - 1}
	}
}

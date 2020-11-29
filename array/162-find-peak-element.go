/*
 * https://leetcode-cn.com/problems/find-peak-element/
 * 2020.11.28
 * Golang 4ms(85.76%), 2.7MB(44.48%)
 */
package array

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

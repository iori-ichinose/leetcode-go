/*
 * https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/
 * 2020.11.22
 * Golang 0ms(100%), 2.1MB(66.67%)
 */
package backtrack

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) == 0 {
		return true
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	var backTrack func([]int) bool
	backTrack = func(groups []int) bool {
		if len(nums) == 0 {
			return true
		}
		v := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		for i, group := range groups {
			if group+v <= target {
				groups[i] += v
				if backTrack(groups) {
					return true
				}
				groups[i] -= v
			}
			if group == 0 {
				break
			}
		}
		nums = append(nums, v)
		return false
	}
	sort.Ints(nums)
	if nums[len(nums)-1] > target {
		return false
	}
	return backTrack(make([]int, k))
}

func CanPartitionKSubsetsTb() {

}

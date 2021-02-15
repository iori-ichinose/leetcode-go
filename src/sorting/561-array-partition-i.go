/*
 * https://leetcode-cn.com/problems/array-partition-i/
 * 2020.12.4
 * 方法一：排序
 * Golang 76ms(45.34%), 6.9MB(79.48%)
 */
package sorting

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	n, res := len(nums), 0
	for i := 0; i < n; i += 2 {
		res += nums[i]
	}
	return res
}

func arrayPairSum2(nums []int) int {
	var arr [20001]int
	lim := 10000
	for _, n := range nums {
		arr[n+lim]++
	}
	d, sum := 0, 0
	for i := -10000; i <= 10000; i++ {
		sum += (arr[i+lim] + 1 - d) / 2 * i
		d = (2 + arr[i+lim] - d) % 2
	}
	return sum
}

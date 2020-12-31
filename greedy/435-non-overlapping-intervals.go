/*
 * https://leetcode-cn.com/problems/non-overlapping-intervals/
 * 2020.12.31
 * Golang 4ms(100.00%), 3.9MB(62.50%)
 */
package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] > intervals[j][1]
	})

	right, res := intervals[0][1], 1
	for i := 1; i < n; i++ {
		if intervals[i][0] >= right {
			res++
			right = intervals[i][1]
		}
	}
	return n - res
}

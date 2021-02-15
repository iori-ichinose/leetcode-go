/*
 * 我TM射爆！！
 * https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/
 * 2020.11.23
 * Golang 80ms(91.27%), 7MB(71.20%)
 *
 * 本题的贪心思想体现在每次射箭我们总选择射爆最右方的气球（即最远处），进而得到最少的弓箭数量。
 */
package greedy

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	maxr, ans := points[0][1], 1
	for _, p := range points {
		if p[0] > maxr {
			maxr = p[1]
			ans++
		}
	}
	return ans
}

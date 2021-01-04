/**
 * https://leetcode-cn.com/problems/positions-of-large-groups/
 * 2021.1.5
 * Golang 0ms(100.00%), 2.6MB(100.00%)
 */
package array

func largeGroupPositions(s string) [][]int {
	var res [][]int
	currStart, currCount := 0, 0
	for i := range s {
		if i == 0 || s[i] == s[i-1] {
			currCount++
		} else {
			if currCount >= 3 {
				res = append(res, []int{currStart, currStart + currCount - 1})
			}
			currCount = 1
			currStart = i
		}
	}
	if currCount >= 3 {
		res = append(res, []int{currStart, currStart + currCount - 1})
	}
	return res
}

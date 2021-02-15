/*
 * https://leetcode-cn.com/problems/split-array-into-consecutive-subsequences/
 * 2020.12.4
 * Golang 100ms(65.63%), 7.1MB(40.63%)
 */
package greedy

func isPossible(nums []int) bool {
	left := map[int]int{}
	for _, val := range nums {
		left[val]++
	}
	endCount := map[int]int{}
	for _, val := range nums {
		if left[val] == 0 {
			continue
		}
		if endCount[val-1] > 0 {
			left[val]--
			endCount[val-1]--
			endCount[val]++
		} else if left[val+1] > 0 && left[val+2] > 0 {
			left[val]--
			left[val+1]--
			left[val+2]--
			endCount[val+2]++
		} else {
			return false
		}
	}
	return true
}

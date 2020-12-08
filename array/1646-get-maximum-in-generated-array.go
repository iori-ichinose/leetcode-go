/*
 * https://leetcode-cn.com/problems/get-maximum-in-generated-array/
 * 2020.12.8
 * Golang 0ms(100.00%), 2MB(89.66%)
 */
package array

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	target := make([]int, n+1)
	target[0], target[1] = 0, 1
	maxVal := target[1]
	for i := 1; i <= n/2; i++ {
		target[2*i] = target[i]
		if target[2*i] > maxVal {
			maxVal = target[2*i]
		}
		if i == n/2 && n%2 == 0 {
			break
		}
		target[2*i+1] = target[i] + target[i+1]
		if target[2*i+1] > maxVal {
			maxVal = target[2*i+1]
		}
	}
	return maxVal
}

/*
 * https://leetcode-cn.com/problems/largest-perimeter-triangle/
 * 2020.11.29
 * Golang 48ms(79.79%), 6.6MB(25.81%)
 */
package greedy

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i-1]+A[i-2] > A[i] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}

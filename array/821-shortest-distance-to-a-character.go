/*
 * https://leetcode-cn.com/problems/shortest-distance-to-a-character/
 * 2020.12.13
 * Golang 0ms(100.00%), 2.4MB(100.00%)
 */
package array

func shortestToChar(S string, C byte) []int {
	prev, n := -1<<31, len(S)
	res := make([]int, n)

	for i := range S {
		if S[i] == C {
			prev = i
		}
		res[i] = i - prev
	}

	prev = 1<<31 - 1
	for i := n - 1; i >= 0; i-- {
		if S[i] == C {
			prev = i
		}
		if prev-i < res[i] {
			res[i] = prev - i
		}
	}

	return res
}

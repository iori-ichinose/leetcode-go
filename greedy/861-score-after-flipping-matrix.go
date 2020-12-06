/*
 * https://leetcode-cn.com/problems/score-after-flipping-matrix/
 * 2020.12.7
 * Golang 0ms(100.00%), 2.2MB(81.25%)
 */
package greedy

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	res := m * 1 << (n - 1)
	for j := 1; j < n; j++ {
		ones := 0
		for i := 0; i < m; i++ {
			if A[i][0] == 1 {
				ones += A[i][j]
			} else {
				ones += 1 - A[i][j]
			}
		}
		if m-ones > ones {
			ones = m - ones
		}
		res += (1 << (n - j - 1)) * ones
	}
	return res
}

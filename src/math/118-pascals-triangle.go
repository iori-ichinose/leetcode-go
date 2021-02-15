/*
 * https://leetcode-cn.com/problems/pascals-triangle/
 * 2020.12.6
 * Golang 0ms(100.00%), 2MB(59.85%)
 */
package math

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

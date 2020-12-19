/*
 * https://leetcode-cn.com/problems/rotate-image/
 * 2020.12.19
 * Golang 0ms(100.00%), 2.2MB(58.74%)
 */
package array

func rotate(matrix [][]int) {
	row := len(matrix)

	for i := 0; i < row; i++ {
		for j := i; j < row; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < row/2; j++ {
			matrix[i][j], matrix[i][row-j-1] = matrix[i][row-j-1], matrix[i][j]
		}
	}
}

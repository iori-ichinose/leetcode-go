/*
 * https://leetcode-cn.com/problems/maximum-non-negative-product-in-a-matrix/
 * 2020.12.16
 * Golang 4ms(91.30%), 2.8MB(49.28%)
 */
package dynamicProgramming

func maxProductPath(grid [][]int) int {
	const mod = 1e9 + 7
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	m, n := len(grid), len(grid[0])
	maxgt, minlt := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		maxgt[i] = make([]int, n)
		minlt[i] = make([]int, n)
	}

	maxgt[0][0], minlt[0][0] = grid[0][0], grid[0][0]
	for i := 1; i < n; i++ {
		maxgt[0][i] = maxgt[0][i-1] * grid[0][i]
		minlt[0][i] = maxgt[0][i]
	}
	for i := 1; i < m; i++ {
		maxgt[i][0] = maxgt[i-1][0] * grid[i][0]
		minlt[i][0] = maxgt[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] >= 0 {
				maxgt[i][j] = max(maxgt[i][j-1], maxgt[i-1][j]) * grid[i][j]
				minlt[i][j] = min(minlt[i][j-1], minlt[i-1][j]) * grid[i][j]
			} else {
				maxgt[i][j] = min(minlt[i][j-1], minlt[i-1][j]) * grid[i][j]
				minlt[i][j] = max(maxgt[i][j-1], maxgt[i-1][j]) * grid[i][j]
			}
		}
	}
	if maxgt[m-1][n-1] < 0 {
		return -1
	} else {
		return maxgt[m-1][n-1] % mod
	}
}

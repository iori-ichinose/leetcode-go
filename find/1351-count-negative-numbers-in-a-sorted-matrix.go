package find

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0

	for i := m - 1; i >= 0 && grid[i][n-1] < 0; i-- {
		for j := n - 1; j >= 0 && grid[i][j] < 0; j-- {
			res++
		}
	}
	return res
}

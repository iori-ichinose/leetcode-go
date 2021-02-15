/*
 * https://leetcode-cn.com/problems/path-with-maximum-gold/
 * 2020.11.19
 * Golang 16ms(85%), 2.1MB(95%)
 */
package backtrack

func getMaximumGold(grid [][]int) int {
	maxGold := 0
	rm, cm := len(grid), len(grid[0])

	var backTrack func(i, j int, sum int)
	backTrack = func(row, col int, sum int) {
		if row < 0 || row >= rm || col < 0 || col >= cm || grid[row][col] == 0 {
			if sum > maxGold {
				maxGold = sum
			}
			return
		}

		tmp := grid[row][col]
		sum += tmp

		grid[row][col] = 0
		backTrack(row-1, col, sum)
		backTrack(row+1, col, sum)
		backTrack(row, col-1, sum)
		backTrack(row, col+1, sum)
		grid[row][col] = tmp
	}
	for i := 0; i < rm; i++ {
		for j := 0; j < cm; j++ {
			if grid[i][j] != 0 {
				backTrack(i, j, 0)
			}
		}
	}
	return maxGold
}

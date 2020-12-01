/*
 * https://leetcode-cn.com/problems/unique-paths-ii/
 * 2020.11.30
 * Golang 0ms(100.00%), 2.5MB(51.60%)
 */
package dynamicProgramming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 0 && dp[i-1][0] != 0 {
			dp[i][0] = 1
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && dp[0][j-1] != 0 {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

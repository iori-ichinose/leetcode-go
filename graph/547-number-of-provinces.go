/**
 * https://leetcode-cn.com/problems/number-of-provinces/
 * 2021.1.7
 * Golang 24ms(92.79%), 6.6MB(96.74%)
 */
package graph

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	res := 0
	var dfs func(from int)
	dfs = func(from int) {
		visited[from] = true
		for to, connection := range isConnected[from] {
			if connection == 1 && !visited[to] {
				dfs(to)
			}
		}
	}

	for i, v := range visited {
		if !v {
			res++
			dfs(i)
		}
	}
	return res
}

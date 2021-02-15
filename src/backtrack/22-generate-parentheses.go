/*
 * https://leetcode-cn.com/problems/generate-parentheses/
 * 2020.11.29
 * Golang 0ms(100.00%), 2.7MB(85.70%)
 */
package backtrack

func generateParenthesis(n int) []string {
	var res []string
	var backTrack func(left, right int, path []byte)
	backTrack = func(left, right int, path []byte) {
		if left == right && left == n {
			res = append(res, string(path))
		}

		if left < n {
			path = append(path, '(')
			backTrack(left+1, right, path)
			path = path[:len(path)-1]
		}

		if right < left {
			path = append(path, ')')
			backTrack(left, right+1, path)
			path = path[:len(path)-1]
		}
	}
	backTrack(0, 0, []byte{})
	return res
}

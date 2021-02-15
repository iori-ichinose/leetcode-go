/*
 * https://leetcode-cn.com/problems/count-and-say/
 * 2020.11.29
 * Golang 0ms(100.00%), 2.4MB(70.15%)
 */
package qstring

func countAndSay(n int) string {
	var prev []byte
	curr := []byte{'1'}
	for i := 1; i < n; i++ {
		prev = curr
		curr = []byte{}
		start, end := 0, 0
		for end < len(prev) {
			for end < len(prev) && prev[start] == prev[end] {
				end++
			}
			curr = append(curr, byte(end-start+'0'), prev[start])
			start = end
		}
	}
	return string(curr)
}

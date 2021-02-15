/*
 * https://leetcode-cn.com/problems/remove-duplicate-letters/
 * 2020.12.20
 * Golang 0ms(100.00%), 2.1MB(98.74%)
 */
package qstring

func removeDuplicateLetters(s string) string {
	n := len(s)
	var lastIndex [26]int
	var visited [26]bool
	var stack []byte

	for i := 0; i < n; i++ {
		lastIndex[s[i]-'a'] = i
	}

	for i := 0; i < n; i++ {
		if visited[s[i]-'a'] {
			continue
		}

		for last := len(stack) - 1; last >= 0 && stack[last] > s[i] && lastIndex[stack[last]-'a'] > i; last = len(stack) - 1 {
			top := stack[last]
			stack = stack[:last]
			visited[top-'a'] = false
		}

		stack = append(stack, s[i])
		visited[s[i]-'a'] = true
	}

	return string(stack)
}

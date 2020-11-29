/*
 * https://leetcode-cn.com/problems/reorganize-string/
 * 2020.11.30
 * Golang 0ms(100.00%), 2MB(95.35%)
 */
package qstring

func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}

	var count [26]int
	maxCount := 0
	for _, ch := range S {
		count[ch-'a']++
		if count[ch-'a'] > maxCount {
			maxCount = count[ch-'a']
		}
	}
	if maxCount > (n+1)/2 {
		return ""
	}

	res := make([]byte, n)
	even, odd, half := 0, 1, n/2

	for i, c := range count {
		ch := byte(i + 'a')
		for c > 0 && c <= half && odd < n {
			res[odd] = ch
			c--
			odd += 2
		}
		for c > 0 {
			res[even] = ch
			c--
			even += 2
		}
	}
	return string(res)
}

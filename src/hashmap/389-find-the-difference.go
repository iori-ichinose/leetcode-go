/*
 * https://leetcode-cn.com/problems/find-the-difference/
 * 2020.12.18
 * Golang 0ms(100.00%), 2.1MB(77.31%)
 */
package hashmap

func findTheDifference(s string, t string) byte {
	var count [26]int
	for _, ch := range s {
		count[ch-'a']++
	}
	for _, ch := range t {
		count[ch-'a']--
		if count[ch-'a'] < 0 {
			return byte(ch)
		}
	}
	return ' '
}

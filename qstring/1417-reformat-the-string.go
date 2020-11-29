/*
 * https://leetcode-cn.com/problems/reformat-the-string/
 * 2020.11.29
 * Golang 4ms(67.50%), 3.8MB(37.50%)
 *
 * 怎么写都别扭，我太难了...
 */
package qstring

import "unicode"

func reformat(s string) string {
	var digits, letters []rune
	var res []byte
	for _, n := range s {
		if unicode.IsDigit(n) {
			digits = append(digits, n)
		} else {
			letters = append(letters, n)
		}
	}
	m, n := len(digits), len(letters)
	if m > n+1 || n > m+1 {
		return ""
	} else if m > n {
		res = append(res, byte(digits[0]))
		for i := 0; i < n; i++ {
			res = append(res, byte(letters[i]))
			res = append(res, byte(digits[i+1]))
		}
	} else if m < n {
		res = append(res, byte(letters[0]))
		for i := 0; i < m; i++ {
			res = append(res, byte(digits[i]))
			res = append(res, byte(letters[i+1]))
		}
	} else {
		for i := 0; i < n; i++ {
			res = append(res, byte(digits[i]))
			res = append(res, byte(letters[i]))
		}
	}
	return string(res)
}

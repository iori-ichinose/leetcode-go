/*
 * https://leetcode-cn.com/problems/maximum-number-of-balloons/
 * 2020.12.15
 * Golang 0ms(100.00%), 2.2MB(24.59%)
 */
package greedy

import "strconv"

func monotoneIncreasingDigits(N int) int {
	numArray := []byte(strconv.Itoa(N))
	n, i := len(numArray), 1
	for i < n && numArray[i-1] <= numArray[i] {
		i++
	}
	if i == n {
		return N
	}

	for i > 0 && numArray[i-1] > numArray[i] {
		numArray[i-1]--
		i--
	}
	for i = i + 1; i < n; i++ {
		numArray[i] = '9'
	}
	res, _ := strconv.Atoi(string(numArray))
	return res
}

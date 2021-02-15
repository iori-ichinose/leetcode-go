/*
 * https://leetcode-cn.com/problems/max-consecutive-ones-iii/
 * 2020.12.9
 * Golang 72ms(83.05%), 6.9MB(62.61%)
 */
package slideWindow

func longestOnes(A []int, K int) int {
	left, right, count, length := 0, 0, 0, len(A)
	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	for right < length {
		if A[right] == 0 {
			count++
		}
		right++

		if count > K {
			if A[left] == 0 {
				count--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

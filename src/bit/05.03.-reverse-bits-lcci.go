/*
 * https://leetcode-cn.com/problems/reverse-bits-lcci/
 * 2020.11.25
 * Golang 0ms(100.00%), 1.9MB(60.00%)
 */
package bit

func reverseBits(num int) int {
	if ^num == 0 {
		return 32
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	maxlength, pre, cur := 0, 0, 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			cur++
		} else {
			pre = cur
			cur = 0
		}
		maxlength = max(maxlength, pre+cur+1)
		num >>= 1
	}
	return maxlength
}

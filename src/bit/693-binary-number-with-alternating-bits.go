/*
 * https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
 * 2020.11.22
 * Golang 0ms(100%), 1.9MB(66.00%)
 */
package bit

func hasAlternatingBits(n int) bool {
	// 右移保证二进制位数不变
	s := n >> 1
	// 如果是01交替，相互异或结果为全1
	s = s ^ n
	// 二进制全1 + 1 -> 除去首位全部为0。
	// 此时与s按位取与，结果必然为0
	return s&(s+1) == 0
}

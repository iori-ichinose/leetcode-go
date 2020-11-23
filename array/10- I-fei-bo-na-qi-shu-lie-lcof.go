/*
 * https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
 * 2020.11.21
 * Golang 0ms(100.00%), 1.9MB(91.51%)
 */
package array

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	const MOD int = 1000000007
	a, b := 1, 0
	for i := 1; i < n; i++ {
		a = a + b
		b = a - b
		a %= MOD
	}
	return a
}

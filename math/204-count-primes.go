/*
 * https://leetcode-cn.com/problems/count-primes/
 * 2020.12.3
 * Golang 8ms(96.57%), 4.9MB(64.12%)
 */
package math

// 埃氏筛法
// 对于一个数n，如果它是质数，则2*n, 3*n, 4*n一定不是质数
// 同时我们标记只需从n*n开始，因为任意i*n (i < n)必然已经被之前的数标记过。
func countPrimes(n int) int {
	notPrime, res := make([]bool, n), 0
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			res++
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}

	return res
}

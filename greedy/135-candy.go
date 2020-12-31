/*
 * https://leetcode-cn.com/problems/candy/
 * 2020.12.24
 * Golang 16ms(97.05%), 6.1MB(100.00%)
 */
package greedy

func candy(ratings []int) int {
	n := len(ratings)
	res, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			res += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			res += dec
			pre = 1
		}
	}

	return res
}

/*
 * https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
 * 2020.11.27
 * 动态规划: Golang 72ms(46.73%), 16.3MB(22.43%)
 * 滑动窗口：Golang 72ms(46.73%), 3.4MB(96.26%)
 */
package array

func findLengthDp(A []int, B []int) int {
	m, n, res := len(A), len(B), 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	dp := make([][]int, n+1)
	dp[n] = make([]int, m+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]int, m+1)
		for j := m - 1; j >= 0; j-- {
			if A[i] != B[j] {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i+1][j+1] + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

func findLengthSlideWindows(A []int, B []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	maxLength := func(addA, addB, len int) int {
		res, k := 0, 0
		for i := 0; i < len; i++ {
			if A[addA+i] == B[addB+i] {
				k++
			} else {
				k = 0
			}
			res = max(res, k)
		}
		return res
	}

	res, n, m := 0, len(A), len(B)
	for i := 0; i < n; i++ {
		length := min(m, n-i)
		maxLen := maxLength(i, 0, length)
		res = max(res, maxLen)
	}

	for i := 0; i < n; i++ {
		length := min(n, m-i)
		maxLen := maxLength(0, i, length)
		res = max(res, maxLen)
	}
	return res
}

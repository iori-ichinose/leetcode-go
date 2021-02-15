/*
 * https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/
 * 2020.12.8
 * Golang 0ms(100.00%), 2.1MB(80.39%)
 */
package backtrack

import "math"

func splitIntoFibonacci(S string) []int {
	n := len(S)
	var res []int
	var backtrack func(index, sum, prev int) bool
	backtrack = func(index, sum, prev int) bool {
		if index == n {
			return len(res) >= 3
		}
		curr := 0
		for i := index; i < n; i++ {
			if i > index && S[index] == '0' {
				break
			}
			curr = 10*curr + int(S[i]-'0')
			if curr > math.MaxInt32 {
				break
			}

			if len(res) >= 2 {
				if curr < sum {
					continue
				} else if curr > sum {
					break
				}
			}

			res = append(res, curr)
			if backtrack(i+1, prev+curr, curr) {
				return true
			}
			res = res[:len(res)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return res
}

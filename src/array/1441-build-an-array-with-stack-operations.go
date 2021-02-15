/*
 * https://leetcode-cn.com/problems/build-an-array-with-stack-operations/
 * 2020.12.16
 * Golang 0ms(100.00%), 2.4MB(10.26%)
 */
package array

func buildArray(target []int, n int) []string {
	var res []string
	set := map[int]bool{}
	for _, num := range target {
		set[num] = true
	}

	for i := 1; i < target[len(target)-1]; i++ {
		res = append(res, "Push")
		if !set[i] {
			res = append(res, "Pop")
		}
	}
	return res
}

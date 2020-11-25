/*
 * https://leetcode-cn.com/problems/increasing-decreasing-string/
 * 2020.11.25
 * Golang 0ms(100.00%), 3MB(91.35%)
 */
package qstring

func sortString(s string) string {
	count, n := [26]int{}, len(s)
	res := make([]byte, 0, n)
	for _, n := range s {
		count[n-'a']++
	}
	for len(res) < n {
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				res = append(res, byte(i)+'a')
				count[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if count[i] > 0 {
				res = append(res, byte(i)+'a')
				count[i]--
			}
		}
	}
	return string(res)
}

func SortStringTb() {
	print("Testbench sortString: ")
	s := sortString("leetcode")
	println(s)
}

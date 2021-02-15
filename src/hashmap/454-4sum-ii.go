/*
 * https://leetcode-cn.com/problems/4sum-ii/
 * 2020.11.27
 * Golang 64ms(73.51%), 22.6MB(47.26%)
 */
package hashmap

func fourSumCount(A []int, B []int, C []int, D []int) int {
	count, res := map[int]int{}, 0
	// 遍历A，B将所有可能的组合以及他们的数量放入Hashmap
	for _, n := range A {
		for _, m := range B {
			count[m+n]++
		}
	}
	// 遍历C，D，所有的组合中如果出现了-A-B，说明和必然为0，则等于0的数量增加。
	for _, u := range C {
		for _, v := range D {
			res += count[-u-v]
		}
	}
	return res
}

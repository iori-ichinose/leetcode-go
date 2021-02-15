/*
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/
 * 2020.11.29
 * Golang 8ms(60.43%), 3.8MB(27.01%)
 */
package hashmap

func longestConsecutive(nums []int) int {
	lengths := map[int]int{}
	res := 0

	for _, n := range nums {
		if lengths[n] == 0 {
			// 如果Hashmap中没有n，则进行比较，否则无需比较
			// 因为假如已经存在了n，则n已经参与计算过一遍，无需再次计算
			left, right := lengths[n-1], lengths[n+1]

			// 左右如果存在已有元素，则新的长度为左右已经形成的序列长度和+1
			cur := 1 + left + right
			if cur > res {
				res = cur
			}

			// 每一次放入时因为n不在其中，所以一个新的n只有可能和Hashmap原有的序列边界重合
			// 因此我们只需把序列的两端点设置为序列长度即可
			lengths[n] = cur
			lengths[n-left] = cur
			lengths[n+right] = cur
		}
	}
	return res
}

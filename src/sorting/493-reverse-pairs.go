/*
 * https://leetcode-cn.com/problems/reverse-pairs/
 * 2020.11.28
 * Golang 136ms(67.53%), 8.1MB(61.04%)
 */
package sorting

func reversePairsRecur(nums []int, left, right int) int {
	if left == right {
		return 0
	}

	mid := (left + right) / 2
	n1 := reversePairsRecur(nums, left, mid)
	n2 := reversePairsRecur(nums, mid+1, right)
	res := n1 + n2

	i, j := left, mid+1
	for i <= mid {
		for j <= right && nums[i] > 2*nums[j] {
			j++
		}
		res += j - mid - 1
		i++
	}

	sorted := make([]int, right-left+1)
	p1, p2, p := left, mid+1, 0
	for p1 <= mid || p2 <= right {
		if p1 > mid {
			sorted[p] = nums[p2]
			p++
			p2++
		} else if p2 > right {
			sorted[p] = nums[p1]
			p++
			p1++
		} else if nums[p1] < nums[p2] {
			sorted[p] = nums[p1]
			p++
			p1++
		} else {
			sorted[p] = nums[p2]
			p++
			p2++
		}
	}

	for i := 0; i < len(sorted); i++ {
		nums[left+i] = sorted[i]
	}
	return res
}

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return reversePairsRecur(nums, 0, len(nums)-1)
}

package array

func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}

	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}

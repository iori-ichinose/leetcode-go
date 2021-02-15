package greedy

func minPatches(nums []int, n int) int {
	res, i, x := 0, 0, 1
	for x <= n {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			x *= 2
			res++
		}
	}
	return res
}

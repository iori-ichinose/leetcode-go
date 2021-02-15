/*
 * https://leetcode-cn.com/problems/jump-game/
 * 2020.12.4
 * Golang 8ms(90.74%), 4MB(100.00%)
 */
package greedy

func canJump(nums []int) bool {
	// 记录每一次能跳到最远处
	n, rightmost := len(nums), 0
	for i := 0; i < n; i++ {
		// 假如i在最远处左侧，说明一定能跳到此位置
		// 若在i处能跳到比最远处还要远，更新rightmost
		// 假如跳不到i，那肯定也无法跳到i之后的了，返回false
		if i <= rightmost {
			if nums[i]+i > rightmost {
				rightmost = nums[i] + i
				if rightmost >= n-1 {
					return true
				}
			}
		} else {
			return false
		}
	}
	return true
}

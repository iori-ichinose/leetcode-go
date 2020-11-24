/*
 * https://leetcode-cn.com/problems/count-complete-tree-nodes/
 * 2020.11.24
 * Golang 16ms(92.44%), 7.1MB(80.11%)
 */
package binaryTree

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level, node := 0, root
	for node.Left != nil {
		level++
		node = node.Left
	}
	low, high := 1<<level, (1<<(level+1))-1
	exists := func(root *TreeNode, level, k int) bool {
		bits, node := 1<<(level-1), root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node != nil
	}
	for low < high {
		mid := (high-low+1)/2 + low
		if exists(root, level, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

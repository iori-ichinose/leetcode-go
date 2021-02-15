/*
 * https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
 * 2020.11.22
 * Golang 4ms(90.00%), 4.4MB(77.47%)
 */
package binaryTree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left, right := maxDepth(root.Left)+1, maxDepth(root.Right)+1
		if left > right {
			return left
		} else {
			return right
		}
	}
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, res := []*TreeNode{root}, 0
	for len(queue) != 0 {
		res++
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

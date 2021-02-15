/*
 * https://leetcode-cn.com/problems/symmetric-tree/
 * 2020.11.22
 * Golang 0ms(100.00%), 2.9MB(59.26%)
 */
package binaryTree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var check func(node1, node2 *TreeNode) bool
	check = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		} else if node1 == nil || node2 == nil {
			return false
		} else if node1.Val != node2.Val {
			return false
		} else {
			return check(node1.Left, node2.Right) && check(node1.Right, node2.Left)
		}
	}
	return check(root.Left, root.Right)
}

/*
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
 * 2020.11.27
 * Golang 0ms(100.00%), 2.8MB(75.84%)
 */
package binaryTree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		var p []*TreeNode
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}

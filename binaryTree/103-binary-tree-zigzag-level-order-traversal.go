/*
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
 * 2020.12.22
 * Golang 0ms(100.00%), 2.5MB(47.17%)
 */
package binaryTree

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	q := []*TreeNode{root}
	for i := 0; len(q) != 0; i++ {
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
		if i%2 == 1 {
			for j, n := 0, len(res[i]); j < n/2; j++ {
				res[i][j], res[i][n-1-j] = res[i][n-1-j], res[i][j]
			}
		}
		q = p
	}

	return res
}

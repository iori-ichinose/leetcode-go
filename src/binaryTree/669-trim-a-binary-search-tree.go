/*
 * https://leetcode-cn.com/problems/trim-a-binary-search-tree/
 * 2020.11.19
 * Golang 8ms(98%), 6.2MB(86%)
 */
package binaryTree

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if root.Val < low {
			return dfs(root.Right)
		} else if root.Val > high {
			return dfs(root.Left)
		} else {
			root.Left = dfs(root.Left)
			root.Right = dfs(root.Right)
		}
		return root
	}

	return dfs(root)
}

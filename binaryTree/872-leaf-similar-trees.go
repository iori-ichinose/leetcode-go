/*
 * https://leetcode-cn.com/problems/leaf-similar-trees/
 * 2020.12.9
 * Golang 0ms(100.00%), 2.6MB(48.15%)
 */
package binaryTree

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var tree1, tree2 []int
	var dfs func(root *TreeNode, dst *[]int)
	dfs = func(root *TreeNode, dst *[]int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			*dst = append(*dst, root.Val)
		}
		dfs(root.Left, dst)
		dfs(root.Right, dst)
	}
	dfs(root1, &tree1)
	dfs(root2, &tree2)

	if len(tree2) != len(tree1) {
		return false
	}
	for i := range tree1 {
		if tree1[i] != tree2[i] {
			return false
		}
	}
	return true
}

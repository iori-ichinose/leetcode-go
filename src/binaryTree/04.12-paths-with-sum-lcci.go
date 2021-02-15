package binaryTree

func pathSum(root *TreeNode, sum int) int {
	res := 0
	var dfs func(root *TreeNode, currLeft int)
	dfs = func(root *TreeNode, currLeft int) {
		if root == nil || currLeft < root.Val {
			return
		}

		if currLeft == root.Val {
			res++
			return
		}

		dfs(root.Left, currLeft-root.Val)
		dfs(root.Right, currLeft-root.Val)
	}

	var search func(root *TreeNode)
	search = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root, sum)
		search(root.Left)
		search(root.Right)
	}
	search(root)
	return res
}

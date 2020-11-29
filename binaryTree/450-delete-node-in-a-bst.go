package binaryTree

func afterward(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func ancestor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if root.Left == nil && root.Right == nil {
		root = nil
	} else if root.Left != nil {
		root.Val = ancestor(root)
		root.Left = deleteNode(root.Left, root.Val)
	} else {
		root.Val = afterward(root)
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

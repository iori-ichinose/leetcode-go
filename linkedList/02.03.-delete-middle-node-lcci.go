/*
 * https://leetcode-cn.com/problems/delete-middle-node-lcci/
 * 2020.11.27
 * Golang 0ms(100.00%), 2.8MB(47.36%)
 *
 * 这特么什么鬼题？
 */
package linkedList

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

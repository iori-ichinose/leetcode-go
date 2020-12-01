/*
 * https://leetcode-cn.com/problems/middle-of-the-linked-list/
 * 2020.12.1
 * Golang 0ms(100.00%), 1.9MB(67.14%)
 */
package linkedList

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

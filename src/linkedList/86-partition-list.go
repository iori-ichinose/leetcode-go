/**
 * https://leetcode-cn.com/problems/partition-list/
 * 2021.1.3
 * Golang 0ms(100.00%), 2.4MB(68.33%)
 */
package linkedList

func partition(head *ListNode, x int) *ListNode {
	left, right := new(ListNode), new(ListNode)
	leftHead, rightHead := left, right
	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	left.Next = rightHead.Next
	right.Next = nil
	return leftHead.Next
}

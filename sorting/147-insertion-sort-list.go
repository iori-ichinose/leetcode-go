/*
 * https://leetcode-cn.com/problems/insertion-sort-list/
 * 2020.11.20
 * Golang 4ms(96.90%), 3.2MB(65.95%)
 */
package sorting

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := &ListNode{Next: head}
	sorted, cur := head, head.Next
	for cur != nil {
		if sorted.Val <= cur.Val {
			sorted = sorted.Next
		} else {
			prev := p
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}
			sorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = sorted.Next
	}
	return p.Next
}

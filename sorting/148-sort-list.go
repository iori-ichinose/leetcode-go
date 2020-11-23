/*
 * https://leetcode-cn.com/problems/sort-list/
 * 2020.11.21
 * 进阶：O(nlogn)时间复杂度，O(1)空间复杂度
 * Golang 36ms(40.18%), 6.9MB(40.84%)
 */
package sorting

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}
	mid := slow.Next
	slow.Next = nil

	left, right := sortList(head), sortList(mid)
	h := &ListNode{}
	res := h
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next, left = left, left.Next
		} else {
			h.Next, right = right, right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}
	return res.Next
}

func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	merge := func(head1, head2 *ListNode) *ListNode {
		phead := &ListNode{}
		p0, p1, p2 := phead, head1, head2
		for p1 != nil && p2 != nil {
			if p1.Val <= p2.Val {
				p0.Next = p1
				p1 = p1.Next
			} else {
				p0.Next = p2
				p2 = p2.Next
			}
			p0 = p0.Next
		}
		if p1 != nil {
			p0.Next = p1
		} else if p2 != nil {
			p0.Next = p2
		}
		return phead.Next
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	nhead := &ListNode{Next: head}
	for i := 1; i < length; i <<= 1 {
		pre, cur := nhead, nhead.Next
		for cur != nil {
			head1 := cur
			for j := 1; j < i && cur.Next != nil; j++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for j := 1; j < i && cur != nil; j++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			pre.Next = merge(head1, head2)

			for pre.Next != nil {
				pre = pre.Next
			}
			cur = next
		}
	}
	return nhead.Next
}

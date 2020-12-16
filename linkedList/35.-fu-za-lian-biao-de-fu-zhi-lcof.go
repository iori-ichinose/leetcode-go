/*
 * https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/
 * 2020.12.10
 * Golang 0ms(100.00%), 3.4MB(73.22%)
 */
package linkedList

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	pNode := head
	for pNode != nil {
		cloned := &Node{Val: pNode.Val, Next: pNode.Next, Random: nil}
		pNode.Next = cloned
		pNode = pNode.Next.Next
	}

	pNode = head
	for pNode != nil {
		if pNode.Random != nil {
			pNode.Next.Random = pNode.Random.Next
		}
		pNode = pNode.Next.Next
	}

	copyHead, curCopy := head.Next, head.Next
	pNode = head
	for pNode != nil {
		pNode.Next = pNode.Next.Next
		pNode = pNode.Next
		if curCopy.Next != nil {
			curCopy.Next = curCopy.Next.Next
			curCopy = curCopy.Next
		}
	}
	return copyHead
}

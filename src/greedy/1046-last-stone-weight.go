/*
 * https://leetcode-cn.com/problems/last-stone-weight/
 * 2020.12.10
 * Golang 0ms(100.00%), 2MB(22.97%)
 */
package greedy

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func lastStoneWeight(stones []int) int {
	var h IntHeap
	heap.Init(&h)
	for _, w := range stones {
		heap.Push(&h, w)
	}

	for len(h) > 1 {
		y := heap.Pop(&h)
		x := heap.Pop(&h)
		diff := y.(int) - x.(int)
		if diff != 0 {
			heap.Push(&h, diff)
		}
	}

	if len(h) == 0 {
		return 0
	} else {
		return heap.Pop(&h).(int)
	}
}

/*
 * https://leetcode-cn.com/problems/reveal-cards-in-increasing-order/
 * 2020.12.3
 * Golang 4ms(76.92%), 3.2MB(69.23%)
 */
package sorting

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	res := make([]int, n)
	sort.Ints(deck)
	for _, n := range deck {
		res[index[0]] = n
		index = index[1:]
		if len(index) != 0 {
			index = append(index, index[0])
			index = index[1:]
		}
	}
	return res
}

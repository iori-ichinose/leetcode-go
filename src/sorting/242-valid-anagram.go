/*
 * https://leetcode-cn.com/problems/valid-anagram/
 * 2020.11.22
 * 一般要求：排序
 * Golang 8ms(60.17%), 3MB(20.51%)
 * 进阶要求：Unicode, Map
 * Golang 8ms(60.17%), 2.7MB(75.85%)
 *
 * 这排序的效率怎么比我想象的低呢...
 */
package sorting

import "sort"

func isAnagram(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	return string(s1) == string(s2)
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s1 := map[rune]int{}
	for _, ch := range s {
		s1[ch]++
	}

	for _, ch := range t {
		s1[ch]--
		if s1[ch] < 0 {
			return false
		}
	}
	return true
}

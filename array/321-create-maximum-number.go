/*
 * https://leetcode-cn.com/problems/create-maximum-number/
 * 2020.12.2
 * 没做出来，太难了
 */
package array

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	maxSubsequence := func(a []int, k int) []int {
		var s []int
		for i, v := range a {
			for len(s) > 0 && len(s)+len(a)-i-1 >= k && v > s[len(s)-1] {
				s = s[:len(s)-1]
			}
			if len(s) < k {
				s = append(s, v)
			}
		}
		return s
	}

	lexicographicalLess := func(a, b []int) bool {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return a[i] < b[i]
			}
		}
		return len(a) < len(b)
	}

	merge := func(a, b []int) []int {
		merged := make([]int, len(a)+len(b))
		for i := range merged {
			if lexicographicalLess(a, b) {
				merged[i], b = b[0], b[1:]
			} else {
				merged[i], a = a[0], a[1:]
			}
		}
		return merged
	}

	start := 0
	var res []int
	if k > len(nums2) {
		start = k - len(nums2)
	}

	for i := start; i <= k && i <= len(nums1); i++ {
		s1, s2 := maxSubsequence(nums1, i), maxSubsequence(nums2, k-i)
		merged := merge(s1, s2)
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return res
}

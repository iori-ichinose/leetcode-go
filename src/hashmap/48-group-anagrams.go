/*
 * https://leetcode-cn.com/problems/group-anagrams/
 * 2020.12.14
 * Golang 24ms(90.63%), 8.1MB(46.37%)
 */
package hashmap

import "sort"

func groupAnagrams(strs []string) [][]string {
	counts := map[string][]string{}
	var res [][]string

	for _, str := range strs {
		arr := []byte(str)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		tmp := string(arr)
		counts[tmp] = append(counts[tmp], str)
	}

	for _, arr := range counts {
		res = append(res, arr)
	}
	return res
}

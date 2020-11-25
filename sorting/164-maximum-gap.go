/*
 * https://leetcode-cn.com/problems/maximum-gap/
 * 2020.11.26
 * Golang 4ms(91.67%), 3.3MB(67.74%)
 */
package sorting

func maximumGap(nums []int) int {
	n, res := len(nums), 0
	if n < 2 {
		return 0
	}

	type pair struct {
		min, max int
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		} else if num < minVal {
			minVal = num
		}
	}

	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	d := max(1, (maxVal-minVal)/(n-1))
	bucketSize := (maxVal-minVal)/d + 1
	bucket := make([]pair, bucketSize)
	for i := range bucket {
		bucket[i] = pair{-1, -1}
	}

	for _, num := range nums {
		index := (num - minVal) / d
		if bucket[index].min == -1 {
			bucket[index].min = num
			bucket[index].max = num
		} else {
			bucket[index].min = min(bucket[index].min, num)
			bucket[index].max = max(bucket[index].max, num)
		}
	}

	prev := -1
	for i, b := range bucket {
		if b.min == -1 {
			continue
		}
		if prev != -1 {
			res = max(res, b.min-bucket[prev].max)
		}
		prev = i
	}
	return res
}

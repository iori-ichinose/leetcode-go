package qstring

func maxPower(s string) int {
	max := func(x int, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	count, maxCount, n := 1, 1, len(s)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
			maxCount = max(maxCount, count)
		} else {
			count = 1
		}
	}

	return maxCount
}

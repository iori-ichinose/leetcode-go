package greedy

import "sort"

func leastInterval(tasks []byte, n int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	length := len(tasks)
	tasklist := make([]int, 26)
	for _, ch := range tasks {
		tasks[ch-'a']++
	}

	sort.Slice(tasklist, func(i, j int) bool { return i < j })
	count := 1
	for count < len(tasklist) && tasklist[count] == tasklist[0] {
		count++
	}
	return max(length, count+(n+1)*(tasklist[0]-1))
}

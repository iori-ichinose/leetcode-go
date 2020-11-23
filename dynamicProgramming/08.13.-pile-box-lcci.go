package dynamicProgramming

import "sort"

type boxSort [][]int

func (b boxSort) Less(x, y int) bool { return b[x][2] < b[y][2] }
func (b boxSort) Len() int           { return len(b) }
func (b boxSort) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func pileBox(box [][]int) int {
	sort.Sort(boxSort(box))
	var dp = make([]int, len(box))
	dp[0] = box[0][2]
	res := dp[0]
	for i := 1; i < len(box); i++ {
		maxHeight := 0
		for j := 0; j < i; j++ {
			if box[j][0] < box[i][0] &&
				box[j][1] < box[i][1] &&
				box[j][2] < box[i][2] && maxHeight < dp[j] {
				maxHeight = dp[j]
			}
		}
		dp[i] = maxHeight + box[i][2]
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

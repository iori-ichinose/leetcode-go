package dynamicProgramming

import "fmt"

func numRollsToTarget(d int, f int, target int) int {
	const MOD int = 1000000007
	dp := make([][]int, d+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	for i := 1; i < d+1; i++ {
		e := i * f
		if e > target {
			e = target
		}
		for j := i; j <= e; j++ {
			cnt := 0
			for k := 1; k <= f && j-k >= 0; k++ {
				cnt = (cnt + dp[i-1][j-k]) % MOD
			}
			dp[i][j] = cnt
		}

	}
	return dp[d][target]
}

func NumRollsToTargetTb() {
	fmt.Print("Testbench numRollsToTarget: ")
	fmt.Println(numRollsToTarget(1, 6, 3))
}

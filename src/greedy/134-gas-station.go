package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	n, i := len(gas), 0
	for i < n {
		sumGas, sumCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumGas += gas[j]
			sumCost += cost[j]
			if sumCost > sumGas {
				break
			} else {
				cnt++
			}
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}

	return -1
}

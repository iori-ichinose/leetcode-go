package dynamicProgramming

func maxProfit3(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	for i := range prices {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

package _063_股票的最大利润

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxProfit(prices []int) int {
	result, cost := 0, (1<<31)-1
	for _, price := range prices {
		cost = min(cost, price)
		result = max(result, price-cost)
	}

	return result
}

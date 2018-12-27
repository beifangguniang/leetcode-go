package Problem0122

func maxProfit(prices []int) int {
	count := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			count += prices[i] - prices[i-1]
		}
	}
	return count
}

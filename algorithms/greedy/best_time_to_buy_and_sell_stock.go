package greedy

// MaxProfit returns the maximum profit from buying once and selling once
// afterward. It returns 0 when no profitable transaction is possible.
// The input is not modified.
func MaxProfit(prices []int) int {
	maxProfit := 0
	buy := 0
	sale := 1
	for buy < len(prices)-1 && sale < len(prices) {
		if prices[buy] > prices[sale] {
			buy = sale
			sale++
			continue
		}
		profit := prices[sale] - prices[buy]
		maxProfit = max(profit, maxProfit)
		sale++
	}

	return maxProfit
}

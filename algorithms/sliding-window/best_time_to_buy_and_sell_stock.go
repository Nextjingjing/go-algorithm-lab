package slidingwindow

// MaxProfit returns the maximum profit from buying once and selling once
// afterward. It returns 0 when no profitable transaction is possible.
// The input is not modified.
func MaxProfit(prices []int) int {
	buy := 0
	sale := 1
	maxProfit := 0
	for sale < len(prices) {
		if prices[buy] > prices[sale] {
			buy = sale
		} else {
			profit := prices[sale] - prices[buy]
			maxProfit = max(profit, maxProfit)
		}
		sale++
	}
	return maxProfit
}

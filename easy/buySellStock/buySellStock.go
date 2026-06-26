package buysellstock

// Thinking:
// Because buying needs to happen before selling
// If we look at back first
// We will know what cheapest price a item might reach
// we can track it
// and soon as we see a profitable bid,
// we change our profits to that bid

func MaxProfit(prices []int) int {
	profit := 0

	var lbo int = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < lbo {
			lbo = prices[i]
		} else if prices[i]-lbo > profit {
			profit = prices[i] - lbo
		}
	}

	return profit
}

func MaxProfit2(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

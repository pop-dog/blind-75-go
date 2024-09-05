package main

func maxProfit(prices []int) int {
	best_profit := 0
	current_buy := -1
	for _, price := range prices {
		// Check if we should buy at this price,
		// i.e., this price is better than our current buy price
		if current_buy == -1 || price < current_buy {
			current_buy = price
			continue
		}
		// Check if this is our new best sell price
		if current_profit := price - current_buy; current_profit > best_profit {
			best_profit = current_profit
		}
	}
	return best_profit
}

func main() {
	x := maxProfit([]int{7, 1, 5, 3, 6, 4})
	println("x:", x)
}

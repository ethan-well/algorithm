package main

func main() {
	prices := []int{7,6,4,3,1}

	println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]

	var maxProfit int

	for i := 0; i < len(prices); i ++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
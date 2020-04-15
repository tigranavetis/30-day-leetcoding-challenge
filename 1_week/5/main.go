package main

import (
	"fmt"
)

// maxProfit finds the best time to buy and sell stock.
func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

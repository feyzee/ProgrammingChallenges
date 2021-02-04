package main

import "fmt"

func maximumProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	if len(prices) <= 1 {
		return maxProfit
	}

	for _, value := range prices {
		if value < minPrice {
			minPrice = value
		} else if (value - minPrice) > maxProfit {
			maxProfit = value - minPrice
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 7, 3, 6, 4}

	x := maximumProfit(prices)
	fmt.Println(x)
}

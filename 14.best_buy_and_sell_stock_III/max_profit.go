package main

import (
	"fmt"
)

func maxProfitLeft(prices []int) []int {
	profits := make([]int, len(prices))
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			profits[i] = profits[i-1]
		} else {
			if prices[i]-buy > profits[i-1] {
				profits[i] = prices[i] - buy
			} else {
				profits[i] = profits[i-1]
			}

		}
	}
	return profits
}
func maxProfitRight(prices []int) []int {
	profits := make([]int, len(prices))
	sell := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > sell {
			sell = prices[i]
			profits[i] = profits[i+1]
		} else {
			if sell-prices[i] > profits[i+1] {
				profits[i] = sell - prices[i]
			} else {
				profits[i] = profits[i+1]
			}
		}
	}
	return profits
}

func maxProfit(prices []int) int {
	left := maxProfitLeft(prices)
	right := maxProfitRight(prices)
	maxResult := 0
	for i := 0; i < len(prices); i++ {
		if left[i]+right[i] > maxResult {
			maxResult = left[i] + right[i]
		}
	}
	return maxResult
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	// prices := []int{1, 2, 3, 4, 5}
	// prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	// memoProfit := make([]int, len(prices))
	// fmt.Println(memoProfit)
}

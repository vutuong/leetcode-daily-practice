package main

import (
	"fmt"
	"math/bits"
)

// 21. Best Time to Buy and Sell Stock Easy
// Add to List

// Share
// Say you have an array for which the ith element is the price of a given stock on day i.

// If you were only permitted to complete at most one transaction
// (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

// Note that you cannot sell a stock before you buy one.

// Example 1:

// Input: [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// 			 Not 7-1 = 6, as selling price needs to be larger than buying price.

// Example 2:

// Input: [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.

// # Solution 1
// func maxProfit(prices []int) int {
// 	maxResult := 0
// 	for i := 0; i <= len(prices)-2; i++ {
// 		sortTheRest := make([]int, len(prices[i+1:]))
// 		copy(sortTheRest, prices[i+1:])
// 		sort.Ints(sortTheRest)
// 		if sortTheRest[len(sortTheRest)-1] < prices[i] {
// 			continue
// 		}
// 		if maxResult < sortTheRest[len(sortTheRest)-1]-prices[i] {
// 			maxResult = sortTheRest[len(sortTheRest)-1] - prices[i]
// 		}
// 	}
// 	return maxResult
// }

// # Solution 2 : One Pass
func maxProfit(prices []int) int {
	const maxInt = 1<<(bits.UintSize-1) - 1
	minPrice := maxInt
	maxResult := 0
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxResult {
			maxResult = prices[i] - minPrice
		}
	}
	return maxResult
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	// prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

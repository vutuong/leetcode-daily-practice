package main

import "testing"

func TestSum(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	test1 := maxProfit(prices)
	if test1 != 7 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test1, 5)
	}

	prices = []int{1, 2, 3, 4, 5}
	test2 := maxProfit(prices)
	if test2 != 4 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test2, 0)
	}

	prices = []int{7, 6, 4, 3, 1}
	test3 := maxProfit(prices)
	if test3 != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test3, 0)
	}
}

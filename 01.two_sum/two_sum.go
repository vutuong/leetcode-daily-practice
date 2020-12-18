package main

import "fmt"

// Problem: Two Sum
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution
// and you may not use the same element twice.

// You can return the answer in any order.
func twoSum(nums []int, target int) []int {
	results := []int{}
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				results = []int{i, j}
			}
		}
	}
	return results
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 10
	fmt.Println(twoSum(nums, target))
}

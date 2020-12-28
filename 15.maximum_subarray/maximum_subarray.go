package main

import "fmt"

// 53. Maximum Subarray
// Easy

// Share
// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
// Follow up: If you have figured out the O(n) solution,
// try coding another solution using the divide and conquer approach, which is more subtle.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Example 2:

// Input: nums = [1]
// Output: 1
// Example 3:

// Input: nums = [0]
// Output: 0
// Example 4:

func maxSubArray(nums []int) int {
	maxResult := nums[0]
	maxTemporary := 0
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		maxTemporary = maxTemporary + nums[i]
		if maxResult < maxTemporary {
			maxResult = maxTemporary
		}
		if maxTemporary < 0 {
			maxTemporary = 0
		}
	}
	return maxResult
}
func main() {
	// input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	input := []int{8, -19, 5, -4, 20}
	// maxSubArray(input)
	test1 := maxSubArray(input)
	fmt.Println(test1)
}

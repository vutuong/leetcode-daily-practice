package main

import "fmt"

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

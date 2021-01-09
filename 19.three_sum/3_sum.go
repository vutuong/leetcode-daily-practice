package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Initial solution: Brute-force O( n³ ) nested loop scan
func compareSlice(s1, s2 []int) bool {
	sort.Ints(s1)
	sort.Ints(s2)
	return reflect.DeepEqual(s1, s2)
}
func uniqueResult(result [][]int, input []int) [][]int {
	for i := range result {
		if compareSlice(input, result[i]) == true {
			return result
		}
	}
	result = append(result, input)
	return result
}

func threeSumBruteforce(nums []int) [][]int {
	result := [][]int{}
	for i := range nums {
		for left := i + 1; left < len(nums); left++ {
			for right := left + 1; right < len(nums); right++ {
				if nums[i]+nums[left]+nums[right] == 0 {
					result = uniqueResult(result, []int{nums[i], nums[left], nums[right]})
				}
			}
		}
	}
	return result
}

// Accepted solution : left-right scan with pre-sorting
func threeSum(nums []int) [][]int {
	// sort input slice
	sort.Ints(nums)
	// create 2-d slice to store answer
	result := [][]int{}
	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		// srightip the repetition of the same element
		if i > 0 && (nums[i-1] == nums[i]) {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	// input1 := []int{-4, -1, 0, 1, 2, -1}
	result := threeSum(input)
	fmt.Println(result)
	// fmt.Println(compareSlice(input, input1))
}

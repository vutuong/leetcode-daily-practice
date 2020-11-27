package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Problem: Palindrome Number
// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
// Follow up: Could you solve it without converting the integer to a string?

// Example 1:
// Input: x = 121
// Output: true

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Example 4:
// Input: x = -101
// Output: false

func isPalindrome(x int) bool {
	if x < 0 || x == 10 {
		return false
	}
	if x < 10 {
		return true
	}

	// Solution 1: Use the slice and loop over it
	// digits := []int{}
	// for i := 0; i <= len(digits)/2; i++ {
	// 	if digits[i] != digits[len(digits)-i-1] {
	// 		return false
	// 	}
	// }

	//Solution 2: Use log to find a len of the input number
	// Find last half part of number => reverse => compare with first half part
	len := int(math.Floor(math.Log10(float64(x)) + 1))
	var y float64 = 0
	fmt.Println(len)
	for i := 0; i < int(len/2); i++ {
		y = y + float64(x%10)*math.Pow10(len/2-i-1)
		x /= 10
	}
	fmt.Println(float64(y))
	fmt.Println(x)
	if len%2 == 0 {
		if y != float64(x) {
			return false
		}
	} else {
		if y != float64(x/10) {
			return false
		}
	}
	return true
}
func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < 50; i++ {
		num := r1.Intn(100000)
		fmt.Println(num)
		fmt.Println(isPalindrome(num))

	}
}

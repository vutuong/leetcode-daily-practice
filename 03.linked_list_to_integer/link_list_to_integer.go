package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	vals := []int{}
	result := 0
	for {
		vals = append(vals, head.Val)
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
	for i, val := range vals {
		result = result + val*int(math.Pow(2, float64(len(vals)-i-1)))
	}
	return result
}

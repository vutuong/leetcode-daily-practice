package main

import "fmt"

/* Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursion solution
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	if l1 == nil && l2 != nil {
		result = l2
	} else if l1 != nil && l2 == nil {
		result = l1
	} else if l1 == nil && l2 == nil {
		result = nil
	} else {
		if l1.Val < l2.Val {
			result.Val = l1.Val
			result.Next = mergeTwoLists(l1.Next, l2)
		} else {
			result.Val = l2.Val
			result.Next = mergeTwoLists(l1, l2.Next)
		}
	}
	return result
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	// l1 := &ListNode{}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(l1, l2)
	fmt.Println(result)
}

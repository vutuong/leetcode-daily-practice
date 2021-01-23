package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/
83. Remove Duplicates from Sorted List - Easy

Given the head of a sorted linked list, delete all duplicates such that
each element appears only once. Return the linked list sorted as well.

Example 1:
Input: head = [1,1,2]
Output: [1,2]

Example 2:
Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for {
		if cur == nil || cur.Next == nil {
			break
		}
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		},
	}
	result := deleteDuplicates(l1)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

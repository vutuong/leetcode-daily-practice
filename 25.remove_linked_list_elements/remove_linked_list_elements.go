package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode.com/problems/remove-linked-list-elements/
203. Remove Linked List Elements
Remove all elements from a linked list of integers that have value val.

Example:
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
*/

func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{Val: 0, Next: head}
	cur := head

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head.Next
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
	result := removeElements(l1, 4)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}

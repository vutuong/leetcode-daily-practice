package main

import (
	"fmt"
	"reflect"
)

/*
https://leetcode.com/problems/palindrome-linked-list/
234. Palindrome Linked List
Given a singly linked list, determine if it is a palindrome.

Example 1:
Input: 1->2
Output: false

Example 2:
Input: 1->2->2->1
Output: true
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution1: Revere linked list and compare with the original
func isPalindrome(head *ListNode) bool {
	rev := reverseList(head)
	if ok := compareResult(rev, head); ok == true {
		return true
	}
	return false
}
func reverseList(head *ListNode) *ListNode {
	rev := new(ListNode)
	if head == nil {
		return head
	}
	rev.Val = head.Val
	rev.Next = nil
	cur := head.Next
	for cur != nil {
		rev = addHead(rev, cur)
		cur = cur.Next
	}
	return rev
}
func addHead(head, cur *ListNode) *ListNode {
	temp := new(ListNode)
	temp.Val = cur.Val
	temp.Next = head
	head = temp
	return head
}

func compareResult(r1, r2 *ListNode) bool {
	ok := true
	if r1 == nil && r2 == nil {
		return ok
	} else if r1 == nil && r2 != nil {
		return false
	} else if r1 != nil && r2 == nil {
		return false
	} else if r1.Val != r2.Val {
		return false
	} else {
		ok = compareResult(r1.Next, r2.Next)
	}
	return ok
}

// Solution2: Use slice to store all value of node, reverse and compare
func isPalindrome_Slice(head *ListNode) bool {
	storeVal := []int{}
	reverseStoreVal := []int{}
	cur := head
	for cur != nil {
		storeVal = append(storeVal, cur.Val)
		reverseStoreVal = append(reverseStoreVal, cur.Val)
		cur = cur.Next
	}
	reverseStoreVal = reverseSlice(reverseStoreVal)
	return reflect.DeepEqual(reverseStoreVal, storeVal)
}
func reverseSlice(sl []int) []int {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return sl
}

func main() {
	l1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	result := isPalindrome(l1)
	fmt.Println(result)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//Problem: You are given two non-empty linked lists representing two non-negative integers.
//The digits are stored in reverse order, and each of their nodes contains a single digit.
//Add the two numbers and return the sum as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val+l2.Val >= 10 {
		val := l1.Val + l2.Val - 10
		if l1.Next != nil && l2.Next != nil {
			l2.Next.Val = l2.Next.Val + 1
		}
		if l1.Next == nil && l2.Next == nil {
			l1.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
			l2.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
		}
		if l1.Next == nil && l2.Next != nil {
			l1.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
		if l1.Next != nil && l2.Next == nil {
			l2.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
		return &ListNode{val, addTwoNumbers(l1.Next, l2.Next)}
	}
	val := l1.Val + l2.Val
	if l1.Next == nil && l2.Next == nil {
		return &ListNode{val, nil}
	}
	if l1.Next == nil && l2.Next != nil {
		l1.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	if l1.Next != nil && l2.Next == nil {
		l2.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	return &ListNode{val, addTwoNumbers(l1.Next, l2.Next)}
}

func main() {
	list1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  8,
			Next: nil,
		},
	}
	list2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	result := addTwoNumbers(list1, list2)
	fmt.Println(result.Val)
}

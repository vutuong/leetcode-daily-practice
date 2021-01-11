package main

import (
	"errors"
	"fmt"
)

/* Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Remove Nth Node From End of List
// Given the head of a linked list, remove the nth node from the end of the list and return its head.
// Follow up: Could you do this in one pass?
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
	len  int
}

func (l *LinkedList) getSize() int {
	ptr := l.head
	count := 0
	for {
		if ptr.Next != nil {
			count++
			ptr = ptr.Next
		} else {
			break
		}
	}
	count++
	return count
}

// get node at the specfic position
func (l *LinkedList) getAtPos(pos int) (*ListNode, error) {
	ptr := l.head
	// if pos < 0 return head of LinkedList
	if pos < 0 {
		fmt.Println("Position can not be negative")
		return nil, errors.New("Position can not be negative")
	}
	if pos > l.len-1 {
		fmt.Println("Position is bigger than length")
		return nil, errors.New("Position is bigger than length")
	}

	for i := 0; i < pos; i++ {
		ptr = ptr.Next
	}
	return ptr, nil
}

// del a node at a specific position
func (l *LinkedList) deleteAtPos(pos int) error {
	if l.len == 0 {
		fmt.Println("Linkedlist is empty")
		return errors.New("Linkedlist is empty")
	}
	if pos < 0 || pos > l.len {
		fmt.Println("Pos can'nt be negative or bigger than length")
		return errors.New("Pos can'nt be negative or bigger than length")
	}
	if pos == 0 {
		l.head, _ = l.getAtPos(pos + 1)
		l.len--
		return nil
	}

	ptrPrev, _ := l.getAtPos(pos - 1)
	ptrNex, _ := l.getAtPos(pos + 1)

	ptrPrev.Next = ptrNex
	l.len--
	return nil
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := LinkedList{head, 0}
	l.len = l.getSize()
	l.deleteAtPos(l.len - n)
	return l.head
}

func main() {
	fmt.Println("Remove Nth Node From End of List")
}

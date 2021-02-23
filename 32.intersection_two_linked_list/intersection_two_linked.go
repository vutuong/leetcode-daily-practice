package getIntersectionNode

/*
Intersection of Two Linked Lists
https://leetcode.com/problems/intersection-of-two-linked-lists/

Write a program to find the node at which the intersection of two singly linked lists begins.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, ndA := headA, headA
	nodeB, ndB := headB, headB
	lenA, lenB := 0, 0
	for nodeA != nil {
		lenA += 1
		nodeA = nodeA.Next
	}
	for nodeB != nil {
		lenB += 1
		nodeB = nodeB.Next
	}
	if lenA > lenB {
		for lenA > lenB {
			ndA = ndA.Next
			lenA -= 1
		}
	} else if lenB > lenA {
		for lenB > lenA {
			ndB = ndB.Next
			lenB -= 1
		}
	}
	for ndA != nil && ndB != nil {
		if ndA == ndB {
			return ndA
		}
		ndA = ndA.Next
		ndB = ndB.Next
	}
	return nil
}

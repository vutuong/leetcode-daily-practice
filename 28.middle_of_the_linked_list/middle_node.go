package middleNode

/* Middle of the Linked List
https://leetcode.com/problems/middle-of-the-linked-list/

Given a non-empty, singly linked list with head node head, return a middle node of linked list.
If there are two middle nodes, return the second middle node.

Example 1:
Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.

Example 2:
Input: [1,2,3,4,5,6]
Output: Node 4 from this list (Serialization: [4,5,6])
Since the list has two middle nodes with values 3 and 4, we return the second one.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Using 2 pointer
// When two step pointer reach to the end, one step pointer reach to the middle
func middleNode(head *ListNode) *ListNode {
	smallStep := head
	bigStep := head
	for {
		if bigStep.Next == nil {
			return smallStep
		}
		if bigStep.Next.Next == nil {
			return smallStep.Next
		}
		bigStep = bigStep.Next.Next
		smallStep = smallStep.Next
	}
}

// Solution 2: Walk through all node in the linked list, and store every node to a slice
// Get the middle element of the output slice is easier.
func middleNodeSlice(head *ListNode) *ListNode {
	s := []*ListNode{}
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	return s[len(s)/2]
}

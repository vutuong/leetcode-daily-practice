package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode.com/problems/reverse-linked-list/
206. Reverse Linked List

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/
// Iterative solution
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

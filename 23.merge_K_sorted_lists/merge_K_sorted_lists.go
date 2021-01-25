package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode.com/problems/merge-k-sorted-lists/
23. Merge k Sorted Lists - Hard

	You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

	Merge all the linked-lists into one sorted linked-list and return it.
*/
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
func mergeKLists(lists []*ListNode) *ListNode {
	var start *ListNode
	if len(lists) == 0 {
		return nil
	} else {
		start = lists[0]
	}
	for i := 1; i < len(lists); i++ {
		start = mergeTwoLists(start, lists[i])
	}
	return start
}

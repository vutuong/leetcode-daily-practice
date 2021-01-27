package linked_list_cycle

/* 141. Linked List Cycle
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

https://leetcode.com/problems/linked-list-cycle/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution 1: Using hashmap to check if the looped over pointer is contain in hashmap or not
// If yes => return cycle linked list
// If not => add the pointer as a new key to the hash map
func hasCycle_Solution1(head *ListNode) bool {
	cur := head
	store := make(map[*ListNode]bool)
	for cur != nil {
		if _, ok := store[cur]; ok {
			return true
		}
		store[cur] = true
		cur = cur.Next
	}
	return false
}

//Solution 2: Using two pointer with different step to loop over the linked list
// If two pointer meet each other at any point, return True
func hasCycle(head *ListNode) bool {
	oneStep, twoStep := head, head
	for head != nil {
		if twoStep.Next == nil {
			break
		} else if twoStep.Next.Next == nil {
			break
		}
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
		if oneStep == twoStep {
			return true
		}
	}
	return false
}

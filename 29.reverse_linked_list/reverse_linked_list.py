# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
'''
https://leetcode.com/problems/reverse-linked-list/
206. Reverse Linked List
Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
'''
# Recursive solution: add next node to the head of None linkedlist
class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        cur = head
        rev = None
        while cur !=None:
            rev = self.addHead(rev,cur)
            cur = cur.next
        return rev
    def addHead(self, rev: ListNode, cur: ListNode) -> ListNode:
        temp = ListNode()
        temp.val = cur.val
        temp.next = rev
        rev = temp
        return rev
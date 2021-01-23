package main

import (
	"testing"
)

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
func TestSum(t *testing.T) {
	var l1, expected, result *ListNode
	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	expected = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	result = deleteDuplicates(l1)
	if !compareResult(result, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}

	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	expected = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	result = deleteDuplicates(l1)
	if !compareResult(result, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}
}

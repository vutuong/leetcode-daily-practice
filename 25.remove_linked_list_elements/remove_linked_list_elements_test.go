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

func testSum(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	val1 := 6
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	result := removeElements(l1, val1)
	if !compareResult(result, expected) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}
}

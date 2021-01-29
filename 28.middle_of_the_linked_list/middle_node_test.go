package middleNode

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
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
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
	}
	result := middleNode(l1)
	if result.Val != 3 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 3)
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
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
	}
	result = middleNode(l2)
	if result.Val != 4 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 4)
	}

	l3 := &ListNode{
		Val:  1,
		Next: nil,
	}
	result = middleNode(l3)
	if result.Val != 1 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 1)
	}

	l4 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	result = middleNode(l4)
	if result.Val != 2 {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, 2)
	}
}

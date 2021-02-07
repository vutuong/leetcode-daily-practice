package main

import "testing"

func TestSum(t *testing.T) {
	var test1, test2, test3, expected bool
	l1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	test1 = isPalindrome(l1)
	expected = false
	if test1 != expected {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test1, expected)
	}

	l2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	test2 = isPalindrome(l2)
	expected = true
	if test2 != expected {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test2, expected)
	}

	l3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	test3 = isPalindrome(l3)
	expected = true
	if test3 != expected {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test3, expected)
	}
}

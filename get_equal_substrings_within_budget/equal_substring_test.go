package main

import "testing"

func TestSum(t *testing.T) {
	s := "abcd"
	g := "bcdf"
	maxCost := 3
	test1 := equalSubstring(s, g, maxCost)
	if test1 != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", test1, 3)
	}

	s = "abcd"
	g = "cdef"
	maxCost = 3
	test2 := equalSubstring(s, g, maxCost)
	if test1 != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", test2, 3)
	}

	s = "abcd"
	g = "acde"
	maxCost = 0
	test3 := equalSubstring(s, g, maxCost)
	if test1 != 3 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", test3, 3)
	}

}

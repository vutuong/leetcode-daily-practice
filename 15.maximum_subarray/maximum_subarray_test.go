package main

import "testing"

func TestSum(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	test1 := maxSubArray(input)
	if test1 != 6 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test1, 6)
	}

	input = []int{1}
	test2 := maxSubArray(input)
	if test2 != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test2, 1)
	}

	input = []int{-2, 7, 6, 4, 3, -11}
	test3 := maxSubArray(input)
	if test3 != 20 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test3, 20)
	}

	input = []int{-2, 1}
	test4 := maxSubArray(input)
	if test4 != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test4, 1)
	}

	input = []int{-2, -1}
	test5 := maxSubArray(input)
	if test5 != -1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test5, -1)
	}

	input = []int{8, -19, 5, -4, 20}
	test6 := maxSubArray(input)
	if test6 != 21 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", test6, 21)
	}

}

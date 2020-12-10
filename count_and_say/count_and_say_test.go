package main

import "testing"

func TestSum(t *testing.T) {
	var n int

	n = 1
	test1 := countAndSay(n)
	if test1 != "1" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test1, "1")
	}

	n = 4
	test2 := countAndSay(n)
	if test2 != "1211" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test2, "1211")
	}

	// s = "5"
	// numRows = 1
	// test3 := convert(s, numRows)
	// if test3 != "A" {
	// 	t.Errorf("Result was incorrect, got: %s, want: %s.", test3, "A")
	// }
}

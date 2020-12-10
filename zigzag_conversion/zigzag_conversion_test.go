package main

import "testing"

func TestSum(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	test1 := convert(s, numRows)
	if test1 != "PAHNAPLSIIGYIR" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test1, "PAHNAPLSIIGYIR")
	}

	s = "PAYPALISHIRING"
	numRows = 4
	test2 := convert(s, numRows)
	if test2 != "PINALSIGYAHRPI" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test2, "PAYPALISHIRING")
	}

	s = "A"
	numRows = 1
	test3 := convert(s, numRows)
	if test3 != "A" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test3, "A")
	}
}

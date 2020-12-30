package main

import "testing"

func TestSum(t *testing.T) {
	input := "aa"
	p := "a"
	test1 := isMatch(input, p)
	if test1 != false {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test1, false)
	}

	input = "aa"
	p = "*"
	test2 := isMatch(input, p)
	if test2 != true {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test2, true)
	}

	input = "cb"
	p = "?a"
	test3 := isMatch(input, p)
	if test3 != false {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test3, false)
	}

	input = "adceb"
	p = "*a*b"
	test4 := isMatch(input, p)
	if test4 != true {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test4, true)
	}

	input = "acdcb"
	p = "a*c?b"
	test5 := isMatch(input, p)
	if test5 != false {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test5, false)
	}

	input = "a"
	p = "aa"
	test6 := isMatch(input, p)
	if test6 != false {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test6, false)
	}

	input = ""
	p = "*****"
	test7 := isMatch(input, p)
	if test7 != true {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test7, true)
	}

	input = "abcabczzzde"
	p = "*abc???de*"
	test8 := isMatch(input, p)
	if test8 != true {
		t.Errorf("Result was incorrect, got: %t, want: %t.", test8, true)
	}

}

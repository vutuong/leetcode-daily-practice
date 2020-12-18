package main

import (
	"reflect"
	"testing"
)

// func Equal(a, b []string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i, v := range a {
// 		if v != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
func TestSum(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	test1 := fullJustify(words, maxWidth)
	expect := []string{"This    is    an", "example  of text", "justification.  "}
	if !reflect.DeepEqual(test1, expect) {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test1, "x")
	}

	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth = 16
	test2 := fullJustify(words, maxWidth)
	expect = []string{"What   must   be", "acknowledgment  ", "shall be        "}
	if !reflect.DeepEqual(test2, expect) {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test2, "x")
	}

	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth = 20
	test3 := fullJustify(words, maxWidth)
	expect = []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}
	if !reflect.DeepEqual(test3, expect) {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test3, "x")
	}

	words = []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth = 16
	test4 := fullJustify(words, maxWidth)
	expect = []string{"This    is    an", "example  of text", "justification.  "}
	if !reflect.DeepEqual(test4, expect) {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test4, "x")
	}

	words = []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	maxWidth = 16
	test5 := fullJustify(words, maxWidth)
	expect = []string{"ask   not   what", "your country can", "do  for  you ask", "what  you can do", "for your country"}
	if !reflect.DeepEqual(test5, expect) {
		t.Errorf("Result was incorrect, got: %s, want: %s.", test5, "x")
	}
}

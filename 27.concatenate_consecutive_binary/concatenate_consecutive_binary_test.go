package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	var input, expected, result int
	input = 3
	expected = 27
	result = concatenatedBinary(input)
	if result != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}

	input = 1
	expected = 1
	result = concatenatedBinary(input)
	if result != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}

	input = 12
	expected = 505379714
	result = concatenatedBinary(input)
	if result != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}
}

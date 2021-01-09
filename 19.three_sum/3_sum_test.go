package main

import (
	"fmt"
	"testing"
)

// func compareSlice(s1, s2 []int) bool {
// 	sort.Ints(s1)
// 	sort.Ints(s2)
// 	return reflect.DeepEqual(s1, s2)
// }
func checkEqualResult(result [][]int, expected [][]int) bool {
	var ok bool
	for i := range result {
		ok = false
		for j := range expected {
			if compareSlice(expected[j], result[i]) == true {
				ok = true
				break
			}
		}
		if ok == false {
			return false
		}
	}
	return true
}
func TestSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	test1 := threeSum(input)
	expected := make([][]int, 2)
	expected[0] = []int{-1, -1, 2}
	expected[1] = []int{-1, 0, 1}
	fmt.Println(expected)
	if checkEqualResult(test1, expected) != true {
		t.Errorf("Result was incorrect, got: %v, want: %v.", test1, false)
	}
}

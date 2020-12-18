package main

import (
	"fmt"
)

//Longest Common Prefix

//Write a function to find the longest common prefix string amongst an array of strings.
//If there is no common prefix, return an empty string "".

//Example 1:
//Input: strs = ["flower","flow","flight"]
//Output: "fl"

//Example 2:
//Input: strs = ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.

func twoStringLongest(str1, str2 string) string {
	if len(str1) > len(str2) {
		str1, str2 = str2, str1
	}
	longest := ""
	for i := 0; i <= len(str1); i++ {
		if str1[:len(str1)-i] == str2[:len(str1)-i] {
			longest = str1[:len(str1)-i]
			break
		}
	}
	return longest
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	compare := strs[0]
	for i := 1; i < len(strs); i++ {
		compare = twoStringLongest(compare, strs[i])
		if len(compare) == 0 {
			return ""
		}
	}
	return compare
}

func main() {
	input := []string{"", "flcow", "flcight"}
	// input := []string{}
	fmt.Println(longestCommonPrefix(input))
}

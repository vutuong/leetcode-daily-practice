package main

import (
	"fmt"
)

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

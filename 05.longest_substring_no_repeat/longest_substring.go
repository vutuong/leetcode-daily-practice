//Given a string s, find the length of the longest substring without repeating characters.
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLongs := 0
	if len(s) == 1 {
		return 1
	}
	for i := range s {
		include := false
		for j := i + 1; j < len(s); j++ {
			for z := i; z < j; z++ {
				if s[j] == s[z] {
					include = true
					break
				}
			}

			if include == true {
				if maxLongs < j-i {
					maxLongs = j - i
				}
				break
			}
			if j == len(s)-1 {
				if maxLongs < j+1-i {
					maxLongs = j + 1 - i
				}
			}
		}
	}
	return maxLongs
}

func main() {
	s := "tuong"
	fmt.Println(lengthOfLongestSubstring(s))
}

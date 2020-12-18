package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 38. Count and Say
// The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
// countAndSay(1) = "1"
// countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1),
// //  which is then converted into a different digit string.
// Input: n = 4
// Output: "1211"
// Explanation:
// countAndSay(1) = "1"
// countAndSay(2) = say "1" = one 1 = "11"
// countAndSay(3) = say "11" = two 1's = "21"
// countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
func convert(s string) string {
	count := 0
	strs := []string{}

	for i := range s {
		if i == 0 || s[i] == s[i-1] {
			count++
		}
		if i == len(s)-1 || s[i] != s[i+1] {
			strs = append(strs, strconv.Itoa(count), string(s[i]))
			count = 1
		}
	}
	fmt.Println(strs)
	return strings.Join(strs, "")

}

func countAndSay(n int) string {
	var input string
	for i := 0; i < n; i++ {
		if i == 0 {
			input = "1"
		} else {
			input = convert(input)
		}
	}
	return input
}

func main() {
	countAndSay(n)
}

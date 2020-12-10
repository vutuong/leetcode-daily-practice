package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	convert("12333344")
}

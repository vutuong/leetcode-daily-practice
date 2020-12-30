package main

import (
	"fmt"
	"strings"
)

// input = "adceb"
// p = "*a*b"

func isMatch(s string, p string) bool {
	str := strings.Split(s, "")
	pattern := strings.Split(p, "")

	/* Replace multiple * with one *
	Loop over item in pattern slice:
		- If pattern[i] is "*":
			+ If pattern[i] is First "*":
				++ Keep it in pattern: pattern[writeIndex] = pattern[i]
				++ Increase writeIndex
				++ Next "*" will not be first "*" : isFirst = false (Until other character appear)
		- If pattern[i] is not "*":
			+ Keep it in pattern: pattern[writeIndex] = pattern[i]
			+ Increase writeIndex
			+ Reset isFirst = true for the next closed "*"
	*/
	writeIndex := 0
	isFirst := true
	for i := range pattern {
		if pattern[i] == "*" {
			if isFirst {
				pattern[writeIndex] = pattern[i]
				writeIndex++
				isFirst = false
			}
		} else {
			pattern[writeIndex] = pattern[i]
			writeIndex++
			isFirst = true
		}
	}

	// create two dimension slice for dynamic programing
	dP := make([][]bool, len(str)+1)
	for i := range dP {
		dP[i] = make([]bool, writeIndex+1)
	}
	if writeIndex > 0 && pattern[0] == "*" {
		dP[0][1] = true
	}
	// To hard to find a solution by myself :)
	// This solution is ref to
	// https://www.youtube.com/watch?v=3ZDZ-N0EPV0&ab_channel=TusharRoy-CodingMadeSimple&fbclid=IwAR3Y4MV0U7fJVPWpcUKs-FuxDbgW3u4DbuLN8qmiIPHcd3FQ-Ixkwxg_XXY
	dP[0][0] = true
	// for j := 1; j <= len(p); j++ {
	// 	if p[j-1] == '*' {
	// 		dP[0][j] = dP[0][j-1]
	// 	}
	// }

	for i := 1; i < len(dP); i++ {
		for j := 1; j < len(dP[0]); j++ {
			// if p[j-1] == '*' {
			// 	dP[i][j] = dP[i][j-1] || dP[i-1][j]
			// } else if p[j-1] == '?' || p[j-1] == s[i-1] {
			// 	dP[i][j] = dP[i-1][j-1]
			// }
			if pattern[j-1] == "?" || str[i-1] == pattern[j-1] {
				dP[i][j] = dP[i-1][j-1]
			} else if pattern[j-1] == "*" {
				dP[i][j] = dP[i-1][j] || dP[i][j-1]
			}
		}
	}

	return dP[len(str)][writeIndex]
}

func main() {
	// input := "adceb"
	// p := "*a*b"
	// input := "aa"
	// p := "aa"

	// input := "ab"
	// p := "?*"
	// input := "a"
	// p := "aa"
	input := "abcabczzzde"
	p := "*abc???de*"

	result := isMatch(input, p)
	fmt.Println(result)
}

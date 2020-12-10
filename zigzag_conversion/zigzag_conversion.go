package main

import (
	"fmt"
	"strings"
)

// 6.  Conversion
// Medium
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
//  (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	column := 0
	row := 0

	rows := make([]string, numRows)
	for i := range s {
		if i%(numRows-1) == 0 {
			if (i/(numRows-1))%2 == 0 {
				row = 0
				if i != 0 {
					column = column + 1
				}
			} else {
				row = numRows - 1
			}
		} else {
			if (i/(numRows-1))%2 == 0 {
				row = row + 1
			} else {
				row = row - 1
				column = column + 1
			}
		}
		rows[row] = rows[row] + string(s[i])
	}

	return strings.Join(rows, "")
}

func main() {
	fmt.Println(convert("P", 1))
}

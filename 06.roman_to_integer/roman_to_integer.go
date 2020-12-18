package main

import (
	"fmt"
)

//Given a roman numeral, convert it to an integer.
// Roman numerals are usually written largest to smallest from left to right.
// However, the numeral for four is not IIII.
// Instead, the number four is written as IV.
// Because the one is before the five we subtract it making four.
// The same principle applies to the number nine, which is written as IX.
// There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
func specialConvert(s string) int {
	switch s {
	case "IV":
		return 4
	case "IX":
		return 9
	case "XL":
		return 40
	case "XC":
		return 90
	case "CD":
		return 400
	case "CM":
		return 900
	default:
		return 0
	}
}
func normalConvert(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}
func romanToInt(s string) int {
	toInt := 0
	if len(s) == 1 || len(s) == 0 {
		return normalConvert(s)
	}
	for i := range s {
		fmt.Println(string(s[i]))
		if i > 0 && specialConvert(string(s[i-1:i+1])) != 0 {
			continue
		}
		if i != len(s)-1 && i != len(s)-2 {
			if specialConvert(string(s[i:i+2])) == 0 {
				toInt = toInt + normalConvert(string(s[i]))
			} else {
				toInt = toInt + specialConvert(string(s[i:i+2]))
			}
		}
		if i == len(s)-1 || i == len(s)-2 {
			if specialConvert(string(s[i:])) == 0 {
				toInt = toInt + normalConvert(string(s[i]))
			} else {
				toInt = toInt + specialConvert(string(s[i:]))
			}

		}
	}
	return toInt
}

func main() {
	s := "LVIII"
	fmt.Println(romanToInt(s))
}

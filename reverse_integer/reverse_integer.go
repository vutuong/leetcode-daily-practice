package main

import (
	"math"
)

//Assume we are dealing with an environment that could only store integers
// within the 32-bit signed integer range: [−231,  231 − 1].
//For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
func reverse(x int) int {
	var y, z int = 0, 0
	if (x >= 0 && x < 10) || (x < 0 && x > -10) {
		return x
	}
	if x < 0 {
		z = -1 * x
	} else {
		z = x
	}
	len := int(math.Floor(math.Log10(float64(z)) + 1))
	for i := 0; i < int(len); i++ {
		y = y + z%10*int(math.Pow10(len-i-1))
		z /= 10
	}
	if x < 0 {
		y = -1 * y
	}
	if math.MaxInt32 < y || math.MinInt32 > y {
		return 0
	}
	return y
}

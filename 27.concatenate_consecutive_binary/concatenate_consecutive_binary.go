package main

/*
Concatenation of Consecutive Binary Numbers
Given an integer n, return the decimal value of the binary string formed by concatenating
the binary representations of 1 to n in order, modulo 109 + 7.

Example 1:
Input: n = 1
Output: 1
Explanation: "1" in binary corresponds to the decimal value 1.

Example 2:
Input: n = 3
Output: 27
Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
After concatenating them, we have "11011", which corresponds to the decimal value 27.

Example 3:
Input: n = 12
Output: 505379714
Explanation: The concatenation results in "1101110010111011110001001101010111100".
The decimal value of that is 118505380540.
After modulo 109 + 7, the result is 505379714.
*/

func concatenatedBinary(n int) int {
	const m = 1000000007
	result := 0
	shift := 1
	for i := 1; i <= n; i++ {
		// shift++ left if i bigger than 2^shift
		if i == 1<<shift {
			shift++
		}
		// shift result before add i to
		// add i by using |= (Example x |=y, just like x+=y, however it might be useful in case the value can't not be present as int64)
		result = result << shift
		result |= i
		result = result % m
	}
	return int(result)
}

func main() {
	concatenatedBinary(42)
}

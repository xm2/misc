package reverseinteger

import "math"

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [-2**31,  2**31 - 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverse(x int) int {
	ret := 0

	for {
		if x == 0 {
			break
		}

		pop := x % 10

		if ret > math.MaxInt32/10 || (ret == math.MaxInt32/10 && pop > 7) {
			// overflow 2**31-1
			return 0
		}

		if ret < math.MinInt32/10 || (ret == math.MinInt32/10 && pop < -8) {
			// overflow -2**31
			return 0
		}

		ret = ret*10 + pop
		x = x / 10
	}

	return ret

}

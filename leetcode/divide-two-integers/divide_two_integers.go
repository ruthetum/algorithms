package divide_two_integers

import (
	"math"
)

// https://leetcode.com/problems/divide-two-integers/

const (
	max = 2147483647
	min = -2147483648
)

func divide(dividend int, divisor int) int {
	if dividend == min && divisor == -1 {
		return max
	}

	if dividend == 0 {
		return 0
	}

	// without using multiplication, division, and mod operator.
	// return dividend/divisor

	sign := -1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}

	answer := 0
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	for dividend >= divisor {
		temp := divisor
		m := 1

		for temp<<1 <= dividend {
			temp = temp << 1
			m = m << 1
		}

		dividend -= temp
		answer += m
	}

	if sign > 0 {
		return answer
	}
	return -answer
}

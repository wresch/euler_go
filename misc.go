package euler_go

import (
	//"fmt"
)

// Reverse takes an integer n and returns the integer obtained
// by reversing the digits of n
func Reverse(n int64) int64 {
	var r int64 = 0
	for ;n > 0; n /= 10 {
		r = 10 * r + n % 10
	}
	return r
}

// Pow calculates x ** y for int64
func Pow(x, y int64) (int64) {
	var factor int64 = 1
	a, b := x, y
	for ;b > 1; {
		if b % 2 == 0 {
			a = a * a
			b /= 2
		} else {
			factor *= a
			a = a * a
			b = (b - 1) / 2
		}
	}
	return a * factor
}

func SumOfNaturalsUpTo(k int64) int64 {
	return (k * (k + 1)) / 2
}

func SumOfNaturalSquaresUpTo(k int64) int64 {
	return (k * (k + 1) * (2 * k + 1)) / 6
}


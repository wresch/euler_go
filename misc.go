package euler_go

import (
	"fmt"
)

// Reverse takes an integer n and returns the integer obtained
// by reversing the digits of n
func Reverse(n int64) int64 {
	var r int64 = 0
	for ; n > 0; n /= 10 {
		r = 10*r + n%10
	}
	return r
}

// Pow calculates x ** y for int64
func Pow(x, y int64) int64 {
	var factor int64 = 1
	a, b := x, y
	for b > 1 {
		if b%2 == 0 {
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
	return (k * (k + 1) * (2*k + 1)) / 6
}

func Output(i int, desc string, solution interface{}) {
	fmt.Println("%****************************************************************")
	fmt.Printf(`\section{Problem %4d}`, i)
	fmt.Println("\n%****************************************************************")
	fmt.Println(`\subsection{Description and approach}`)
	fmt.Printf("%s\n", desc)
	fmt.Println(`\subsection{Solution}`)
	fmt.Printf("%v\n", solution)
}

var numberSpellingDict = map[int64]string{
	0:             "zero",
	1:             "one",
	2:             "two",
	3:             "three",
	4:             "four",
	5:             "five",
	6:             "six",
	7:             "seven",
	8:             "eight",
	9:             "nine",
	10:            "ten",
	11:            "eleven",
	12:            "twelve",
	13:            "thirteen",
	14:            "fourteen",
	15:            "fifteen",
	16:            "sixteen",
	17:            "seventeen",
	18:            "eighteen",
	19:            "nineteen",
	20:            "twenty",
	30:            "thirty",
	40:            "forty",
	50:            "fifty",
	60:            "sixty",
	70:            "seventy",
	80:            "eighty",
	90:            "ninety",
	100:           "hundred",
	1000:          "thousand",
	1000000:       "million",
	1000000000:    "billion",
	1000000000000: "trillion",
}

// does not include comma; goes up to 1 billion - 1
func BritishSpelledNumber(n int64) (s string) {
	switch {
	case n < 0:
		s = "negative " + BritishSpelledNumber(-1*n)
	case n < 20:
		s = numberSpellingDict[n]
	case n < 100:
		tens := (n / 10) * 10
		rem := n % 10
		if rem == 0 {
			s = fmt.Sprintf(numberSpellingDict[tens])
		} else {
			s = fmt.Sprintf("%s-%s", numberSpellingDict[tens], numberSpellingDict[rem])
		}
	case n < 1000:
		houndreds := n / 100
		rem := n % 100
		if rem == 0 {
			s = fmt.Sprintf("%s hundred", numberSpellingDict[houndreds])
		} else {
			s = fmt.Sprintf("%s hundred and %s", numberSpellingDict[houndreds],
				BritishSpelledNumber(rem))
		}
	case n < 1000000:
		thousands := n / 1000
		rem := n % 1000
		if rem == 0 {
			s = fmt.Sprintf("%s thousand", BritishSpelledNumber(thousands))
		} else if rem < 100 {
			s = fmt.Sprintf("%s thousand and %s", BritishSpelledNumber(thousands),
				BritishSpelledNumber(n%1000))
		} else {
			s = fmt.Sprintf("%s thousand %s", BritishSpelledNumber(thousands),
				BritishSpelledNumber(n%1000))
		}
	case n < 1000000000:
		millions := n / 1000000
		rem := n % 1000000
		if rem == 0 {
			fmt.Sprintf("%s million", BritishSpelledNumber(millions))
		} else if rem < 100 {
			s = fmt.Sprintf("%s million and %s", BritishSpelledNumber(millions),
				BritishSpelledNumber(n%1000000))
		} else {
			s = fmt.Sprintf("%s million %s", BritishSpelledNumber(millions),
				BritishSpelledNumber(n%1000000))
		}
	default:
		s = "fail"
	}
	return
}

func TriangleMaxPathSum(triangle [][]int) (maxNum int) {
	sumTree := make([][]int, len(triangle))
	for ri, row := range triangle {
		sumTree[ri] = make([]int, ri+1)
		if ri == 0 {
			sumTree[0][0] = row[0]
		} else {
			for ci, n := range row {
				switch {
				case ci == 0:
					sumTree[ri][0] = n + sumTree[ri-1][0]
				case ci == ri:
					sumTree[ri][ci] = n + sumTree[ri-1][ci-1]
				case sumTree[ri-1][ci] > sumTree[ri-1][ci-1]:
					sumTree[ri][ci] = n + sumTree[ri-1][ci]
				default:
					sumTree[ri][ci] = n + sumTree[ri-1][ci-1]
				}
			}
		}
	}
	maxNum = 0
	for _, n := range sumTree[len(triangle)-1] {
		if n > maxNum {
			maxNum = n
		}
	}
	return
}

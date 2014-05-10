package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
A palindromic number reads the same both ways. The largest palindrome
made from the product of two 2-digit numbers is 9009 = 91 × 99.

\emph{Find the largest palindrome made from the product of two 3-digit
numbers.}

The first things to do is to only consider unique products and
start from the largest three digit number (since the largest
6 digit palindrome is sought).

Next, let us consider a palindromic 6 digit number with digits a, b, c:
\begin{equation}\begin{align}
        P &=& 100000a + 10000b + 1000c + 100c + 10b + a \\
          &=& 100001a + 10010b + 1100c \\
          &=& 11(9091a + 910b + 100c)
\end{align}\end{equation}
This means that P must have 11 as a factor, which means that one
of the 2 three digit numbers must have 11 as a factor (since 11 is
prime). This can be used to limit the search space.`

func main() {
	var largestPalindrome int64 = 0
	for a := int64(999); a >= 100; a-- {
		var bstart int64 = 990 // largest 3 digit number divisible by 11
		var bstep  int64 =  11
		if a % 11 == 0 {
			bstart = 999
			bstep  = 1
		}
		for b := bstart; b >= 100; b -= bstep {
			prod := a * b
			if prod <= largestPalindrome {
				break
			}
			if prod == eu.Reverse(prod) {
				largestPalindrome = prod
			}
		}
	}
	eu.Output(4, desc, largestPalindrome)
}

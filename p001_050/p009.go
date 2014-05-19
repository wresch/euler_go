package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc string = `
A Pythagorean triplet is a set of three natural numbers, $a < b < c$,
for which
\[a^2 + b^2 = c^2\]

For example,
\[3^2 + 4^2 = 9 + 16 = 25 = 5^2\].

\emph{There exists exactly one Pythagorean triplet for which $a + b + c
= 1000$.  Find the product abc.}

The solution to this problem makes use of Euclid's formula
\[
    a = m^2 - n^2 ,\ \, b = 2mn ,\ \, c = m^2 + n^2
\]
.  I tried to derive this formula myself, but did not succeed. This
needs more work
`

func main() {

	for n := int64(1); n < 100; n++ {
		for m := n + 1; m < 101; m++ {
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			if a+b+c == 1000 {
				eu.Output(9, desc, fmt.Sprintf("(%d, %d, %d); Product = %d",
					a, b, c, a*b*c))
				return
			}
		}
	}
	eu.Output(9, desc, "FAILED TO FIND SOLUTION")
}

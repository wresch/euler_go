package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

\emph{What is the smallest positive number that is evenly divisible by all
of the numbers from 1 to 20?}

Let \(n\) be the smallest number divisible by integers \(2..k\).

At a minimum, n has to contain all the prime numbers \(<= k\)
as factors.  The non-prime integers have to be factorized
and any factors not yet represented by just the prime numbers
need to be included.  For the case of 20:
\begin{equation}
2 * 3 * 2 * 5 * 7 * 2 * 3 * 11 * 13 * 2 * 17 * 19 = 232792560
\end{equation}`


func main() {
	var k int64 = 20
	ps := eu.NewPrimeSieve()
	factors := make(map[int64]int64)
	for i := int64(2); i <= k; i++ {
		if i == ps.Value {
			factors[i] = 1
			ps.Next()
		} else {
			ifact := eu.IntFreq(eu.PrimeFactors(i))
			for ii, ni := range(ifact) {
				if nc, ok := factors[ii]; ok {
					if ni > nc {
						factors[ii] = ni
					}
				} else {
					factors[ii] = ni
				}
			}
		}
	}
	var product int64 = 1
	for i, n := range factors {
		product *= eu.Pow(i, n)
	}
	eu.Output(5, desc, product)
}

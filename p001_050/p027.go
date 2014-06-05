package main

import (
	"fmt"
	eu "github.com/wresch/euler_go"
)

var desc = `


Euler discovered the remarkable quadratic formula:
\[
n^2 + n + 41
\]

It turns out that the formula will produce 40 primes for the
consecutive values n = 0 to 39. However, when $n = 40$,
\[
40^2 + 40 + 41 = 40(40 + 1) + 41
\]

is divisible by 41, and certainly when $n = 41$,
\[
41^2 + 41 + 41
\]
is clearly divisible by 41.

The incredible formula
\[
n^2 − 79n + 1601
\]

was discovered, which produces 80 primes for the consecutive values n
= 0 to 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:
\[
n^2 + an + b, where \abs{a} < 1000 and \abs{b} < 1000
\]

\emph{Find the product of the coefficients, a and b, for the quadratic
expression that produces the maximum number of primes for consecutive
values of n, starting with n = 0.}

\subsection{Approach}

There is probably much more math behind this, but there are some
properties that seem straight forward that should cut down on the
search space of a brute force solution considerably:
\begin{description}[style=nextline]
\item[b must be prime]
If b was not prime, then the formula would produce a non-prime number
for $n=0$
\item[a must be odd]
Since b, a prime, is odd, $n^2 + an = n(n + a)$ must be even, or the
resulting sum would be even, breaking the chain early. If n is even,
it doesn't matter if $n + a$ is even or odd.  However, if n is odd,
$n + a$ must be even, which means a must be odd.

This is still a brute force method and I am not sure it will
run fast enough.
`

func consecutivePrimeResults(a, b int, primes *eu.PrimeSet) (n int) {
	r := 0
	for n = 0; ; n++ {
		r = n*n + a*n + b
		if r < 0 {
			break
		}
		if isPrime, _ := primes.IsPrime(uint(r)); !isPrime {
			break
		}
	}
	return
}

func longestConsecutivePrimeResults(b int, primes *eu.PrimeSet) (aLongest, longest int) {
	current := 0
	for a := -999; a < 1000; a += 2 {
		current = consecutivePrimeResults(a, b, primes)
		if current > longest {
			longest = current
			aLongest = a
		}
	}
	return
}

func main() {
	// generate map of primes < 1e7
	primes := eu.NewPrimeSet(100000)
	longestSeries := 0
	aLongest := 0
	bLongest := 0
	for ps := eu.NewPrimeSieve(); ps.Value < 1000; ps.Next() {
		a, longest := longestConsecutivePrimeResults(int(ps.Value), primes)
		if longest > longestSeries {
			longestSeries = longest
			aLongest = a
			bLongest = int(ps.Value)
		}
	}
	eu.Output(27, desc, fmt.Sprintf(`
Coefficients a = %d and b = %d produce primes for values 0 to %d.
$a\cdot b=%d$.`, aLongest, bLongest, longestSeries-1, aLongest*bLongest))
	/* for n := 0; n <= longestSeries; n++ {
		r := n*n + aLongest*n + bLongest
		isPrime, _ := primes.IsPrime(uint(r))
		fmt.Printf("%2d^2 + %d * %d + %d = %4d (%v)\n", n, n, aLongest,
			bLongest, r, isPrime)
	} */
}

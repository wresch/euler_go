package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

\emph{Find the sum of all the primes below two million.}

Solution reused the prime sive from a previous problem.
Eventually, i probably need a better implementation of
the sieve.  The current implementation has the advantage
of creating an endless series, but it does not store
previous results and is relatively inefficient. This
code took 2m on helix.`

func main() {
	ps := eu.NewPrimeSieve()
	var total int64 = 0
	for ; ps.Value < 2000000; ps.Next() {
		total += ps.Value
	}
	eu.Output(10, desc, total)
}

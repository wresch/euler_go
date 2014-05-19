package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
Each new term in the Fibonacci sequence is generated by adding the
previous two terms. By starting with 1 and 2, the first 10 terms will
be:

	1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

\emph{By considering the terms in the Fibonacci sequence whose values do not
exceed four million, find the sum of the even-valued terms.}

This problem was solved by implementing a generator that supplies an
infinite list of numbers from the Fibonacci series and then summing
the even numbers.  This is essentially a brute force approach.
`

func main() {
	var total int64 = 0
	for i := eu.NewFibonacciCounter(); i.Value < 4000000; i.Next() {
		if i.Value%2 == 0 {
			total += i.Value
		}
	}
	eu.Output(2, desc, total)
}

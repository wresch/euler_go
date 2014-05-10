package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
If we list all the natural numbers below 10 that are multiples
of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

\emph{Find the sum of all the multiples of 3 or 5 below 1000}.

My solution to this is just brute force.
`


//////////////////// problem1 ////////////////////
func main() {
	var total int64 = 0
	for i := eu.NewLinearCounter(3, 3); i.Value < 1000; i.Next() {
		total += i.Value
	}
	for i := eu.NewLinearCounter(5, 5); i.Value < 1000; i.Next() {
		if i.Value % 3 == 0 {
			continue
		}
		total += i.Value
	}
	eu.Output(1, desc, total)
}

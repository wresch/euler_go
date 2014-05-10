package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can
see that the 6th prime is 13.

\emph{What is the 10 001st prime number?}

This just reuses the prime number generator from a previous problem
`

func main() {
	ps := eu.NewPrimeSieve()
	for i := 1; i < 10001; i++ {
		ps.Next()
	}
	eu.Output(7, desc, ps.Value)
}

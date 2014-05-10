package main

import (
	eu "github.com/wresch/euler_go"
)

var desc string = `
The prime factors of 13195 are 5, 7, 13 and 29.

\emph{What is the largest prime factor of the number 600851475143 ?}

Not very many optimizations - my algorithm basically uses Eratosthenes Sieve to
generate prime numbers and then does trial division with them up to
\(\sqrt(k)\).  `

func main() {
	prime_factors := eu.PrimeFactors(600851475143)
	eu.Output(3, desc, prime_factors)
}
